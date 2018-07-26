package main

func AllRoutes() Routes {
	routes := Routes{
		Route{Method: "GET",
			Path:   "/",
			Handle: Index},
		Route{Method: "GET",
			Path:   "/todos",
			Handle: TodoList,
			middlewares: []Middleware{
				Loggin(),
			}},
		Route{Method: "GET",
			Path:   "/todos/:id",
			Handle: TodoById,
			middlewares: []Middleware{
				Loggin(),
			}},
		Route{Method: "POST",
			Path:   "/todos",
			Handle: TodoCreate,
			middlewares: []Middleware{
				Loggin(),
			}},
	}
	return routes
}
