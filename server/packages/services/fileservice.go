package services

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func ListStorageFiles(basePath string) ([]string, error) {
	paths := []string{}

	paths, err := listFiles(basePath, paths, basePath)

	if err != nil {
		return paths, err
	}

	return paths, nil
}

func listFiles(basePath string, paths []string, pathFilter string) ([]string, error) {
	entries, err := os.ReadDir(basePath)

	fmt.Println(basePath)

	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			paths, _ = listFiles(filepath.Join(basePath, entry.Name()), paths, pathFilter)
		} else {
			name := filepath.Join(basePath, entry.Name())
			name = strings.Replace(name, pathFilter, "", 1)
			paths = append(paths, name)
		}
	}

	return paths, nil
}
