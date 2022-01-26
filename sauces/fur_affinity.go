package sauces

import "github.com/asimpleidea/salsa/indexes"

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

// GetSimilarity returns the similarity percentage between the result
// and the provided image.
func (f *FurAffinity) GetSimilarity() float64 {
	return f.Similarity
}

// GetThumbnail returns a URL with the thumbnail of this image.
func (f *FurAffinity) GetThumbnail() string {
	return f.Thumbnail
}

// IsHidden returns true if this result is hidden from the results.
func (f *FurAffinity) IsHidden() bool {
	return f.Hidden
}

// GetIndex returns the index for this.
func (f *FurAffinity) GetIndex() indexes.Index {
	return f.IndexID
}

// GetIndexName returns the index name or page title where this image is
// hosted.
func (f *FurAffinity) GetIndexName() string {
	return f.IndexName
}
