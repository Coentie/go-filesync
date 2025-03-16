package download

import (
	"encoding/base64"
	"encoding/json"
	"github.com/fatih/color"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func Download() {
	res, err := http.Get(os.Getenv("API_URL") + "/files")

	if err != nil || res == nil {
		color.Red("could not fetch file list")
	}

	defer res.Body.Close()

	bodyBytes, err := io.ReadAll(res.Body)

	if err != nil {
		color.Red("could not read response body")
	}

	var files []string
	err = json.Unmarshal(bodyBytes, &files)

	if err != nil {
		color.Red("could not read json from response body")
	}

	for _, file := range files {
		qualifiedPath := os.Getenv("SYNC_PATH")
		elements := strings.Split(file, "/")

		for _, element := range elements {

			// If file has extension we are at the end and we should download the file.
			if strings.Contains(element, ".") {
				filestring := base64.StdEncoding.EncodeToString([]byte(file))

				out, err := os.Create(filepath.Join(os.Getenv("SYNC_PATH"), file))

				if err != nil {
					color.Red("could not create file %s", filepath.Join(qualifiedPath, file))
					continue
				}

				defer out.Close()

				resp, _ := http.Get(os.Getenv("API_URL") + "/download/" + filestring)

				defer resp.Body.Close()

				_, err = io.Copy(out, resp.Body)

				if err != nil {
					color.Red("could not copy body %s", file)
				}

				continue
			}

			qualifiedPath = filepath.Join(qualifiedPath, element)
			color.Yellow(qualifiedPath)
			if _, err := os.Stat(qualifiedPath); os.IsNotExist(err) {
				err = os.Mkdir(qualifiedPath, os.ModePerm)

				if err != nil {
					color.Red(err.Error())
					break
				}
			}
		}
	}
}
