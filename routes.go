package main

import "github.com/gorilla/mux"

func (s *WebServer) Routes(r *mux.Router) {
	r.HandleFunc("/version", handleVersion).Methods("GET")
}
