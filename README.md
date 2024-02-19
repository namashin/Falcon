# Falcon

Falcon is a lightweight and flexible web framework for Go, inspired by Gin.

## Tags

- `Go`
  
## Features

- Simple and intuitive routing
- Middleware support for customizing request handling
- Easy-to-use API for building web applications

## Installation

**Clone or download the Falcon repository by following these steps:**
    - If you have Git installed, open a terminal and run the following command:
    
```bash
git clone https://github.com/namashin/Falcon.git
```

### Running Falcon example

```go
package main

import (
	"fmt"
	"framework/framework"
)

func main() {
	e := framework.NewEngine()

	e.Router.Get("/test", func(ctx *framework.Context) {
		fmt.Fprint(ctx.ResponseWriter(), "test")
	})

	e.Run()
}
```
