package upload

import (
	"bytes"
	"fmt"
	"github.com/coentie/filesync/packages/paths"
	"github.com/fatih/color"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

func Upload() {
	contents, err := paths.GetContents()

	if err != nil {
		panic(err)
	}

	for _, path := range contents.Paths {
		handle(path)
	}
}

func handle(path string) {
	entries, err := os.ReadDir(path)

	if err != nil {
		panic(err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			handle(filepath.Join(path, entry.Name()))
			continue
		}

		doRequest(path, entry)
	}
}

func doRequest(path string, entry os.DirEntry) {
	qualifiedName := filepath.Join(path, entry.Name())
	file, err := os.Open(qualifiedName)

	if err != nil {
		color.Red("could not open file")
		return
	}

	defer file.Close()

	// Create a buffer to store the multipart form data
	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	color.Cyan(qualifiedName)

	part, err := writer.CreateFormFile("file", qualifiedName)

	if err != nil {
		color.Red("could not create form file")
		return
	}

	if _, err = io.Copy(part, file); err != nil {
		color.Red("could not copy file")
		return
	}

	writer.Close()
	req, err := http.NewRequest("POST", os.Getenv("API_URL"), &requestBody)

	if err != nil {
		color.Red("could not create upload request")
		return
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		color.Red("could not upload file", err)
		return
	}

	defer resp.Body.Close()

	fmt.Println("uploaded", entry.Name())
}
