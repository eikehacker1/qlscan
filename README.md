# qlscan

This script made in "GO" checks some initial series of graqphql APIs. It only performs a basic introspection query for you to begin your exploration.
#### Example:

````JSON
'''
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
'''
````
Example endpoints that may be consulted:
````URL
/graphql/v1
/graphql/v2
/graphql
/api/v1
/api
/api/graphql
/graphql/api
/graphql/graphql
````
These are just examples, you are the one who recognizes the endpoints
# INSTALL:
````zsh
go install -v github.com/eikehacker1/qlscan@latest
````

## USE: 
````bash
cat urls.txt | qlscan 
````
## OR:
````bash
qlscan --url https://exemplo.com/v1 
````
![exemplo](https://raw.githubusercontent.com/eikehacker1/qlscan/main/Captura%20de%20tela%202023-10-14%20090726.png)
