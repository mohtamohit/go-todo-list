package todo

import (
	"fmt"
	"practice/go-todo-list/config"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	task         = "some random testing task"
	taskID       = 1
	updateTask   = "updated testing task"
	updateTaskID = 50
	deleteTaskID = 51
)

func TestCreate(t *testing.T) {
	config.Load()
	task_id, err := Create(task)
	fmt.Println(task_id)

	// facing a problem here. create function is always returning task id as zero
	// unable to fix this after lot of time too.

	require.NoError(t, err)
}

func TestReadForExistingTask(t *testing.T) {
	config.Load()
	task, err := Read(taskID)
	require.NoError(t, err)
	require.Equal(t, task, "some random testing task")
}

func TestReadForNoTask(t *testing.T) {
	config.Load()
	_, err := Read(10000000)
	require.EqualError(t, err, "Task Id is non-existent")
}

func TestShowAll(t *testing.T) {
	config.Load()
	err := ShowAll()
	require.NoError(t, err)
}

func TestUpdate(t *testing.T) {
	config.Load()
	err := Update(updateTaskID, updateTask)
	require.NoError(t, err)
}

func TestDelete(t *testing.T) {
	config.Load()
	err := Delete(deleteTaskID)
	require.NoError(t, err)
}
