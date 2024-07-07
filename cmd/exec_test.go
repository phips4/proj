package cmd

import (
	"github.com/phips4/proj/internal/model"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

type mockedGetter struct {
	executed bool
}

func (getter *mockedGetter) Get(idOrName string) (*model.Project, error) {
	getter.executed = true
	return &model.Project{Name: "mockedProject"}, nil
}

func (getter *mockedGetter) All() ([]*model.Project, error) {
	return nil, nil
}

func Test_runExecFunc(t *testing.T) {
	getter := &mockedGetter{}
	cmd := &cobra.Command{
		Run: runExecFunc(getter),
	}

	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"test", "test_project"}

	err := cmd.Execute()

	assert.NoError(t, err)
	assert.True(t, getter.executed, "Get method should be called")
}
