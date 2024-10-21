package cloudinary

import (
	"log"
	"os"

	"context"

	"github.com/joho/godotenv"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/admin"

	cloudinarymodels "blog/cloudinary/models"
)

// Cloudinary instance to be used across the package
var cld *cloudinary.Cloudinary

func Init() {
	log.Println("Starting Cloudinary initialization...")

	cloudinaryURL := os.Getenv("CLOUDINARY_URL")
	if cloudinaryURL == "" {
		log.Fatal("CLOUDINARY_URL environment variable is missing")
	}

	if os.Getenv("FLY_APP_NAME") == "" {
		log.Println("Running locally, loading .env file...")
		envErr := godotenv.Load()
		if envErr != nil {
			log.Fatal("Error loading .env file")
		}
	} else {
		log.Println("Running on Fly.io, using environment variables...")
	}

	// Initialize Cloudinary
	var err error
	cld, err = cloudinary.NewFromURL(cloudinaryURL)
	if err != nil {
		log.Fatalf("Failed to initialize Cloudinary: %v", err)
	}

	log.Println("Cloudinary initialized")
}

func GetAllImagesInFolder(folderName string) (cloudinarymodels.ThumbnailsResponse, error) {
	log.Println("Getting images in foldfdfder", folderName)

	cld := GetCloudinary()

	// Create a context
	ctx := context.Background()

	// TODO get by folder name!
	// TODO pagination with nextcursor
	resources, err := cld.Admin.Assets(ctx, admin.AssetsParams{
		MaxResults: 20,
	})

	if len(resources.Assets) == 0 {
		log.Println("No images found in the folder:", folderName)
	}

	if err != nil {
		log.Fatalf("Failed to list resources: %v", err)
	}

	return cloudinarymodels.ThumbnailsResponse{
		Assets:     resources.Assets,
		NextCursor: resources.NextCursor,
	}, nil
}

// Returns the Cloudinary instance for use in other parts of the package.
func GetCloudinary() *cloudinary.Cloudinary {
	if cld == nil {
		log.Fatalf("Cloudinary instance is not initialized. Call Init() first.")
	}
	return cld
}
