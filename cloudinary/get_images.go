package cloudinary

import (
	"context"
	"log"

	"github.com/cloudinary/cloudinary-go/v2/api"
	"github.com/cloudinary/cloudinary-go/v2/api/admin"
)

func GetAllImagesInFolder(folderName string) ([]api.BriefAssetResult, error) {
	log.Println("Getting images in folder", folderName)

	cld := GetCloudinary()

	// Create a context
	ctx := context.Background()
	// TODO get by folder name!
	resources, err := cld.Admin.Assets(ctx, admin.AssetsParams{})

	log.Printf("%+v\n", resources)

	if len(resources.Assets) == 0 {
		log.Println("No images found in the folder:", folderName)
	}

	log.Printf("Total assets found: %d\n", len(resources.Assets))

	if err != nil {
		log.Fatalf("Failed to list resources: %v", err)
	}

	for _, asset := range resources.Assets {
		log.Println("Image URL:", asset.SecureURL)
	}

	return resources.Assets, nil
}
