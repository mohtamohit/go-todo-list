package db

import (
	"database/sql"
	"fmt"
	"log"
	"practice/go-todo-list/config"
)

var db *sql.DB

func InitDB() *sql.DB {
	var err error

	fmt.Println(config.Db().ConnString())
	//jaadu := "postgres://localhost:5432/todo_db_test?sslmode=disable"
	//db, err = sql.Open("postgres", jaadu)
	db, err = sql.Open("postgres", config.Db().ConnString())
	if err != nil {
		log.Fatalf("failed to load the database: %s", err)
	}
	//	fmt.Println(err.Error)
	fmt.Println("Ping ke upar!!")
	if err = db.Ping(); err != nil {
		log.Fatalf("ping to database failed :%s", err)
	}
	return db
}

func Close() error {
	return db.Close()
}
