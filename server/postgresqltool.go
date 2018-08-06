package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type PGConnection struct {
	Host     string
	Port     int
	User     string
	Password string
	Dbname   string
}

type PostgreSqlDB struct {
	Connstr PGConnection
	Sql     string
	DB      *sql.DB
}

type QueryResult func(rows *sql.Rows) interface{}

func (dbs *PostgreSqlDB) Open() *PostgreSqlDB {
	connstr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbs.Connstr.Host, dbs.Connstr.Port, dbs.Connstr.User, dbs.Connstr.Password, dbs.Connstr.Dbname)
	db, err := sql.Open("postgres", connstr)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	dbs.DB = db
	return dbs
}

func (dbs *PostgreSqlDB) Query(qr QueryResult, args ...interface{}) interface{} {
	defer dbs.DB.Close()

	rows, err := dbs.DB.Query(dbs.Sql, args...)
	if err != nil {
		panic(err)
	}
	return qr(rows)
}

func (dbs *PostgreSqlDB) Prepare(args ...interface{}) int64 {
	defer dbs.DB.Close()

	stmt, _ := dbs.DB.Prepare(dbs.Sql)
	res, _ := stmt.Exec(args...)
	affect, _ := res.RowsAffected()
	return affect
}
