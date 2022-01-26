package sauces

import "github.com/asimpleidea/salsa/indexes"

// SeiganIllustration...
// TODO
type SeigaIllustration struct {
	// SauceHeader is the header for this source.
	*SauceHeader

	// ExternalURLs is a list of urls where the source for this picture is
	// hosted.
	ExternalURLs []string
	// ID of the post/page.
	ID int
	// Title of the post/page.
	Title string
	// MemberName is the name of the user that posted this.
	MemberName string
	// MemberID is the ID of the user that posted this.
	MemberID int
}

// GetSimilarity returns the similarity percentage between the result
// and the provided image.
func (s *SeigaIllustration) GetSimilarity() float64 {
	return s.Similarity
}

// GetThumbnail returns a URL with the thumbnail of this image.
func (s *SeigaIllustration) GetThumbnail() string {
	return s.Thumbnail
}

// IsHidden returns true if this result is hidden from the results.
func (s *SeigaIllustration) IsHidden() bool {
	return s.Hidden
}

// GetIndex returns the index for this.
func (s *SeigaIllustration) GetIndex() indexes.Index {
	return s.IndexID
}

// GetIndexName returns the index name or page title where this image is
// hosted.
func (s *SeigaIllustration) GetIndexName() string {
	return s.IndexName
}
