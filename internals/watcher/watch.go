package watcher
// simple mental model ,
// linux will create event thru inotify , then go will read those events , print on terminal

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
)

func Watch(path string) error {
	w, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	defer w.Close()

	if err := addRecursive(w, path); err != nil {
		return err
	}

	fmt.Println()
	fmt.Println("VEIL   Watching workspace")
	fmt.Println()
	fmt.Println("Path:", path)
	fmt.Println("Press Ctrl+C to stop")
	fmt.Println()

	for {
		select {
		case event, ok := <-w.Events:
			if !ok {
				return nil
			}

			fmt.Println(event.Op, event.Name)

			if event.Has(fsnotify.Create) {
				info, err := os.Stat(event.Name)
				if err == nil && info.IsDir() {
					_ = addRecursive(w, event.Name)
				}
			}

		case err, ok := <-w.Errors:
			if !ok {
				return nil
			}

			fmt.Println("watch error:", err)
		}
	}
}

func addRecursive(w *fsnotify.Watcher, root string) error {
	return filepath.WalkDir(root, func(path string, entry os.DirEntry, err error) error {
		if err != nil {
			return nil
		}

		if entry.IsDir() {
			return w.Add(path)
		}

		return nil
	})
}