package sauces

import "github.com/asimpleidea/salsa/indexes"

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

// GetSimilarity returns the similarity percentage between the result
// and the provided image.
func (d *DeviantArt) GetSimilarity() float64 {
	return d.Similarity
}

// GetThumbnail returns a URL with the thumbnail of this image.
func (d *DeviantArt) GetThumbnail() string {
	return d.Thumbnail
}

// IsHidden returns true if this result is hidden from the results.
func (d *DeviantArt) IsHidden() bool {
	return d.Hidden
}

// GetIndex returns the index for this.
func (d *DeviantArt) GetIndex() indexes.Index {
	return d.IndexID
}

// GetIndexName returns the index name or page title where this image is
// hosted.
func (d *DeviantArt) GetIndexName() string {
	return d.IndexName
}
