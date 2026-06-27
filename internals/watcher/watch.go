package watcher
// simple mental model ,
// linux will create event thru inotify , then go will read those events , print on terminal
import (
	"fmt"

	"github.com/fsnotify/fsnotify"
)

func Watch(path string) error {
	w, err := fsnotify.NewWatcher()  // we have used the inotify to track the events
	if err != nil {
		return err
	}
	defer w.Close()

	if err := w.Add(path); err != nil {
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

		case err, ok := <-w.Errors:
			if !ok {
				return nil
			}

			fmt.Println("watch error:", err)
		}
	}
}