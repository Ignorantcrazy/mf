package main

import (
	"log"
	"net/http"

	"github.com/Ignorantcrazy/mfgo"
)

func main() {
	r := mfgo.NewRoute(AllRoutes())
	log.Fatal(http.ListenAndServe(":8080", r))
}
