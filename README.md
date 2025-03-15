# Greetings GO

This is a simple program that prints a greetings message to the console.

## Instalation

```bash
go get -u github.com/miafate/greetings
```

## Usage

```go
package main

import (
    "fmt"
    "github.com/miafate/greetings"
)

func main() {
    message := greetings.Hello("Gopher")
    fmt.Println(message)
}
```