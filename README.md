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
