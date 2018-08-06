package main

func AllRoutes() Routes {
	routes := Routes{
		Routes: []Route{
			Route{Method: "GET",
				Path:   "/",
				Handle: Index},
			Route{Method: "GET",
				Path:   "/todos",
				Handle: TodoList,
				Middlewares: []Middleware{
					CORS(),
				}},
			Route{Method: "GET",
				Path:   "/todos/:id",
				Handle: TodoById,
			},
			Route{Method: "POST",
				Path:   "/todos",
				Handle: TodoCreate,
			},
		},
		Middlewares: []Middleware{
			Loggin(),
		},
	}
	return routes
}
