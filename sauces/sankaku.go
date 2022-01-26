package sauces

import "github.com/asimpleidea/salsa/indexes"

// Sankaku is a source for a picture found from chan.sankakucomplex.com.
type Sankaku struct {
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
func (s *Sankaku) GetSimilarity() float64 {
	return s.Similarity
}

// GetThumbnail returns a URL with the thumbnail of this image.
func (s *Sankaku) GetThumbnail() string {
	return s.Thumbnail
}

// IsHidden returns true if this result is hidden from the results.
func (s *Sankaku) IsHidden() bool {
	return s.Hidden
}

// GetIndex returns the index for this.
func (s *Sankaku) GetIndex() indexes.Index {
	return s.IndexID
}

// GetIndexName returns the index name or page title where this image is
// hosted.
func (s *Sankaku) GetIndexName() string {
	return s.IndexName
}
