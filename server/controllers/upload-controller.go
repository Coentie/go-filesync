package controllers

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func Upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in controller")
	// Limit the size of the request body (e.g., 10MB)
	r.ParseMultipartForm(10 << 20)

	// Retrieve the file from form-data
	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error retrieving file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Create a new file in the local storage
	dst, err := os.Create("./storage/" + handler.Filename)
	if err != nil {
		http.Error(w, "Error creating file", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	// Copy the uploaded file data to the new file
	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, "Error saving file", http.StatusInternalServerError)
		return
	}

	// Respond with success message
	fmt.Fprintf(w, "File uploaded successfully: %s", handler.Filename)
}
