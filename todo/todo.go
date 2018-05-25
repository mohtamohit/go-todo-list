package todo

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
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
		fmt.Println("Encountered error:", err)
	}

	var task_id int
	statement, err := db.Prepare("INSERT INTO todo_db(task) VALUES($1);")
	if err != nil {
		fmt.Println("Encountered Error: ", err)
	}
	row, err := statement.Query(task)
	row.Scan(&task_id)

	if err != nil {
		fmt.Println("Encountered Error: ", err)
	}
	return task_id, err
}

func Read(task_id int) (string, error) {
	return "", nil
}
