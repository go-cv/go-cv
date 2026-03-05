package main

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
)

// startWatcher starts watching the content directory for changes
func startWatcher() (*fsnotify.Watcher, error) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, fmt.Errorf("failed to create watcher: %w", err)
	}

	// Add content directory to watcher
	err = watcher.Add(contentPath)
	if err != nil {
		watcher.Close()
		return nil, fmt.Errorf("failed to watch content directory: %w", err)
	}

	fmt.Printf("Watching %s for changes...\n", contentPath)

	// Start goroutine to handle events
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				// Only process write and create events for .md files
				if event.Has(fsnotify.Write) || event.Has(fsnotify.Create) {
					fmt.Printf("File changed: %s\n", event.Name)
					// Regenerate output
					if err := generateOutput(); err != nil {
						fmt.Printf("Error regenerating: %s\n", err)
					}
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				fmt.Printf("Watcher error: %s\n", err)
			}
		}
	}()

	return watcher, nil
}
