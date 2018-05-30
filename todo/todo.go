package todo

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

func Create(dbIns *sql.DB, task string) (int, error) {
	var task_id int

	if task == "" {
		return -1, fmt.Errorf("Cannot Create an empty task")
	}
	statement, err := dbIns.Prepare("INSERT INTO todo_table(task, timestamp) VALUES($1, $2) RETURNING task_id;")
	if err != nil {
		return -1, err
	}
	rows := statement.QueryRow(task, fmt.Sprintf("%v-%d-%v", time.Now().Year(), int(time.Now().Month()), time.Now().Day()))
	rows.Scan(&task_id)

	return task_id, nil
}

func Read(dbIns *sql.DB, task_id int) (string, error) {

	statement, err := dbIns.Prepare("SELECT task FROM todo_table WHERE task_id= $1;")
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

func ShowAll(dbIns *sql.DB) error {

	statement, err := dbIns.Prepare("SELECT task_id, task FROM todo_table;")
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

func Update(dbIns *sql.DB, task_id int, task string) error {

	if task == "" {
		return fmt.Errorf("Cannot update with an empty task")
	}
	statement, err := dbIns.Prepare("UPDATE todo_table SET task = $1 WHERE task_id = $2 RETURNING task_id;")
	if err != nil {
		return err
	}
	row := statement.QueryRow(task, task_id)
	return row.Scan(&task_id)
}

func Delete(dbIns *sql.DB, task_id int) error {

	statement, err := dbIns.Prepare("DELETE FROM todo_table WHERE task_id = $1 RETURNING task_id;")

	if err != nil {
		return err
	}

	row := statement.QueryRow(task_id)
	return row.Scan(&task_id)
}
