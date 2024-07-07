package cmd

import (
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

type mockedAdder struct {
	added bool
}

func (adder *mockedAdder) Add(name string, description string, path string, execute string, labels []string) error {
	adder.added = true
	return nil
}

func Test_runFunc(t *testing.T) {
	adder := &mockedAdder{}
	cmd := &cobra.Command{
		Run: runAddFunc(adder),
	}

	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"test", "test_project", "test_description"}

	err := cmd.Execute()

	assert.NoError(t, err)
	assert.True(t, adder.added, "add method did not get called")
}

func Test_runFunc_empty_args(t *testing.T) {
	adder := &mockedAdder{}
	cmd := &cobra.Command{
		Run:  runAddFunc(adder),
		Args: cobra.ExactArgs(2),
	}

	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"test", "test_project"}

	err := cmd.Execute()

	assert.Error(t, err)
	assert.False(t, adder.added, "add method did get called")
}
