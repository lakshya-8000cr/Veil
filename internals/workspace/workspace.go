package workspace

// here we started mking our first layer which workspace layer
// this layer will talk to the overlay layer 

import (
	"os"
	"path/filepath"
		"encoding/json"

)

type Workspace struct {
	Name    string
	Project string
	Path    string
	Upper   string
	Work    string
	Merged  string
}

func New(name string, project string) (*Workspace, error) { //  this will create the dirs and also for nwo it is not ma,ing dir these are just the objects
	absProject, err := filepath.Abs(project)
	if err != nil {
		return nil, err
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	base := filepath.Join(home, ".veil", "workspaces", name)

	return &Workspace{
		Name:    name,
		Project: absProject,
		Path:    base,
		Upper:   filepath.Join(base, "upper"),
		Work:    filepath.Join(base, "work"),
		Merged:  filepath.Join(base, "merged"),
	}, nil
}

func (w *Workspace) Create() error {  // this will create the filepath /veil/workspace/project-name/ upper  ./work  ./merged  ./config.json
	if err := os.MkdirAll(w.Upper, 0755); err != nil {
		return err
	}

	// config.json will contain the meta data of the project , maimly name and the path  

	if err := os.MkdirAll(w.Work, 0755); err != nil {
		return err
	}

	if err := os.MkdirAll(w.Merged, 0755); err != nil {
		return err
	}

	configPath := filepath.Join(w.Path, "config.json")

	data, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(configPath, data, 0644)
}