package api

import (
	"encoding/json"
	"net/http"
	// "path/filepath"
	"strings"

	"go-share/internal/config"
	"go-share/internal/service"
)

func HandleHome(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/index.html")
}

func HandleFileList(w http.ResponseWriter, r *http.Request) {
	folder := r.URL.Query().Get("folder")
	files, err := service.ListFiles(folder)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(files)
}

func HandleFileServe(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/files/")
	parts := strings.SplitN(path, "/", 2)
	if len(parts) < 2 {
		http.Error(w, "Invalid file path", http.StatusBadRequest)
		return
	}
	folder, file := parts[0], parts[1]
	fullPath, err := service.GetFilePath(folder, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	http.ServeFile(w, r, fullPath)
}

func HandleFolderList(w http.ResponseWriter, r *http.Request) {
	folders := make([]string, 0, len(config.AppConfig.Folders))
	for k := range config.AppConfig.Folders {
		folders = append(folders, k)
	}
	json.NewEncoder(w).Encode(folders)
}
