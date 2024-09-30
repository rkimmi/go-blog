package cloudinary

import (
    "context"
    "fmt"
    "log"

    "github.com/cloudinary/cloudinary-go/v2/api/admin"
)

func GetAllImagesInFolder(folderName string) {

    cld := GetCloudinary()

    // Create a context
    ctx := context.Background()
    // TODO get by file name!
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
}
