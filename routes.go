package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (s *WebServer) Routes(r *mux.Router) {
	r.HandleFunc("/version", handleVersion).Methods("GET")

	// Serve generated HTML files from output directory
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(outputPath))))
}
