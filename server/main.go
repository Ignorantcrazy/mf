package main

import (
	"log"
	"net/http"
)

func main() {
	routers := AllRoutes()
	//routers.Use(Loggin())
	r := NewRoute(routers)
	log.Fatal(http.ListenAndServe(":8080", r))
}
