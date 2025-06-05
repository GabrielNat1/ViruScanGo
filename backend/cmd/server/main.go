package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/GabrielNat1/ViruScanGo/internal/config"
	"github.com/GabrielNat1/ViruScanGo/internal/scanner"
)

var (
	cfg            *config.Config
	scannerService scanner.Scanner
)

func init() {
	cfg = config.NewDefaultConfig()
	scannerService = scanner.NewScanner()
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

	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error reading file: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	tempFile, err := os.CreateTemp("", "scan-*")
	if err != nil {
		http.Error(w, "Error creating temp file: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer os.Remove(tempFile.Name())
	defer tempFile.Close()

	_, err = io.Copy(tempFile, file)
	if err != nil {
		http.Error(w, "Error saving file: "+err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := scannerService.ScanFile(r.Context(), tempFile.Name())
	if err != nil {
		http.Error(w, "Error scanning file: "+err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"filename":   header.Filename,
		"infected":   result.IsInfected,
		"threatName": result.ThreatName,
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
