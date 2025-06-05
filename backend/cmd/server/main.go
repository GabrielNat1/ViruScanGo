package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/GabrielNat1/ViruScanGo/internal/config"
	"github.com/GabrielNat1/ViruScanGo/internal/scanner"
)

var (
	cfg            *config.Config
	scannerService scanner.Scanner
)

func init() {
	cfg = config.NewDefaultConfig()
}

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

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"status": "scan request received",
	})
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
