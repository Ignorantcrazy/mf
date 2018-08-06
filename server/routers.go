package main

func AllRoutes() Routes {
	routes := Routes{
		Route{Method: "GET",
			Path:   "/",
			Handle: Index},
		Route{Method: "GET",
			Path:   "/todos",
			Handle: TodoList,
			Middlewares: []Middleware{
				Loggin(),
				CORS(),
			}},
		Route{Method: "GET",
			Path:   "/todos/:id",
			Handle: TodoById,
			Middlewares: []Middleware{
				Loggin(),
			}},
		Route{Method: "POST",
			Path:   "/todos",
			Handle: TodoCreate,
			Middlewares: []Middleware{
				Loggin(),
			}},
	}
	return routes
}
