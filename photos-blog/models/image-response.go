package imagemodels

type ImageThumbnail struct {
	ID  string `json:"id"`
	URL string `json:"url"`
}

type GetThumbnailsResponse struct {
	Images         []ImageThumbnail `json:"images"`
	NextImageStart string           `json:"nextPageStart"`
}
