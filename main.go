package main

import (
	routes "backend-api/routes/product"
	"net/http"
)

func main() {
	http.HandleFunc("/", routes.Handler)
	http.HandleFunc("/test", routes.Test)

	http.ListenAndServe(":8080", nil)
}
