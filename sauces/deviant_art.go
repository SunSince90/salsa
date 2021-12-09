package sauces

// DeviantArt is a source for a picture found from deviantart.com.
type DeviantArt struct {
	// SauceHeader is the header for this source.
	*SauceHeader

	// ExternalURLs is a list of urls where the source for this picture is
	// hosted.
	ExternalURLs []string
	// ID of the post.
	ID int
	// Title of the post.
	Title string
	// AuthorName is the name of the author of this picture.
	AuthorName string
	// AuthorURL is the URL of the author's profile.
	AuthorURL string
}
