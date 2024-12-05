## Map

Map works similar to dictionaries in python.
Map are dynamic while structure are not.
Map allow you any value as key while struct not.

- Creating map
```go
websites := map[string]string{
    "Google": "https://google.com",
    "Amazon": "https://amazon.com",
} // [type of keys]values datatype to be used{"key": "value"} // empty map {}
```

- print map
```go
fmt.Println(websites)
```

- pint values using keys
```go
fmt.Println(websites["Google"])
fmt.Println(websites["Amazon"])
```

- app new key value pairs
```go
// add new key value pairs in map[type]type
websites["Github"] = "https://github.com"
fmt.Println(websites)
```

- delete key value in map
```go
delete(websites, "Google")
fmt.Println(websites)
```

- create map() using make() function

```go
// create map() using make() function
animals := make(map[string]string, 3) // create map of keys string and values string has capacity of 3
animals["cat"] = "simple and cute"
animals["tiger"] = "dangerous and bad"
animals["dog"] = "honest and cute"

fmt.Println(animals)
```