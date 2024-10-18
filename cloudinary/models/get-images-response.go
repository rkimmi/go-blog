package cloudinarymodels

import (
	"github.com/cloudinary/cloudinary-go/v2/api"
)

type ThumbnailsResponse struct {
	Assets     []api.BriefAssetResult
	NextCursor string
}
