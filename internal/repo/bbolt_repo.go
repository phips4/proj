package repo

import (
	"encoding/json"
	"fmt"
	"github.com/phips4/proj/internal/model"
	"go.etcd.io/bbolt"
	"log"
	"os/exec"
	"strconv"
)

var ProjectBucket = []byte("projects")

type BBoltProjectRepository struct {
	db *bbolt.DB
}

func NewBBoltProjectRepository(db *bbolt.DB) *BBoltProjectRepository {
	return &BBoltProjectRepository{db: db}
}

// Add a new project to the repository.
func (r *BBoltProjectRepository) Add(name, description, path, execute string, labels []string) error {
	project := &model.Project{
		Name:        name,
		Description: description,
		Path:        path,
		Execute:     execute,
		Labels:      labels,
	}

	projectBytes, err := json.Marshal(project)
	if err != nil {
		return err
	}

	return r.db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket(ProjectBucket)
		id, _ := b.NextSequence()
		return b.Put(itob(int(id)), projectBytes)
	})
}

// UpdatePath updates the path of a project.
func (r *BBoltProjectRepository) UpdatePath(idOrName, path string) error {
	return r.updateField(idOrName, "Path", path)
}

// UpdateExecute updates the execute command of a project.
func (r *BBoltProjectRepository) UpdateExecute(idOrName, execute string) error {
	return r.updateField(idOrName, "Execute", execute)
}

// UpdateDescription updates the description of a project.
func (r *BBoltProjectRepository) UpdateDescription(idOrName, description string) error {
	return r.updateField(idOrName, "Description", description)
}

// AddLabel adds a label to a project.
func (r *BBoltProjectRepository) AddLabel(idOrName, label string) error {
	return r.updateLabel(idOrName, label, true)
}

// RemoveLabel removes a label from a project.
func (r *BBoltProjectRepository) RemoveLabel(idOrName, label string) error {
	return r.updateLabel(idOrName, label, false)
}

// Delete a project from the repository.
func (r *BBoltProjectRepository) Delete(idOrName string) error {
	return r.db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket(ProjectBucket)
		return b.ForEach(func(k, v []byte) error {
			project := &model.Project{}
			err := json.Unmarshal(v, &project)
			if err != nil {
				return err
			}
			if project.Name == idOrName || string(k) == idOrName {
				return b.Delete(k)
			}
			return nil
		})
	})
}

// Get retrieves a project by its name or ID.
func (r *BBoltProjectRepository) Get(idOrName string) (*model.Project, error) {
	var result *model.Project

	err := r.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket(ProjectBucket)
		return b.ForEach(func(k, v []byte) error {
			project := &model.Project{}
			err := json.Unmarshal(v, &project)
			if err != nil {
				return err
			}
			if project.Name == idOrName || string(k) == idOrName {
				result = project
				return nil
			}
			return nil
		})
	})

	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, nil // Project not found
	}

	return result, nil
}

// Execute runs the execute command of a project.
func (r *BBoltProjectRepository) Execute(idOrName string) error {
	var execute string

	err := r.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket(ProjectBucket)
		return b.ForEach(func(k, v []byte) error {
			project := &model.Project{}
			err := json.Unmarshal(v, &project)
			if err != nil {
				return err
			}
			if project.Name == idOrName || string(k) == idOrName {
				execute = project.Execute
				return nil
			}
			return nil
		})
	})

	if err != nil {
		return err
	}

	if execute == "" {
		return fmt.Errorf("no execute command found for project %s", idOrName)
	}

	cmdExec := exec.Command(execute)
	out, err := cmdExec.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error executing command: %s", err)
	}
	log.Println(string(out))

	return nil
}

// Close closes the BBolt database.
func (r *BBoltProjectRepository) Close() {
	if r.db != nil {
		r.db.Close()
	}
}

func (r *BBoltProjectRepository) updateField(idOrName, field, value string) error {
	return r.db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket(ProjectBucket)
		return b.ForEach(func(k, v []byte) error {
			project := &model.Project{}
			err := json.Unmarshal(v, &project)
			if err != nil {
				return err
			}
			if project.Name == idOrName || string(k) == idOrName {
				switch field {
				case "Path":
					project.Path = value
				case "Execute":
					project.Execute = value
				case "Description":
					project.Description = value
				}
				projectBytes, err := json.Marshal(project)
				if err != nil {
					return err
				}
				return b.Put(k, projectBytes)
			}
			return nil
		})
	})
}

func (r *BBoltProjectRepository) updateLabel(idOrName, label string, add bool) error {
	return r.db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket(ProjectBucket)
		return b.ForEach(func(k, v []byte) error {
			project := &model.Project{}
			err := json.Unmarshal(v, &project)
			if err != nil {
				return err
			}

			if project.Name != idOrName && string(k) != idOrName {
				return nil
			}

			if add {
				project.Labels = append(project.Labels, label)
			} else {
				project.Labels = remove(project.Labels, label)
			}

			projectBytes, err := json.Marshal(project)
			if err != nil {
				return err
			}

			return b.Put(k, projectBytes)
		})
	})
}

func (r *BBoltProjectRepository) All() ([]*model.Project, error) {
	var projects []*model.Project
	err := r.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket(ProjectBucket)

		return b.ForEach(func(k, v []byte) error {
			project := &model.Project{}
			err := json.Unmarshal(v, &project)
			if err != nil {
				return err
			}

			projects = append(projects, project)

			return nil
		})
	})
	if err != nil {
		log.Println("error listing projects:", err)
		return nil, err
	}
	return projects, nil
}

// Helper function to remove a label from the slice
func remove(labels []string, label string) []string {
	for i, l := range labels {
		if l == label {
			return append(labels[:i], labels[i+1:]...)
		}
	}
	return labels
}

// Convert an int to a []byte
func itob(v int) []byte {
	return []byte(strconv.Itoa(v))
}
