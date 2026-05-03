package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"scannerbridge/scanner"
)

func init() {
	loadEnv()
}

func loadEnv() {
	content, err := os.ReadFile(".env")
	if err != nil {
		return
	}
	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			os.Setenv(strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1]))
		}
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "7878"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/health", handleHealth)
	mux.HandleFunc("/scanners", handleListScanners)
	mux.HandleFunc("/scan", handleScan)

	srv := &http.Server{
		Addr:         ":" + port,
		Handler:      corsMiddleware(mux),
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 120 * time.Second, // scan bisa butuh waktu lama
	}

	log.Printf("Scanner Bridge berjalan di port %s", port)
	log.Fatal(srv.ListenAndServe())
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v) //nolint:errcheck
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func handleListScanners(w http.ResponseWriter, r *http.Request) {
	scanners, err := scanner.ListScanners()
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	if scanners == nil {
		scanners = []scanner.ScannerInfo{}
	}
	writeJSON(w, http.StatusOK, scanners)
}

func handleScan(w http.ResponseWriter, r *http.Request) {
	var req scanner.ScanRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "request tidak valid"})
		return
	}

	if req.ScannerID == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "scanner_id wajib diisi"})
		return
	}
	if req.DPI == 0 {
		req.DPI = 300
	}
	if req.ColorMode == "" {
		req.ColorMode = "color"
	}
	if req.Format == "" {
		req.Format = "jpeg"
	}

	log.Printf("Scan dimulai: scanner=%s dpi=%d color=%s format=%s", req.ScannerID, req.DPI, req.ColorMode, req.Format)

	result, err := scanner.Scan(req)
	if err != nil {
		log.Printf("Scan gagal: %v", err)
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	log.Printf("Scan selesai: %dx%d %s", result.Width, result.Height, result.Format)
	writeJSON(w, http.StatusOK, result)
}
