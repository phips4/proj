package model

type Project struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Path        string   `json:"path"`
	Execute     string   `json:"execute"`
	Labels      []string `json:"labels"`
}
