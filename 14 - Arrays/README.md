## Arrays

- example arrays
```go
myArray := [4]float64{1.2, 2.2, 3.1, 5.6} // create float array // store 4 values
fmt.Println(myArray)
```

- another int array with 10 spaces
```go
prices := [10]int{1, 2, 3}
fmt.Println(prices)
```
```text
output:
[1 2 3 0 0 0 0 0 0 0]
```

- string array
```go
var productName [4]string = [4]string{"di", "go"}
fmt.Println(productName)
```
```text
output:
[di go  ]
```

- access value by index in array
```go
// access value by index in array
fmt.Println(prices[0])
fmt.Println(prices[5])
fmt.Println(prices[1])
```
```text
output:
1
0
2
```

- set value of index in array
```go
// set value of index in array
prices[5] = 50
fmt.Println(prices)
```
```text
output: 
[1 2 3 0 0 50 0 0 0 0]
```

- slicing in arrays
```go
// slicing in arrays
fmt.Println(prices[2:6])
fmt.Println(prices[:4])
fmt.Println(prices[4:])
```
```text
output:
[3 0 0 50]
[1 2 3 0]
[0 50 0 0 0 0]
```

- get array length and capacity
```go
fmt.Println(len(prices), cap(prices))
```

- Dynamic Array
```go
// dynamic arrays // auto scalable array size
dynamicArray := []int{1, 23, 3}
fmt.Println(dynamicArray)
// add new value in dynamic array
updatedDynamicArray := append(dynamicArray, 4)
fmt.Println(updatedDynamicArray, dynamicArray)
// reassign array
dynamicArray = updatedDynamicArray
fmt.Println(dynamicArray)
```
```text
output:
[1 23 3]
[1 23 3 4] [1 23 3]
[1 23 3 4]
```