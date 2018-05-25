package todo

import (
	"database/sql"
	"fmt"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = ""
	dbname   = "todo_db"
)

func dbConn(user, password, dbname string) (*sql.DB, error) {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		user, password, dbname)
	db, err := sql.Open("postgres", dbinfo)
	return db, err
}

func Create(task string) (int, error) {
	db, err := dbConn(user, password, dbname)
	defer db.Close()

	if err != nil {
		fmt.Errorf("Encountered error: ", err)
	}

	var taskID int
	err = db.QueryRow("INSERT INTO todo_db(task) VALUES ($1);", task).Scan(&taskID)
	if err != nil {
		fmt.Errorf("Error: ", err)
	}
	return taskID, err
}

func Read(taskID int) {
}
