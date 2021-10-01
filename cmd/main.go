package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/hi-sasaki/clean-architecture-golang-sample/pkg/driver/handler/chi"
	"github.com/hi-sasaki/clean-architecture-golang-sample/pkg/registry"
)

var addr = ":8080"

func main() {
	mux := chi.NewRouter(registry.NewProvider())
	fmt.Printf("[START] server. port: %s\n", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}
