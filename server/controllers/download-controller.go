package controllers

import (
	"encoding/base64"
	"encoding/json"
	"github.com/coentie/filesync-server/packages/services"
	"github.com/go-chi/chi/v5"
	"net/http"
	"os"
	"path/filepath"
)

// List files for download.
func Files(w http.ResponseWriter, r *http.Request) {
	files, err := services.ListStorageFiles(os.Getenv("STORAGE_PATH"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	content, err := json.Marshal(files)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(content)
}

// Download files.
func Download(w http.ResponseWriter, r *http.Request) {
	base64FileName := chi.URLParam(r, "filename")
	filename, err := base64.StdEncoding.DecodeString(base64FileName)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	file, err := os.ReadFile(filepath.Join(os.Getenv("STORAGE_PATH"), string(filename)))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(file)
}
