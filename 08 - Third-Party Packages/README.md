## Go Third Party Package

- get package from: https://pkg.go.dev/

- Build main package
```sh
go mod init example.com/packages
```

- Example package link: https://pkg.go.dev/github.com/Pallinder/go-randomdata#section-readme
- Install Package command
```sh
go get github.com/Pallinder/go-randomdata
```

- Import Package in ```main.go```
```go
import (
    "fmt"
    "github.com/Pallinder/go-randomdata"
)
```

- Command to install already described packages
```sh
go get
```