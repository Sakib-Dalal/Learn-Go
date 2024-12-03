## Learning Go 👨‍💻

- To run Go app in terminal 
```zsh
go run main.go 
```

- To build executable Go app use 
```zsh
go build main.go 
```

- To execute Go app use 

```zsh
./main 
```

- The ```main.go``` app is considered as module, while any go app is considered as module.

- Command to add Go app as module (Turn's this project into a module)
```zsh
 go mod init example.com/first-app 
```
```zsh
 go build 
```
```zsh
 ./first-app 
```

- Only the main.go should consider ```main()``` function, no other Go file should contain ```main()``` function.