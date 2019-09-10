package main

import (
	"fmt"
	"github.com/elazarl/goproxy"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)

	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = true

	log.Fatal(http.ListenAndServe(":8080", proxy))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World from Go.")
}
