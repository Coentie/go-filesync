package controllers

import (
	"encoding/json"
	"github.com/coentie/filesync-server/packages/services"
	"net/http"
	"os"
)

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
