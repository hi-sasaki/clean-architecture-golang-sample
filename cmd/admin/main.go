package main

import (
	"net/http"

	"github.com/hi-sasaki/clean-architecture-golang-sample/pkg/driver/handler/chi"
	"github.com/hi-sasaki/clean-architecture-golang-sample/pkg/registry"
)

func main() {
	mux := chi.NewRouter(registry.NewProvider())
	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(err)
	}
}
