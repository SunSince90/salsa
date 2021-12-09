package sauces

// ArtStation is a source for a picture found from artstation.com.
type ArtStation struct {
	// SauceHeader is the header for this source.
	*SauceHeader

	// ExternalURLs is a list of urls where the source for this picture is
	// hosted.
	ExternalURLs []string
	// Title of the picture/page.
	Title string
	// Project...
	// TODO
	Project string
	// AuthorName is the name of the author of this picture.
	AuthorName string
	// AuthorURL is the URL where to find the author of this picture.
	AuthorURL string
}
