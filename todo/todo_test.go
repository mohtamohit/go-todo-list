package todo

import (
	"fmt"
	"os"
	"practice/go-todo-list/config"
	"practice/go-todo-list/db"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	task         = "some random testing task"
	taskID       = 1
	updateTask   = "updated testing task"
	updateTaskID = 1
	deleteTaskID = 1
)

func TestCreate(t *testing.T) {
	os.Setenv("ENVIRONMENT", "test")
	config.Load()
	db := db.InitDB()
	task_id, err := Create(task)
	require.NoError(t, err)

	fmt.Println(task_id)

	var task_check string
	s, err := db.Prepare("SELECT task AS task_check FROM todo_table WHERE task_id = $1")
	rows := s.QueryRow(task_id)
	rows.Scan(&task_check)

	require.Equal(t, task, task_check)
	require.NoError(t, err)
	// db.Exec("truncate table todo_table;")
}

func TestReadForExistingTask(t *testing.T) {
	os.Setenv("ENVIRONMENT", "test")
	config.Load()

	db := db.InitDB()
	db.Exec("INSERT INTO todo_table(task_id, task, timestamp) VALUES(2, 'randomTask', '2018-05-18');")

	task, err := Read(taskID)
	require.NoError(t, err)
	require.Equal(t, task, "randomTask")

	// db.Exec("truncate table todo_table;")
}

func TestReadForNoTask(t *testing.T) {
	os.Setenv("ENVIRONMENT", "test")
	config.Load()
	_, err := Read(10000000)
	require.EqualError(t, err, "Task Id is non-existent")
}

func TestShowAll(t *testing.T) {
	os.Setenv("ENVIRONMENT", "test")
	config.Load()
	err := ShowAll()
	require.NoError(t, err)
}

func TestUpdate(t *testing.T) {
	os.Setenv("ENVIRONMENT", "test")
	config.Load()
	db := db.InitDB()
	var task string
	db.Exec("INSERT INTO todo_table(task_id, task, timestamp) VALUES(1, 'randomTask', '2018-05-18');")

	err := Update(1, "updatedTask")

	row := db.QueryRow("SELECT task from todo_table where task_id=1")
	row.Scan(&task)
	require.Equal(t, task, "updatedTask")
	require.NoError(t, err)
	db.Exec("truncate table todo_table;")
}

func TestDelete(t *testing.T) {
	os.Setenv("ENVIRONMENT", "test")
	config.Load()
	err := Delete(deleteTaskID)
	require.NoError(t, err)
}
