package todo

import (
	"fmt"
	"practice/go-todo-list/db"
	"time"

	_ "github.com/lib/pq"
)

func Create(task string) (int, error) {
	db := db.InitDB()
	// db := dbConn(user, password, dbname)
	defer db.Close()

	var task_id int
	statement, err := db.Prepare("INSERT INTO todo_db(task, timestamp) VALUES($1, $2);")
	if err != nil {
		fmt.Println("Encountered Error: ", err)
		return 0, err
	}
	row := statement.QueryRow(task, time.Now())
	row.Scan(&task_id)
	return task_id, nil
}

func Read(task_id int) (string, error) {
	// db := dbConn(user, password, dbname)
	db := db.InitDB()
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
	// db := dbConn(user, password, dbname)
	db := db.InitDB()
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
	// db := dbConn(user, password, dbname)
	db := db.InitDB()
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
	// db := dbConn(user, password, dbname)
	db := db.InitDB()
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
