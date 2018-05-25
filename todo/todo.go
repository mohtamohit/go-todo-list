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

func dbConn(user, password, dbname string) *sql.DB {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		user, password, dbname)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		fmt.Println("Encountered error: ", err)
		panic(err)
	}
	return db
}

func Create(task string) (int, error) {
	db := dbConn(user, password, dbname)
	defer db.Close()

	var task_id int
	statement, err := db.Prepare("INSERT INTO todo_db(task) VALUES($1);")
	if err != nil {
		fmt.Println("Encountered Error: ", err)
		return 0, err
	}
	row := statement.QueryRow(task)
	row.Scan(&task_id)
	return task_id, err
}

func Read(task_id int) (string, error) {
	db := dbConn(user, password, dbname)
	defer db.Close()

	statement, err := db.Prepare("SELECT task FROM todo_db WHERE task_id= $1;")
	if err != nil {
		fmt.Println("Enountered error: ", err)
		return "", err
	}
	row := statement.QueryRow(task_id)
	var task string
	row.Scan(&task)
	if task == "" {
		return task, fmt.Errorf("Task Id is non-existent")
	}
	return task, err
}

func ShowAll() error {
	db := dbConn(user, password, dbname)
	defer db.Close()

	statement, err := db.Prepare("SELECT task FROM todo_db;")
	if err != nil {
		fmt.Println("Enountered error: ", err)
		return err
	}
	rows, err := statement.Query()
	if err != nil {
		fmt.Println("Encountered error: ", err)
		return err
	}

	var task string
	i := 0
	for rows.Next() {
		i++
		rows.Scan(&task)
		fmt.Println("", i, " ", task)
	}
	return err
}

func Update(task_id int, task string) error {
	db := dbConn(user, password, dbname)
	defer db.Close()

	statement, err := db.Prepare("UPDATE todo_db SET task = $1 WHERE task_id = $2;")
	if err != nil {
		fmt.Println("Enountered error: ", err)
		return err
	}
	row := statement.QueryRow(task, task_id)
	row.Scan(&task_id)
	return err
}

func Delete(task_id int) error {
	db := dbConn(user, password, dbname)
	defer db.Close()

	statement, err := db.Prepare("DELETE FROM todo_db WHERE task_id = $1;")

	if err != nil {
		fmt.Println("Enountered error: ", err)
		return err
	}

	row := statement.QueryRow(task_id)
	row.Scan(&task_id)

	return err
}
