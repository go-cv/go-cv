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

func main() {
	// Create a channel to receive the OS signals
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM)

	// Initialize the WebService structure
	ws = new(WebServer)
	ws.Initialize()

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
