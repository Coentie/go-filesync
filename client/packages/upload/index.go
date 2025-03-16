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
		entries, err := os.ReadDir(path)

		if err != nil {
			panic(err)
		}

		for _, entry := range entries {
			if entry.IsDir() {
				//ToDo: make recursion so it uploads entire subfolders.
				//ToDo: should figure out server structure first.
				continue
			}

			color.Green("processing file: %s", entry.Name())

			file, err := os.Open(filepath.Join(path, entry.Name()))

			if err != nil {
				color.Red("could not open file")
				continue
			}

			defer file.Close()

			// Create a buffer to store the multipart form data
			var requestBody bytes.Buffer
			writer := multipart.NewWriter(&requestBody)

			part, err := writer.CreateFormFile("file", filepath.Base(entry.Name()))

			if err != nil {
				color.Red("could not create form file")
				continue
			}

			if _, err = io.Copy(part, file); err != nil {
				color.Red("could not copy file")
				continue
			}
			writer.Close()
			req, err := http.NewRequest("POST", os.Getenv("API_URL"), &requestBody)

			if err != nil {
				color.Red("could not create upload request")
				continue
			}

			req.Header.Set("Content-Type", writer.FormDataContentType())
			resp, err := http.DefaultClient.Do(req)

			if err != nil {
				color.Red("could not upload file", err)
				continue
			}

			fmt.Println(resp)

			defer resp.Body.Close()

			fmt.Println("uploaded", entry.Name())
		}
	}
}
