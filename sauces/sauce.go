package sauces

import "github.com/asimpleidea/salsa/indexes"

// Sauce defines a general source returned by Saucenao.
type Sauce interface {
	// GetSimilarity returns the similarity percentage between the result
	// and the provided image.
	GetSimilarity() float64
	// GetThumbnail returns a URL with the thumbnail of this image.
	GetThumbnail() string
	// IsHidden returns true if this result is hidden from the results.
	IsHidden() bool
	// GetIndex returns the index for this.
	GetIndex() indexes.Index
	// GetIndexName returns the index name or page title where this image is
	// hosted.
	GetIndexName() string
}
