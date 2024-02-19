package framework

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"path"
	"strings"
	"time"
)

const TimeOutSecond = 5

func TimeOutMiddleWare(ctx *Context) {
	successCh := make(chan struct{})
	panicCh := make(chan struct{})
	durationContext, cancel := context.WithTimeout(ctx.Request().Context(), TimeOutSecond*time.Second)
	defer cancel()

	go func() {
		defer func() {
			if err := recover(); err != nil {
				panicCh <- struct{}{}
			}
		}()
		ctx.Next()
		successCh <- struct{}{}
	}()

	select {
	case <-durationContext.Done():
		ctx.WriteString("timeout")
		ctx.SetHasTimeout(true)
	case <-panicCh:
		ctx.WriteString("panic")
	case <-successCh:
		fmt.Println("success")
	}
}

func StaticFileMiddleware(ctx *Context) {
	fileServer := http.FileServer(http.Dir("./static"))

	pathname := ctx.Request().URL.Path
	pathname = strings.TrimSuffix(pathname, "/")

	fPath := path.Join("./static", pathname)
	fInfo, err := os.Stat(fPath)

	if err == nil && !fInfo.IsDir() {
		fileServer.ServeHTTP(ctx.ResponseWriter(), ctx.Request())
		ctx.Abort()
		return
	}
}
