package main

import (
	"fmt"
	"log"
	"net/http"

	"go-share/internal/config"
	"go-share/internal/entry"
	"go-share/internal/util"
)

const PORT = 8080

func main() {
	if err := config.LoadConfig("config.json"); err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	mux := entry.RegisterRoutes()
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	ip := util.GetLocalIP()
	log.Printf("Server started: http://%s:%d\n", ip, PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", PORT), mux))
}
