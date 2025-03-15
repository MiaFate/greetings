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
	message, err := greetings.Hello("Gopher")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(message)
}
```