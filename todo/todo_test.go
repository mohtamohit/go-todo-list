package todo

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	task         = "some random testing task"
	taskID       = 48
	updateTask   = "updated testing task"
	updateTaskID = 50
	deleteTaskID = 51
)

func TestCreate(t *testing.T) {
	taskID, err := Create(task)
	fmt.Println("b ", taskID)
	require.NotNil(t, taskID)
	require.NoError(t, err)
}

func TestReadForExistingTask(t *testing.T) {
	task, err := Read(taskID)
	require.Equal(t, task, "some random testing task")
	require.NoError(t, err)
}

func TestReadForNoTask(t *testing.T) {
	_, err := Read(10000000)
	require.EqualError(t, err, "Task Id is non-existent")
}

func TestShowAll(t *testing.T) {
	err := ShowAll()
	require.NoError(t, err)
}

func TestUpdate(t *testing.T) {
	err := Update(updateTaskID, updateTask)
	require.NoError(t, err)
}

func TestDelete(t *testing.T) {
	err := Delete(deleteTaskID)
	require.NoError(t, err)
}
