package sauces

import "github.com/SunSince90/salsa/indexes"

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

// GetSimilarity returns the similarity percentage between the result
// and the provided image.
func (a *ArtStation) GetSimilarity() float64 {
	return a.Similarity
}

// GetThumbnail returns a URL with the thumbnail of this image.
func (a *ArtStation) GetThumbnail() string {
	return a.Thumbnail
}

// IsHidden returns true if this result is hidden from the results.
func (a *ArtStation) IsHidden() bool {
	return a.Hidden
}

// GetIndex returns the index for this.
func (a *ArtStation) GetIndex() indexes.Index {
	return a.IndexID
}

// GetIndexName returns the index name or page title where this image is
// hosted.
func (a *ArtStation) GetIndexName() string {
	return a.IndexName
}
