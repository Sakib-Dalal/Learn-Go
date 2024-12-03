## String Formatting

- <a href="https://pkg.go.dev/fmt#example-Printf">Printf</a>
- ```go
 	fmt.Printf("This is number %v which is int data type\n", num) ```
- ```go
 	fmt.Printf("This is number %v and square is %v \n", num, num*num)```

<br>

- <a href="https://pkg.go.dev/fmt#Sprintf">Sprintf</a>
```go
combinedString := fmt.Sprintf("The number is %d, and float is %f and string is '%s'.\n", num, floatNum, stringVal) 
```