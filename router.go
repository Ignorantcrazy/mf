package main

import (
	"github.com/julienschmidt/httprouter"
)

type Route struct {
	Method      string
	Path        string
	Handle      httprouter.Handle
	middlewares []Middleware
}

type Routes []Route

func NewRoute(routes Routes) *httprouter.Router {
	router := httprouter.New()
	for _, route := range routes {
		handle := Chain(route.Handle, route.middlewares...)
		router.Handle(route.Method, route.Path, handle)
	}
	return router
}
