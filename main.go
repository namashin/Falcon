package main

import (
	"fmt"
	"framework/framework"
)

func main() {
	/* example */
	e := framework.NewEngine()

	e.Router.Get("/test", func(ctx *framework.Context) {
		fmt.Fprint(ctx.ResponseWriter(), "test")
	})

	e.Run()
}
