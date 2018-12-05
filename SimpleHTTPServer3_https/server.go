package main

import (
	"flag"
	"github.com/valyala/fasthttp"
	"log"
)

var addres = flag.String("addres", "127.0.0.1:8800",
	"TCP address to listen to for incoming connections")

func main() {
	flag.Parse()

	s := fasthttp.Server{
		Handler: handler1,
	}

	err := s.ListenAndServe(*addres)
	if err != nil {
		log.Fatalf("error in ListenAndServe: %s", err)
	}
}

func handler1(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("Hello, world!3\n")
}
