package cloudinary

import (
    "context"
    "fmt"
    "log"
    "net/http"
    "encoding/json"
    "reflect"

    // cloudinarySDK "github.com/cloudinary/cloudinary-go/v2"
    "github.com/cloudinary/cloudinary-go/v2/api/admin" 

    "blog/photos-blog/cloudinary/models" 
)

func GetThumbnailsHandler(w http.ResponseWriter, r *http.Request) {
    // Extract the 'folder' parameter from the URL query string
    folder := r.URL.Query().Get("folder")
        
    if folder == "" {
        // If no folder is provided, return a bad request response
        http.Error(w, "Missing folder parameter", http.StatusBadRequest)
        return
    }

    // Call the function to get all images in the specified folder
    images, err := GetAllImagesInFolder(folder)
    if err != nil {
        http.Error(w, fmt.Sprintf("Error retrieving images: %v", err), http.StatusInternalServerError)
        return
    }

    // Set the response content type to JSON
    w.Header().Set("Content-Type", "application/json")

    // Return the list of images as JSON
    if err := json.NewEncoder(w).Encode(images); err != nil {
        http.Error(w, "Failed to encode response", http.StatusInternalServerError)
        return
    }
}

// TODO get thumbnail urls rather than full size
func GetAllImagesInFolder(folderName string) (cloudinarymodels.GetThumbnailsResponse, error) {
    fmt.Println("Getting images in folder", folderName)

    cld := GetCloudinary()

    // Create a context
    ctx := context.Background()
    // TODO get by folder name!
    resources, err := cld.Admin.Assets(ctx, admin.AssetsParams{})

    if len(resources.Assets) == 0 {
        fmt.Println("No images found in the folder:", folderName)
    }
    
    fmt.Printf("Total assets found: %d\n", len(resources.Assets))

    if err != nil {
        log.Fatalf("Failed to list resources: %v", err)
    }

    for _, asset := range resources.Assets {
        fmt.Println("Image URL:", asset.SecureURL)
    }

    thumbnails := []cloudinarymodels.ImageThumbnail{}

    for _, resource := range resources.Assets {
        fmt.Printf("Resource: %+v\n", resource)

        // Print the type of the resource
        fmt.Printf("Type of resource: %v\n", reflect.TypeOf(resource))
        
        thumbnails = append(thumbnails, cloudinarymodels.ImageThumbnail{
            ID:  resource.PublicID, 
            URL: resource.SecureURL,
        })
    }

    return cloudinarymodels.GetThumbnailsResponse{
        Images: thumbnails,
    }, nil
}