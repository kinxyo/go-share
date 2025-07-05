package entry

import (
	"net/http"

	"go-share/internal/api"
	"go-share/internal/middleware"
)

func RegisterRoutes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", api.HandleHome)
	mux.HandleFunc("/api/files", api.HandleFileList)
	mux.HandleFunc("/api/folders", api.HandleFolderList) // NEW endpoint for folders
	mux.HandleFunc("/files/", api.HandleFileServe)

	// Wrap with logging middleware
	return middleware.Logging(mux)
}
