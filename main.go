package main

import (
	"blog/cloudinary"
	"log"
	"net/http"
	"strings"

	photosblog "blog/photos-blog"
)

func main() {
	initDependencies()
	setUpEndpoints()

	runServer()
}

func initDependencies() {
	cloudinary.Init()
}

func runServer() {
	log.Println("Server started at http://localhost:8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.Host, "localhost:") {
			w.Header().Set("Access-Control-Allow-Origin", "*")
		} else {
			w.Header().Set("Access-Control-Allow-Origin", "https://rkimmi.github.io")
		}

		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			return
		}

		next(w, r)
	}
}

func setUpEndpoints() {
	http.HandleFunc("/api/thumbnails", corsMiddleware(photosblog.GetThumbnailsHandler))
}
