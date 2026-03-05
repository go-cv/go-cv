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
	AppName    string   `yaml:"app_name"`
	Theme      string   `yaml:"theme"`
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
	s.AppName = "Go Template Container Web Server"
	s.Theme = "default"

	// Attempt to read the config file (try both config.yml and config.yaml)
	var configFile []byte
	var configPath string

	configFile, err := os.ReadFile("config.yml")
	if err == nil {
		configPath = "config.yml"
	} else {
		configFile, err = os.ReadFile("config.yaml")
		if err == nil {
			configPath = "config.yaml"
		}
	}

	if err != nil {
		if os.IsNotExist(err) {
			// Neither file exists, log and use default config
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
			fmt.Println("Error parsing config file:", configPath, err)
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
