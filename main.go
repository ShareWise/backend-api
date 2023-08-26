package main

import (
	api "backend-api/api/product"
	"net/http"
)

func main() {
	http.HandleFunc("/", api.Handler)

	http.ListenAndServe(":8080", nil)
}
