package todo

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

type taskStruct struct {
	Task_id    int
	Task       string
	Created_at string
	Status     string
}

func Create(dbIns *sql.DB, task string) (int, error) {
	var task_id int

	if task == "" {
		return -1, fmt.Errorf("Cannot Create an empty task")
	}
	statement, err := dbIns.Prepare("INSERT INTO tasks(task, created_at, status) VALUES($1, $2, $3) RETURNING task_id;")
	if err != nil {
		return -1, err
	}
	rows := statement.QueryRow(task, fmt.Sprintf("%v-%d-%v", time.Now().Year(), int(time.Now().Month()), time.Now().Day()), false)
	rows.Scan(&task_id)

	return task_id, nil
}

func Read(dbIns *sql.DB, task_id int) (taskStruct, error) {

	statement, err := dbIns.Prepare("SELECT task, created_at, status FROM tasks WHERE task_id= $1;")
	if err != nil {
		return taskStruct{}, err
	}
	row := statement.QueryRow(task_id)
	var task string
	var created_at string
	var status bool
	err = row.Scan(&task, &created_at, &status)

	var doneOrNot string
	if status {
		doneOrNot = "Completed"
	} else {
		doneOrNot = "Not Completed"
	}

	if task == "" {
		return taskStruct{}, fmt.Errorf("Task Id is non-existent")
	}

	return taskStruct{
		Task_id:    task_id,
		Task:       task,
		Created_at: created_at,
		Status:     doneOrNot,
	}, err
}

func ShowAll(dbIns *sql.DB) ([]taskStruct, error) {

	var ts []taskStruct

	statement, err := dbIns.Prepare("SELECT task_id, task, created_at, status FROM tasks;")
	if err != nil {
		return ts, err
	}
	rows, err := statement.Query()
	if err != nil {
		return ts, err
	}

	var task string
	var task_id int
	var created_at string
	var status bool
	var doneOrNot string
	i := 0
	for rows.Next() {
		i++
		rows.Scan(&task_id, &task, &created_at, &status)
		if status {
			doneOrNot = "Completed"
		} else {
			doneOrNot = "Not Completed"
		}
		ts = append(ts, taskStruct{
			Task_id:    task_id,
			Task:       task,
			Created_at: created_at,
			Status:     doneOrNot,
		})
	}
	return ts, err
}

func Update(dbIns *sql.DB, task_id int, task string) error {

	if task == "" {
		return fmt.Errorf("Cannot update with an empty task")
	}
	statement, err := dbIns.Prepare("UPDATE tasks SET task = $1 WHERE task_id = $2 RETURNING task_id;")
	if err != nil {
		return err
	}
	row := statement.QueryRow(task, task_id)
	return row.Scan(&task_id)
}

func MarkDone(dbIns *sql.DB, task_id int) error {

	statement, err := dbIns.Prepare("UPDATE tasks SET status = true WHERE task_id = $1 RETURNING task_id;")
	if err != nil {
		return err
	}
	row := statement.QueryRow(task_id)
	return row.Scan(&task_id)
}

func Delete(dbIns *sql.DB, task_id int) error {

	statement, err := dbIns.Prepare("DELETE FROM tasks WHERE task_id = $1 RETURNING task_id;")

	if err != nil {
		return err
	}

	row := statement.QueryRow(task_id)
	return row.Scan(&task_id)
}
