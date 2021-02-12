package main

import (
	"fmt"
	"log"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

func Index(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("Hello from docker!")
}

func Hello(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "Hi, %s!\n", ctx.UserValue("name"))
}

func main() {
	r := router.New()
	r.GET("/", Index)
	r.GET("/hello/{name}", Hello)

	log.Fatal(fasthttp.ListenAndServe(":3000", r.Handler))
}
