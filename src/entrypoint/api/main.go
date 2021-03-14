package main

import "net/http"

func main() {
	controller := initializeController()
	http.ListenAndServe(":3000", handler(controller))
}