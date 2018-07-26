package main

import (
	"github.com/julienschmidt/httprouter"
)

type Middleware func(httprouter.Handle) httprouter.Handle

func Chain(h httprouter.Handle, middlewares ...Middleware) httprouter.Handle {
	for _, m := range middlewares {
		h = m(h)
	}
	return h
}
