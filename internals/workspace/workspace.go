package workspace

// here we started mking our first layer which workspace layer
// this layer will talk to the overlay layer 

import (
	"os"
	"path/filepath"
	"veil/internals/overlay"
		"encoding/json"
		"veil/internals/fs"
		"veil/internals/watcher"  

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

func Load(name string) (*Workspace, error) {  // this func will read the path from the config.json then unmarshal it afetr that it wi will create the workspace object 
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	configPath := filepath.Join(
		home,
		".veil",
		"workspaces",
		name,
		"config.json",
	)

	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var ws Workspace

	if err := json.Unmarshal(data, &ws); err != nil {
		return nil, err
	}

	return &ws, nil
}

func (w *Workspace) Unmount() error {
	return overlay.Unmount(w.Merged)
}

func (w *Workspace) Destroy() error {
	_ = w.Unmount()

	return os.RemoveAll(w.Path)
} // this wil try to unmount first but if already unmounted then delete the directory


func (w *Workspace) Apply() error {
	return filepath.WalkDir(w.Upper, func(path string, entry os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if entry.IsDir() {
			return nil
		}

		relPath, err := filepath.Rel(w.Upper, path)
		if err != nil {
			return err
		}

		destPath := filepath.Join(w.Project, relPath)

		return fs.CopyFile(path, destPath)
	})
}


func List() ([]Workspace, error) {  // this will list all the workspaces user have created
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	base := filepath.Join(home, ".veil", "workspaces")

	entries, err := os.ReadDir(base)
	if err != nil {
		if os.IsNotExist(err) {
			return []Workspace{}, nil
		}
		return nil, err
	}

	workspaces := []Workspace{}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		ws, err := Load(entry.Name())
		if err != nil {
			continue
		}

		workspaces = append(workspaces, *ws)
	}

	return workspaces, nil
}



func FindByProject(project string) (*Workspace, error) { // this will help to detec the duplicate project 
	absProject, err := filepath.Abs(project)
	if err != nil {
		return nil, err
	}

	workspaces, err := List()
	if err != nil {
		return nil, err
	}

	for _, ws := range workspaces {
		if ws.Project == absProject {
			return &ws, nil
		}
	}

	return nil, nil
}


func (w *Workspace) IsMounted() bool {
	return overlay.IsMounted(w.Merged)
}


func (w *Workspace) Watch() error {
	return watcher.Watch(w.Merged)
}