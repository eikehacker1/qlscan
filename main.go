package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/fatih/color"
	"strings"
	"io"
	"os"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/pflag"
	"bufio"
)

type QueryResponse struct {
	Data interface{} `json:"data"`
}

type Type struct {
	Name        string `json:"name"`
	Kind        string `json:"kind"`
	Description string `json:"description"`
	Fields      []struct {
		Name string `json:"name"`
	} `json:"fields"`
}

type Schema struct {
	Types []Type `json:"types"`
}

func fetchAndPrintQueries(url string) {
	query := `
	{
	  __schema {
		types {
		  name
		  kind
		  description
		  fields {
			name
		  }
		}
	  }
	}
	`

	resp, err := http.Post(url, "application/json", strings.NewReader(fmt.Sprintf(`{"query": %q}`, query)))
	if err != nil {
		color.Red("Erro ao fazer a solicitação para a URL '%s': %s", url, err)
		color.Yellow("URL não vulnerável, passando para a próxima.")
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		color.Red("Erro ao ler a resposta da URL '%s': %s", url, err)
		color.Yellow("URL não vulnerável, passando para a próxima.")
		return
	}

	if resp.StatusCode == http.StatusOK {
		var data QueryResponse
		if err := json.Unmarshal(body, &data); err != nil {
			color.Red("Erro ao analisar a resposta da URL '%s': %s", url, err)
			color.Yellow("URL não vulnerável, passando para a próxima.")
			return
		}

		schema, ok := data.Data.(map[string]interface{})["__schema"]
		if !ok {
			color.Red("Erro ao acessar o esquema da URL '%s'", url)
			color.Yellow("URL não vulnerável, passando para a próxima.")
			return
		}

		var schemaData Schema
		if err := mapstructure.Decode(schema, &schemaData); err != nil {
			color.Red("Erro ao analisar o esquema da URL '%s': %s", url, err)
			color.Yellow("URL não vulnerável, passando para a próxima.")
			return
		}

		color.Yellow("A URL '%s' é permitida para estas consultas:", url)
		for _, t := range schemaData.Types {
			if t.Kind == "OBJECT" {
				fmt.Printf("Query: %s\n", t.Name)
				fmt.Printf("Description: %s\n", t.Description)
				fields := []string{}
				for _, field := range t.Fields {
					fields = append(fields, field.Name)
				}
				fmt.Printf("Fields: %s\n\n", strings.Join(fields, ", "))
			}
		}
	} else {
		color.Red("A solicitação para a URL '%s' falhou com código de status %d: %s", url, resp.StatusCode, string(body))
		color.Yellow("URL não vulnerável, passando para a próxima.")
	}
}

func main() {
	var url string
	pflag.StringVar(&url, "url", "", "URL do servidor GraphQL")
	pflag.Parse()

	if url != "" {
		fetchAndPrintQueries(url)
	} else {
		stat, _ := os.Stdin.Stat()
		if (stat.Mode() & os.ModeCharDevice) == 0 {
			scanner := bufio.NewScanner(os.Stdin)
			for scanner.Scan() {
				url = scanner.Text()
				fetchAndPrintQueries(url)
			}
		} else {
			color.Red("Nenhuma URL válida fornecida.")
		}
	}
}
