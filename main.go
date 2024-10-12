package main

import (
    "blog/cloudinary"
    "blog/photos-blog"
    "net/http"
    "log"
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

func setUpEndpoints() {
    http.HandleFunc("/api/thumbnails", photosblog.GetThumbnailsHandler)
}
