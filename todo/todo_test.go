package todo

import (
	"os"
	"practice/go-todo-list/config"
	"practice/go-todo-list/db"
	"testing"

	"github.com/stretchr/testify/assert"
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
	assert.NoError(t, err)

	var task_check string
	s, err := db.Prepare("SELECT task AS task_check FROM todo_table WHERE task_id = $1")
	rows := s.QueryRow(task_id)
	rows.Scan(&task_check)

	db.Exec("truncate table todo_table;")
	assert.Equal(t, task, task_check)
	assert.NoError(t, err)
}

func TestCannotCreateEmptyTask(t *testing.T) {
	os.Setenv("ENVIRONMENT", "test")
	config.Load()
	task_id, err := Create("")
	assert.EqualError(t, err, "Cannot Create an empty task")
	assert.Equal(t, -1, task_id)
}
func TestReadForExistingTask(t *testing.T) {
	os.Setenv("ENVIRONMENT", "test")
	config.Load()

	db := db.InitDB()
	var task_id int
	statement, err := db.Prepare("INSERT INTO todo_table(task, timestamp) VALUES($1, $2) RETURNING task_id;")
	rows := statement.QueryRow("read existing test task", "2018-01-01")
	rows.Scan(&task_id)
	task, err := Read(task_id)

	db.Exec("truncate table todo_table;")
	assert.NoError(t, err)
	assert.Equal(t, "read existing test task", task)
}

func TestReadForNoTask(t *testing.T) {
	os.Setenv("ENVIRONMENT", "test")
	config.Load()
	_, err := Read(-10000000)
	assert.EqualError(t, err, "Task Id is non-existent")
}

func TestShowAll(t *testing.T) {
	os.Setenv("ENVIRONMENT", "test")
	config.Load()
	err := ShowAll()
	assert.NoError(t, err)
}

func TestUpdate(t *testing.T) {
	os.Setenv("ENVIRONMENT", "test")
	config.Load()
	db := db.InitDB()
	var task_id int
	var task string
	statement, err := db.Prepare("INSERT INTO todo_table(task, timestamp) VALUES($1, $2) RETURNING task_id;")
	rows := statement.QueryRow("update test task", "2018-01-01")
	rows.Scan(&task_id)

	err = Update(task_id, "updated task")

	statement, err = db.Prepare("SELECT task from todo_table where task_id=$1;")
	row := statement.QueryRow(task_id)
	row.Scan(&task)
	assert.Equal(t, "updated task", task)
	assert.NoError(t, err)
	db.Exec("truncate table todo_table;")
}

func TestCannotUpdateWithEmptyTask(t *testing.T) {
	os.Setenv("ENVIRONMENT", "test")
	config.Load()
	db := db.InitDB()
	var task_id int
	statement, err := db.Prepare("INSERT INTO todo_table(task, timestamp) VALUES($1, $2) RETURNING task_id;")
	rows := statement.QueryRow("update test task", "2018-01-01")
	rows.Scan(&task_id)

	err = Update(task_id, "")
	assert.EqualError(t, err, "Cannot update with an empty task")
	db.Exec("truncate table todo_table;")
}

func TestDelete(t *testing.T) {
	os.Setenv("ENVIRONMENT", "test")
	config.Load()
	db := db.InitDB()
	var task_id int
	statement, err := db.Prepare("INSERT INTO todo_table(task, timestamp) VALUES($1, $2) RETURNING task_id;")
	rows := statement.QueryRow("delete test task", "2018-01-01")
	rows.Scan(&task_id)

	err = Delete(task_id)
	var counter int
	statement, err = db.Prepare("SELECT COUNT(*) from todo_table where task_id=$1;")
	rows = statement.QueryRow(task_id)
	rows.Scan(&counter)
	db.Exec("truncate table todo_table;")
	assert.Zero(t, counter)
	assert.NoError(t, err)
}
