package main

import "database/sql"

type Todo struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type Todos []Todo

func GetTodos(rows *sql.Rows) interface{} {
	var todos Todos
	for rows.Next() {
		var id, title, content string
		err := rows.Scan(&id, &title, &content)
		if err != nil {
			panic(err)
		}
		todos = append(todos, Todo{ID: id, Title: title, Content: content})
	}
	return todos
}
