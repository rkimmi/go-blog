package photosblog

import (
	"encoding/json"
	"log"
	"net/http"
	"reflect"

	"blog/cloudinary"
	imagemodels "blog/photos-blog/models"
)

func GetThumbnailsHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received %s request for %s", r.Method, r.URL.Path)
	// Extract the 'folder' parameter from the URL query string
	folder := r.URL.Query().Get("folder")

	if folder == "" {
		// If no folder is provided, return a bad request response
		http.Error(w, "Missing folder parameter", http.StatusBadRequest)
		return
	}

	// Call the function to get all images in the specified folder
	images, err := cloudinary.GetAllImagesInFolder(folder)
	if err != nil {
		http.Error(w, "Error retrieving images: %v", http.StatusInternalServerError)

		return
	}

	thumbnails := []imagemodels.ImageThumbnail{}

	for _, resource := range images {
		log.Printf("Resource: %+v\n", resource)

		// Print the type of the resource
		log.Printf("Type of resource: %v\n", reflect.TypeOf(resource))

		thumbnails = append(thumbnails, imagemodels.ImageThumbnail{
			ID:  resource.PublicID,
			URL: resource.SecureURL,
		})
	}

	response := imagemodels.GetThumbnailsResponse{
		Images: thumbnails,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		log.Printf("Error encoding response: %v", err)
		return
	}
}
