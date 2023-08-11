package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/elazarl/goproxy"
)

var laddr string

func main() {
	flag.StringVar(&laddr, "l", ":8088", "listen addr")

	flag.Parse()

	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = true
	log.Fatal(http.ListenAndServe(laddr, proxy))
}
