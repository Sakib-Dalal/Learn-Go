## for loop with arrays, slice, maps

- for loop
```go
for i := 0; i < 3; i++ {
	fmt.Println("Hello World")
}
```  

- for loop through arrays
```go
myArray := []int{1, 2, 3, 4, 5}
for i := range myArray {
    fmt.Println(myArray[i], i)
}
```

- for loop through maps
```go
myMap := map[string]int{
    "macOS":   10,
    "Linux":   9,
    "Windows": 5,
}
for key, value := range myMap {
    fmt.Println(key, value)
}
```