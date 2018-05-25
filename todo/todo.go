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
		return 0, err
	}

	var task_id int
	statement, err := db.Prepare("INSERT INTO todo_db(task) VALUES($1);")
	if err != nil {
		fmt.Println("Encountered Error: ", err)
		return 0, err
	}
	row, err := statement.Query(task)
	row.Scan(&task_id)

	if err != nil {
		fmt.Println("Encountered Error: ", err)
		return 0, err
	}
	return task_id, err
}

func Read(task_id int) (string, error) {
	db, err := dbConn(user, password, dbname)
	defer db.Close()

	if err != nil {
		fmt.Println("Encountered error: ", err)
		return "", err
	}

	statement, err := db.Prepare("SELECT task FROM todo_db WHERE task_id= $1;")
	if err != nil {
		fmt.Println("Enountered error: ", err)
		return "", err

	}
	row, err := statement.Query(task_id)
	if err != nil {
		fmt.Println("Encountered error: ", err)
		return "", err
	}

	var task string
	//i := 0
	for row.Next() {
		//i++
		row.Scan(&task)
		//fmt.Println("", i, " ", task)
		//fmt.Println(task)
	}

	return task, err
}
