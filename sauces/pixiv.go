package sauces

import "github.com/asimpleidea/salsa/indexes"

// Pixiv is a source for a picture found from pixiv.net
type Pixiv struct {
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
func (p *Pixiv) GetSimilarity() float64 {
	return p.Similarity
}

// GetThumbnail returns a URL with the thumbnail of this image.
func (p *Pixiv) GetThumbnail() string {
	return p.Thumbnail
}

// IsHidden returns true if this result is hidden from the results.
func (p *Pixiv) IsHidden() bool {
	return p.Hidden
}

// GetIndex returns the index for this.
func (p *Pixiv) GetIndex() indexes.Index {
	return p.IndexID
}

// GetIndexName returns the index name or page title where this image is
// hosted.
func (p *Pixiv) GetIndexName() string {
	return p.IndexName
}
