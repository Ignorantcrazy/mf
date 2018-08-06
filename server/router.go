package main

import (
	"github.com/julienschmidt/httprouter"
)

type Route struct {
	Method      string
	Path        string
	Handle      httprouter.Handle
	Middlewares []Middleware
}

type Routes struct {
	Routes      []Route
	Middlewares []Middleware
}

func NewRoute(routes Routes) *httprouter.Router {
	router := httprouter.New()
	for _, route := range routes.Routes {
		for _, middle := range routes.Middlewares {
			route.Middlewares = append(route.Middlewares, middle)
		}
		handle := Chain(route.Handle, route.Middlewares...)
		router.Handle(route.Method, route.Path, handle)
	}
	return router
}

func (routes *Routes) Use(middleware Middleware) {
	routes.Middlewares = append(routes.Middlewares, middleware)
}
