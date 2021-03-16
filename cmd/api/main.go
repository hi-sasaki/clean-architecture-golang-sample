package main

import (
	"net/http"
)

func main() {
	router, _ :=InitializeRouter()
	http.ListenAndServe(":3000", router)
}