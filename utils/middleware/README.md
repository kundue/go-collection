# Middleware

## Usage

```go
package main

import (
	"fmt"

  "github.com/kundue/go-collection/utils/middleware"
)

func main() {
	// Create an instance of app
	app := middleware.App{}

	// Define some sample handler functions
	handler1 := func() {
		fmt.Println("Handler 1")
	}

	handler2 := func() {
		fmt.Println("Handler 2")
	}

	handler3 := func() {
		fmt.Println("Handler 3")
	}

	// Collect the handlers
	app.Use(handler1, handler2)
	app.Use(handler3)

	// Run all collected handlers
	app.Run()
}
```