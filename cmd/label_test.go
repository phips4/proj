package cmd

import (
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

type mockedLabelManager struct {
	added   bool
	removed bool
}

func (l *mockedLabelManager) AddLabel(idOrName, label string) error {
	l.added = true
	return nil
}

func (l *mockedLabelManager) RemoveLabel(idOrName, label string) error {
	l.removed = true
	return nil
}

func Test_runAddLabelFunc(t *testing.T) {
	label := &mockedLabelManager{}
	cmd := &cobra.Command{
		RunE: runAddLabelFunc(label),
		Args: cobra.ExactArgs(2),
	}

	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"test", "test_project", "test_label"}

	err := cmd.Execute()

	assert.NoError(t, err)
	assert.True(t, label.added, "AddLabel method should be called")
	assert.False(t, label.removed, "RemoveLabel method should not be called")
}

func Test_runRemoveLabel(t *testing.T) {
	label := &mockedLabelManager{}
	cmd := &cobra.Command{
		RunE: runRemoveLabelFunc(label),
		Args: cobra.ExactArgs(2),
	}

	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"test", "test_project", "test_label"}

	err := cmd.Execute()

	assert.NoError(t, err)
	assert.False(t, label.added, "AddLabel method should not be called")
	assert.True(t, label.removed, "RemoveLabel method should not be called")
}
