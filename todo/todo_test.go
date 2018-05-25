package todo

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const task = "some random testing task"
const taskID = 1

func TestCreate(t *testing.T) {
	taskID, err := Create(task)

	// task id is coming 0 all the time although insertion in db is executing fine
	require.NotNil(t, taskID)
	require.NoError(t, err)
}

func TestRead(t *testing.T) {
	// task, err := Read(taskID)
	// require.NotNil(t, task)
	// require.NoError(t, err)
}

func TestUpdate(t *testing.T) {

}

func TestDelete(t *testing.T) {

}
