package cmd

import (
	"github.com/phips4/proj/internal/model"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

type mockedAllGetter struct {
	listed bool
}

func (getter *mockedAllGetter) Get(string) (*model.Project, error) {
	return nil, nil // this method is not used at all in this test
}

func (getter *mockedAllGetter) All() ([]*model.Project, error) {
	getter.listed = true
	return nil, nil
}

func Test_runListFunc(t *testing.T) {
	getter := &mockedAllGetter{}
	cmd := &cobra.Command{
		Run: runListFunc(getter),
	}

	err := cmd.Execute()

	assert.NoError(t, err)
	assert.True(t, getter.listed, "All method should be called")
}

func Test_runListFunc_empty_args(t *testing.T) {
	getter := &mockedAllGetter{}
	cmd := &cobra.Command{
		Run:  runListFunc(getter),
		Args: cobra.ExactArgs(0),
	}

	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"mycliapp_test", "arg1", "arg2"}

	err := cmd.Execute()

	assert.Error(t, err)
	assert.False(t, getter.listed, "All method should not be called")
}
