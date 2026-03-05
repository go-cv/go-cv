package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var APP_VERSION string = "latest"
var COMMIT_ID string = "undefined"
var ws *WebServer

const (
	contentPath = "./content"
	outputPath  = "./output"
)

func main() {
	// Initialize the WebService structure
	ws = new(WebServer)
	ws.Initialize()

	// Determine mode based on command line argument
	// No CLI flags - just check os.Args
	if len(os.Args) > 1 && os.Args[1] == "serve" {
		// Serve mode: Start HTTP server with file watching
		runServeMode()
	} else {
		// CLI mode: Generate static files and exit
		runCLIMode()
	}
}

func runCLIMode() {
	fmt.Println("Running in CLI mode...")
	fmt.Printf("Reading content from: %s\n", contentPath)
	fmt.Printf("Writing output to: %s\n", outputPath)

	// Read content and generate output
	err := generateOutput()
	if err != nil {
		fmt.Printf("Error generating output: %s\n", err)
		os.Exit(1)
	}

	fmt.Println("Generation complete.")
}

func runServeMode() {
	fmt.Println("Running in Serve mode...")

	// Create a channel to receive the OS signals
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM)

	// Start the WebService in a separate goroutine
	go ws.Start()

	// Wait for a signal
	<-sc
	fmt.Println("Shutting down...")

	// Create a context with a timeout for graceful shutdown
	shCtx, shCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shCancel()

	// Shutdown the HTTP server
	err := ws.HTTPServer.Shutdown(shCtx)
	if err != nil {
		fmt.Printf("Server shutdown error: %s", err)
		os.Exit(1)
	}
}
