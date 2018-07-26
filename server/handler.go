package main

import (
	"fmt"
	"net/http"

	"github.com/Ignorantcrazy/mfgo"
	"github.com/julienschmidt/httprouter"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "Todo"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Index")
}

func TodoList(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	defer func() {
		if err := recover(); err != nil {
			result := Result{Code: "500", Data: err}
			mfgo.ResponseWithJson(w, result, 200)
		}
	}()
	connstr := mfgo.PGConnection{Host: host, Port: port, User: user, Password: password, Dbname: dbname}
	db := mfgo.PostgreSqlDB{Connstr: connstr, Sql: "select * from todo"}
	todos := db.Open().Query(GetTodos)
	result := Result{Code: "200", Data: todos}
	mfgo.ResponseWithJson(w, result, 200)
}

func TodoById(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	defer func() {
		if err := recover(); err != nil {
			result := Result{Code: "500", Data: err}
			mfgo.ResponseWithJson(w, result, 200)
		}
	}()
	id := ps.ByName("id")[1:]
	connstr := mfgo.PGConnection{Host: host, Port: port, User: user, Password: password, Dbname: dbname}
	db := mfgo.PostgreSqlDB{Connstr: connstr, Sql: "select * from todo where tid = $1"}
	todos := db.Open().Query(GetTodos, id)
	result := Result{Code: "200", Data: todos}
	mfgo.ResponseWithJson(w, result, 200)
}

func TodoCreate(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	defer func() {
		if err := recover(); err != nil {
			result := Result{Code: "500", Data: err}
			mfgo.ResponseWithJson(w, result, 200)
		}
	}()
	connstr := mfgo.PGConnection{Host: host, Port: port, User: user, Password: password, Dbname: dbname}
	db := mfgo.PostgreSqlDB{Connstr: connstr, Sql: "insert into todo(tid,title,content) VALUES ($1,$2,$3)"}
	affect := db.Open().Prepare("tid3", "title3", "content3")
	if affect > 0 {
		result := Result{Code: "200", Data: affect}
		mfgo.ResponseWithJson(w, result, 200)
	}
}

func TodoDelete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	defer func() {
		if err := recover(); err != nil {
			result := Result{Code: "500", Data: err}
			mfgo.ResponseWithJson(w, result, 200)
		}
	}()
	id := ps.ByName("id")
	connstr := mfgo.PGConnection{Host: host, Port: port, User: user, Password: password, Dbname: dbname}
	db := mfgo.PostgreSqlDB{Connstr: connstr, Sql: "delete from todo where tid = $1"}
	affect := db.Open().Prepare(id)
	if affect > 0 {
		result := Result{Code: "200", Data: affect}
		mfgo.ResponseWithJson(w, result, 200)
	}
}

func TodoUpdate(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	defer func() {
		if err := recover(); err != nil {
			result := Result{Code: "500", Data: err}
			mfgo.ResponseWithJson(w, result, 200)
		}
	}()
	id := ps.ByName("id")
	connstr := mfgo.PGConnection{Host: host, Port: port, User: user, Password: password, Dbname: dbname}
	db := mfgo.PostgreSqlDB{Connstr: connstr, Sql: "update todo set title = $1,content = $2 where id = $3"}
	affect := db.Open().Prepare("title", "content", id)
	if affect > 0 {
		result := Result{Code: "200", Data: affect}
		mfgo.ResponseWithJson(w, result, 200)
	}
}
