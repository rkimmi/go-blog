package cloudinary

import (
	"log"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/joho/godotenv"
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
	_, err := cloudinary.NewFromURL(cloudinaryURL)
	if err != nil {
		log.Fatalf("Failed to initialize Cloudinary: %v", err)
	}
}

// Returns the Cloudinary instance for use in other parts of the package.
func GetCloudinary() *cloudinary.Cloudinary {
	if cld == nil {
		log.Fatalf("Cloudinary instance is not initialized. Call Init() first.")
	}
	return cld
}
