package repo

import (
	"github.com/phips4/proj/internal/model"
)

type ProjectAdder interface {
	Add(name, description, path, execute string, labels []string) error
}

type ProjectPathUpdater interface {
	UpdatePath(idOrName, path string) error
}

type ProjectExecuteUpdater interface {
	UpdateExecute(idOrName, execute string) error
}

type ProjectDescriptionUpdater interface {
	UpdateDescription(idOrName, description string) error
}

type ProjectLabelManager interface {
	AddLabel(idOrName, label string) error
	RemoveLabel(idOrName, label string) error
}

type ProjectDeleter interface {
	Delete(idOrName string) error
}

type ProjectGetter interface {
	Get(idOrName string) (*model.Project, error)
	All() ([]*model.Project, error)
}

type ProjectExecutor interface {
	Execute(idOrName string) error
}

type ProjectCloser interface {
	Close()
}

type ProjectRepo interface {
	ProjectAdder
	ProjectPathUpdater
	ProjectDescriptionUpdater
	ProjectLabelManager
	ProjectDeleter
	ProjectGetter
	ProjectExecutor
	ProjectCloser
}
