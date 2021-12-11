package sauces

import "github.com/SunSince90/salsa/indexes"

// Danbooru is a source for a picture found from danbooru.donmai.us.
type Danbooru struct {
	// SauceHeader is the header for this source.
	*SauceHeader

	// ExternalURLs is a list of urls where the source for this picture is
	// hosted.
	ExternalURLs []string
	// ID of picture/post.
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
func (d *Danbooru) GetSimilarity() float64 {
	return d.Similarity
}

// GetThumbnail returns a URL with the thumbnail of this image.
func (d *Danbooru) GetThumbnail() string {
	return d.Thumbnail
}

// IsHidden returns true if this result is hidden from the results.
func (d *Danbooru) IsHidden() bool {
	return d.Hidden
}

// GetIndex returns the index for this.
func (d *Danbooru) GetIndex() indexes.Index {
	return d.IndexID
}

// GetIndexName returns the index name or page title where this image is
// hosted.
func (d *Danbooru) GetIndexName() string {
	return d.IndexName
}
