package upload

import (
	"bytes"
	"fmt"
	"github.com/coentie/filesync/packages/paths"
	"net/http"
	"os"
	"path/filepath"
)

func Upload() {
	fmt.Println("in upload")
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

			fmt.Println(entry.Name())

			contents, err := os.ReadFile(filepath.Join(path, entry.Name()))

			if err != nil {
				fmt.Println("could not read file", path)
				continue
			}

			_, err = http.Post("http://127.0.0.1:5000", "multipart/form-data", bytes.NewBuffer(contents))

			if err != nil {
				fmt.Println("could not upload file", err)
			}
			fmt.Println("uploaded", entry.Name())
		}
	}
}
