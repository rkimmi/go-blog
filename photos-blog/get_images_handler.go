package photosblog

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	cloudinary "blog/cloudinary"
	imagemodels "blog/photos-blog/models"
)

func GetThumbnailsHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received %s request for %s", r.Method, r.URL.Path)
	queryParams := r.URL.Query()

	folder := queryParams.Get("folder")
	if folder == "" {
		folder = "photos-blog"
	}

	limitStr := queryParams.Get("limit")
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 20
	}

	nextImageStartKey := queryParams.Get("startKey")
	resources, err := cloudinary.GetAllImagesInFolder(limit, nextImageStartKey, folder)

	if err != nil {
		http.Error(w, "Error retrieving images: %v", http.StatusInternalServerError)

		return
	}

	thumbnails := []imagemodels.ImageThumbnail{}

	for _, resource := range resources.Assets {

		thumbnails = append(thumbnails, imagemodels.ImageThumbnail{
			ID:  resource.PublicID,
			URL: resource.SecureURL,
		})
	}

	response := imagemodels.GetThumbnailsResponse{
		Images:         thumbnails,
		NextImageStart: resources.NextCursor,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		log.Printf("Error encoding response: %v", err)
		return
	}
}
