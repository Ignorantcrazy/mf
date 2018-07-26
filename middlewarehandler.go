package main

import (
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

func Loggin() Middleware {
	return func(h httprouter.Handle) httprouter.Handle {
		return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
			start := time.Now()
			log.Printf(" Started %s %s", r.Method, r.URL.Path)
			h(w, r, ps)
			log.Printf(" Completed in %v", time.Since(start))
		}
	}
}
