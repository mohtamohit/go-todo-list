package todo

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const task = "some random testing task"

func TestCreate(t *testing.T) {
	taskID, err := Create(task)
	require.NotNil(t, taskID)
	require.NoError(t, err)
}

func TestRead(t *testing.T) {

}

func TestUpdate(t *testing.T) {

}

func TestDelete(t *testing.T) {

}
