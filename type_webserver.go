package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"gopkg.in/yaml.v2"
)

type WebServer struct {
	HTTPServer *http.Server
	Listen     WSListen `yaml:"listen"`
}

type WSListen struct {
	Address string `yaml:"address"`
	Port    string `yaml:"port"`
}

func (s *WebServer) Initialize() {
	// Initialize default values
	s.Listen = WSListen{
		Address: "0.0.0.0",
		Port:    "80",
	}

	// Attempt to read the config file
	configFile, err := os.ReadFile("config.yml")
	if err != nil {
		if os.IsNotExist(err) {
			// File does not exist, log and use default config
			fmt.Println("Config file not found, using default settings.")
		} else {
			// Some other error occurred when trying to read the file, exit
			fmt.Println("Error reading config file:", err)
			os.Exit(1)
		}
	} else {
		// If the file exists, unmarshal it into the ServiceSettings struct
		err = yaml.Unmarshal(configFile, &s)
		if err != nil {
			fmt.Println("Error parsing config file:", err)
			os.Exit(1)
		}
	}
}

func (s *WebServer) Start() error {
	// Create a new MUX router and an HTTP server
	r := mux.NewRouter()
	s.HTTPServer = &http.Server{
		Addr:    s.Listen.Address + ":" + s.Listen.Port,
		Handler: r,
	}

	// Associate the various handlers (routes)
	s.Routes(r)

	// Start the server
	fmt.Println("Listening on", s.Listen.Address+":"+s.Listen.Port)
	err := s.HTTPServer.ListenAndServe()

	// Return error, or nil
	return err
}
