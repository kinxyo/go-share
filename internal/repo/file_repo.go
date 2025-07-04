package repo

import (
	"errors"
	"go-share/internal/config"
	"os"
)

func ReadDirFiles(folderKey string) ([]string, error) {
	path, ok := config.AppConfig.Folders[folderKey]
	if !ok {
		return nil, errors.New("invalid folder")
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var files []string
	for _, e := range entries {
		if !e.IsDir() {
			files = append(files, e.Name())
		}
	}
	return files, nil
}
