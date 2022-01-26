package sauces

import (
	"time"

	"github.com/asimpleidea/salsa/indexes"
)

// Pawoo is a source for a picture found from Pawoo.
type Pawoo struct {
	// SauceHeader is the header for this source.
	*SauceHeader

	// ExternalURLs is a list of urls where the source for this picture is
	// hosted.
	ExternalURLs []string
	// CreatedAt is the time when this was created.
	CreatedAt time.Time
	// ID of the post/page.
	ID int
	// UserAccount of the user that posted this.
	UserAccount string
	// Username of the user that posted this.
	Username string
	// UserDisplayName is the display name of the user that posted this.
	UserDisplayName string
}

// GetSimilarity returns the similarity percentage between the result
// and the provided image.
func (p *Pawoo) GetSimilarity() float64 {
	return p.Similarity
}

// GetThumbnail returns a URL with the thumbnail of this image.
func (p *Pawoo) GetThumbnail() string {
	return p.Thumbnail
}

// IsHidden returns true if this result is hidden from the results.
func (p *Pawoo) IsHidden() bool {
	return p.Hidden
}

// GetIndex returns the index for this.
func (p *Pawoo) GetIndex() indexes.Index {
	return p.IndexID
}

// GetIndexName returns the index name or page title where this image is
// hosted.
func (p *Pawoo) GetIndexName() string {
	return p.IndexName
}
