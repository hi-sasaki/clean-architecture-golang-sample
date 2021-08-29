package main

import (
	"net/http"
)

func main() {
	router := NewHandler()
	http.ListenAndServe(":3000", router)
}
