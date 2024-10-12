package cloudinary

import (
	"context"
	"log"

	"github.com/cloudinary/cloudinary-go/v2/api"
	"github.com/cloudinary/cloudinary-go/v2/api/admin"
)

// TODO getting getting thumbnail urls rather than full size
func GetAllImagesInFolder(folderName string) ([]api.BriefAssetResult, error) {
	log.Println("Getting images in folder", folderName)

	cld := GetCloudinary()

	// Create a context
	ctx := context.Background()

	// TODO get by folder name!
	resources, err := cld.Admin.Assets(ctx, admin.AssetsParams{})

	if len(resources.Assets) == 0 {
		log.Println("No images found in the folder:", folderName)
	}

	if err != nil {
		log.Fatalf("Failed to list resources: %v", err)
	}

	return resources.Assets, nil
}
