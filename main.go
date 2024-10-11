package main

import (
    "blog/photos-blog/cloudinary"
    "net/http"
)

func main() {
    cloudinary.Init()
    setUpEndpoints()
}

func setUpEndpoints() {
    http.HandleFunc("/api/thumbnails", cloudinary.GetThumbnailsHandler)
}
