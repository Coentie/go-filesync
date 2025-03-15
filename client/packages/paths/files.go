package paths

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

var FILENAME = "paths.json"

func CreatePathsFile() bool {
	if _, err := os.Stat(getConfigFilePath()); os.IsNotExist(err) {
		_, err = os.Create(getConfigFilePath())

		if err != nil {
			return false
		}
		err = WriteContents(NewContent())

		if err != nil {
			return false
		}
	}

	return true
}

// GetContents Fetches the content of the paths.
func GetContents() (CONTENT, error) {
	contents, err := os.ReadFile(getConfigFilePath())
	var content CONTENT

	if err != nil {
		fmt.Println("Error reading config file: ", err)
		return content, err
	}

	err = json.Unmarshal(contents, &content)

	if err != nil {
		return content, handleInvalidJson()
	}

	return content, nil
}

// WriteContents Writing contents to the fall.
func WriteContents(content CONTENT) error {
	file, err := os.Create(getConfigFilePath())

	if err != nil {
		return err
	}

	defer file.Close()

	results, err := json.Marshal(content)

	if err != nil {
		return err
	}

	_, err = file.Write(results)

	if err != nil {
		return err
	}

	return nil
}

// Handle invalid JSON.
func handleInvalidJson() error {
	var res string
	fmt.Println("Error reading JSON. Would you like to create new content? y/n")
	_, err := fmt.Scan(&res)

	if err != nil {
		return err
	}

	if res == "n" || res == "no" {
		return nil
	}

	return WriteContents(NewContent())
}

// Fetches the config file path.
func getConfigFilePath() string {
	return filepath.Join(PATHNAME, FILENAME)
}
