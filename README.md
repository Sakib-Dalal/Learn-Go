## Learning Go 👨‍💻

- To run Go app in terminal 
``` go run main.go ```

- To build executable Go app use 
``` go build main.go ```

- To execute Go app use 
``` ./main ```

- The ```main.go``` app is considered as module, while any go app is considered as module.

- Command to add Go app as module (Turn's this project into a module)
``` go mod init example.com/first-app ```
``` go build ```
``` ./first-app ```

- Only the main.go should consider ```main()``` function, no other Go file should contain ```main()``` function.