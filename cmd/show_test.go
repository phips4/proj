package cmd

import (
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_runShowFunc(t *testing.T) {
	getter := &mockedGetter{}
	cmd := &cobra.Command{
		RunE: runShowFunc(getter),
		Args: cobra.ExactArgs(1),
	}

	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"test", "test_project"}

	err := cmd.Execute()

	assert.NoError(t, err)
	assert.True(t, getter.executed, "Get method should be called")
}
