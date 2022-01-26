package sauces

import "github.com/asimpleidea/salsa/indexes"

// E621 is a source for a picture found from e621.net.
type E621 struct {
	// SauceHeader is the header for this source.
	*SauceHeader

	// ExternalURLs is a list of urls where the source for this picture is
	// hosted.
	ExternalURLs []string
	// ID of the picture/post.
	ID int
	// Creator is a space-separated list of creators for this picture.
	Creator string
	// Material...
	// TODO
	Material string
	// Characters is a space-separated list of supposed characters included in
	// this picture.
	Characters string
	// Source for this picture. Can be blank.
	Source string
}

// GetSimilarity returns the similarity percentage between the result
// and the provided image.
func (e *E621) GetSimilarity() float64 {
	return e.Similarity
}

// GetThumbnail returns a URL with the thumbnail of this image.
func (e *E621) GetThumbnail() string {
	return e.Thumbnail
}

// IsHidden returns true if this result is hidden from the results.
func (e *E621) IsHidden() bool {
	return e.Hidden
}

// GetIndex returns the index for this.
func (e *E621) GetIndex() indexes.Index {
	return e.IndexID
}

// GetIndexName returns the index name or page title where this image is
// hosted.
func (e *E621) GetIndexName() string {
	return e.IndexName
}
