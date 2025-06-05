package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	setupRoutes()
	startServer()
}

func setupRoutes() {
	http.HandleFunc("/api/scan", handleScan)
	http.HandleFunc("/api/status", handleStatus)
}

func handleScan(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"status": "not implemented yet"}`)
}

func handleStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"status": "running"}`)
}

func startServer() {
	port := "8080"
	log.Printf("Server starting on port %s\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
