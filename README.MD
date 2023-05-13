### Typical go-gin project directory structure
```
|-
  |-app
    |-config
    |-controllers
    |-middlewares
    |-models
    |-routes
    |-services
    - main.go
  - go.mod
```
### Colorful printing lib

```Go
package main

import (
  "fmt"
  "github.com/fatih/color" // 在console中着色
)

func main() {
  blueStr := color.New(color.FgBlue).Sprintf("console output: blue")
  fmt.Println(blueStr)
}
```
