package entry

import (
	"net/http"

	"go-share/internal/api"
)

func RegisterRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", api.HandleHome)
	mux.HandleFunc("/api/files", api.HandleFileList)
	mux.HandleFunc("/api/folders", api.HandleFolderList) // NEW endpoint for folders
	mux.HandleFunc("/files/", api.HandleFileServe)

	return mux
}
