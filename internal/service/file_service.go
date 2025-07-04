package service

import (
	"errors"
	"path/filepath"

	"go-share/internal/config"
	"go-share/internal/repo"
)

func ListFiles(folder string) ([]string, error) {
	// Just check if folder exists in config, else error
	if _, ok := config.AppConfig.Folders[folder]; !ok {
		return nil, errors.New("invalid folder")
	}
	// Use repo function to read files
	return repo.ReadDirFiles(folder)
}

func GetFilePath(folder, filename string) (string, error) {
	basePath, ok := config.AppConfig.Folders[folder]
	if !ok {
		return "", errors.New("invalid folder")
	}
	return filepath.Join(basePath, filename), nil
}
