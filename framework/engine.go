package framework

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
)

type Engine struct {
	Router *Router
}

func NewEngine() *Engine {
	return &Engine{
		Router: &Router{
			routingTables: map[string]*TreeNode{
				"get":    NewTreeNode(),
				"post":   NewTreeNode(),
				"patch":  NewTreeNode(),
				"put":    NewTreeNode(),
				"delete": NewTreeNode(),
			},
			middlewares: make([]func(ctx *Context), 0),
			noRoute:     DefaultNotFoundHandler,
		},
	}
}

func (e *Engine) Run() {
	ch := make(chan os.Signal)
	signal.Notify(ch)

	server := &http.Server{Addr: "localhost:8080", Handler: e}

	go func() {
		log.Fatal(server.ListenAndServe())
	}()

	<-ch

	if err := server.Shutdown(context.Background()); err != nil {
		fmt.Println("Shutdown failed")
		return
	}

	fmt.Println("shutdown successes")
}

func (e *Engine) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	ctx := NewContext(rw, r)

	routingTable := e.Router.routingTables[strings.ToLower(r.Method)]

	pathname := r.URL.Path
	pathname = strings.TrimSuffix(pathname, "/")

	var targetHandler func(ctx *Context)

	targetNode := routingTable.Search(pathname)

	if targetNode == nil || targetNode.handler == nil {

		targetHandler = e.Router.noRoute
		if targetHandler == nil {
			targetHandler = DefaultNotFoundHandler
		}

	} else {
		targetHandler = targetNode.handler
		paramDicts := targetNode.ParseParams(r.URL.Path)
		ctx.SetParams(paramDicts)
	}

	handlers := append(e.Router.middlewares, targetHandler)
	ctx.SetHandlers(handlers)
	ctx.Next()
}

func DefaultNotFoundHandler(ctx *Context) {
	ctx.ResponseWriter().WriteHeader(http.StatusNotFound)
}
