package todo

import (
	"fmt"
	"practice/go-todo-list/db"

	_ "github.com/lib/pq"
)

func Create(task string) (int, error) {
	db := db.InitDB()
	defer db.Close()
	var task_id int

	statement, err := db.Prepare("INSERT INTO todo_table(task, timestamp) VALUES($1, $2) RETURNING task_id;")
	if err != nil {
		return -1, err
	}
	rows := statement.QueryRow(task, "2018-06-18")
	rows.Scan(&task_id)

	return task_id, nil
}

func Read(task_id int) (string, error) {
	db := db.InitDB()
	defer db.Close()

	statement, err := db.Prepare("SELECT task FROM todo_table WHERE task_id= $1;")
	if err != nil {
		return "", err
	}
	row := statement.QueryRow(task_id)
	var task string
	err = row.Scan(&task)

	if task == "" {
		return task, fmt.Errorf("Task Id is non-existent")
	}
	return task, err
}

func ShowAll() error {
	db := db.InitDB()
	defer db.Close()

	statement, err := db.Prepare("SELECT task_id, task FROM todo_table;")
	if err != nil {
		return err
	}
	rows, err := statement.Query()
	if err != nil {
		return err
	}

	var task string
	var task_id int
	i := 0
	for rows.Next() {
		i++
		rows.Scan(&task_id, &task)
		fmt.Println(task_id, " ", task)
	}
	return err
}

func Update(task_id int, task string) error {
	db := db.InitDB()
	defer db.Close()

	statement, err := db.Prepare("UPDATE todo_table SET task = $1 WHERE task_id = $2 RETURNING task_id;")
	if err != nil {
		return err
	}
	row := statement.QueryRow(task, task_id)
	return row.Scan(&task_id)
}

func Delete(task_id int) error {
	db := db.InitDB()
	defer db.Close()

	statement, err := db.Prepare("DELETE FROM todo_table WHERE task_id = $1;")

	if err != nil {
		return err
	}

	row := statement.QueryRow(task_id)
	return row.Scan(&task_id)
}
