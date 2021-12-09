package sauces

// FurAffinity is a source for a picture found from furaffinity.net.
type FurAffinity struct {
	// SauceHeader is the header for this source.
	*SauceHeader

	// ExternalURLs is a list of urls where the source for this picture is
	// hosted.
	ExternalURLs []string
	// ID of the picture/post.
	ID int
	// Title for this picture.
	Title string
	// AuthorName of the picture.
	AuthorName string
	// AuthorURL is the URL of this author of this picture.
	AuthorURL string
}
