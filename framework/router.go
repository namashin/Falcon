package framework

import (
	"errors"
	"strings"
)

type Router struct {
	routingTables map[string]*TreeNode
	middlewares   []func(ctx *Context)
	noRoute       func(ctx *Context)
}

func (r *Router) Use(middleware func(ctx *Context)) {
	r.middlewares = append(r.middlewares, middleware)
}

func (r *Router) UseNoRoute(handler func(ctx *Context)) {
	r.noRoute = handler
}

func (r *Router) register(method string, pathname string, handler func(ctx *Context)) error {
	routingTable := r.routingTables[method]
	pathname = strings.TrimSuffix(pathname, "/")

	existedHandler := routingTable.Search(pathname)
	if existedHandler != nil {
		return errors.New("pathname already handled")
	}

	routingTable.Insert(pathname, handler)
	return nil
}

func (r *Router) Get(pathname string, handler func(ctx *Context)) {
	err := r.register("get", pathname, handler)
	if err != nil {
		panic(err)
	}
}

func (r *Router) Post(pathname string, handler func(ctx *Context)) {
	err := r.register("post", pathname, handler)
	if err != nil {
		panic(err)
	}
}

func (r *Router) Patch(pathname string, handler func(ctx *Context)) {
	err := r.register("patch", pathname, handler)
	if err != nil {
		panic(err)
	}
}

func (r *Router) Put(pathname string, handler func(ctx *Context)) {
	err := r.register("put", pathname, handler)
	if err != nil {
		panic(err)
	}
}

func (r *Router) Delete(pathname string, handler func(ctx *Context)) {
	err := r.register("delete", pathname, handler)
	if err != nil {
		panic(err)
	}
}
