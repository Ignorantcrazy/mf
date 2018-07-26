package main

import (
	"github.com/Ignorantcrazy/mfgo"
)

func AllRoutes() mfgo.Routes {
	routes := mfgo.Routes{
		mfgo.Route{Method: "GET",
			Path:   "/",
			Handle: Index},
		mfgo.Route{Method: "GET",
			Path:   "/todos",
			Handle: TodoList,
			Middlewares: []mfgo.Middleware{
				Loggin(),
			}},
		mfgo.Route{Method: "GET",
			Path:   "/todos/:id",
			Handle: TodoById,
			Middlewares: []mfgo.Middleware{
				Loggin(),
			}},
		mfgo.Route{Method: "POST",
			Path:   "/todos",
			Handle: TodoCreate,
			Middlewares: []mfgo.Middleware{
				Loggin(),
			}},
	}
	return routes
}
