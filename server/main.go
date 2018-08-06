package main

import (
	"log"
	"net/http"
)

func main() {
	r := NewRoute(AllRoutes())
	log.Fatal(http.ListenAndServe(":8080", r))
}
