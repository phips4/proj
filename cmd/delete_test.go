package cmd

import (
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

type mockedDeleter struct {
	deleted bool
}

func (deleter *mockedDeleter) Delete(idOrName string) error {
	deleter.deleted = true
	return nil
}

func Test_runDeleteFunc(t *testing.T) {
	mockedDeleter := &mockedDeleter{}
	cmd := &cobra.Command{
		RunE: runDeleteFunc(mockedDeleter),
	}

	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"test", "test_project"}

	err := cmd.Execute()

	assert.NoError(t, err)
	assert.True(t, mockedDeleter.deleted, "Delete method should be called")
}

func Test_runDeleteFunc_empty_args(t *testing.T) {
	mockedDeleter := &mockedDeleter{}
	cmd := &cobra.Command{
		RunE: runDeleteFunc(mockedDeleter),
		Args: cobra.ExactArgs(1),
	}

	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"test"}

	err := cmd.Execute()

	assert.Error(t, err)
	assert.Equal(t, false, mockedDeleter.deleted, "delete method did get called")
}
