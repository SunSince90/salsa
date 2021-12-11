package sauces

import "github.com/SunSince90/salsa/indexes"

// Gelbooru is a source for a picture found from gelbooru.com.
type Gelbooru struct {
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
func (g *Gelbooru) GetSimilarity() float64 {
	return g.Similarity
}

// GetThumbnail returns a URL with the thumbnail of this image.
func (g *Gelbooru) GetThumbnail() string {
	return g.Thumbnail
}

// IsHidden returns true if this result is hidden from the results.
func (g *Gelbooru) IsHidden() bool {
	return g.Hidden
}

// GetIndex returns the index for this.
func (g *Gelbooru) GetIndex() indexes.Index {
	return g.IndexID
}

// GetIndexName returns the index name or page title where this image is
// hosted.
func (g *Gelbooru) GetIndexName() string {
	return g.IndexName
}
