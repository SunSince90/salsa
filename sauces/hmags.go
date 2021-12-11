package sauces

import "github.com/SunSince90/salsa/indexes"

// HMags is a source for a picture found from....
// TODO
type HMags struct {
	// SauceHeader is the header for this source.
	*SauceHeader

	// ExternalURLs is a list of urls where the source for this picture is
	// hosted.
	ExternalURLs []string
	// Title for this picture/page.
	Title string
	// Part...
	// TODO
	Part string
	// Date when this was published.
	Date string
}

// GetSimilarity returns the similarity percentage between the result
// and the provided image.
func (h *HMags) GetSimilarity() float64 {
	return h.Similarity
}

// GetThumbnail returns a URL with the thumbnail of this image.
func (h *HMags) GetThumbnail() string {
	return h.Thumbnail
}

// IsHidden returns true if this result is hidden from the results.
func (h *HMags) IsHidden() bool {
	return h.Hidden
}

// GetIndex returns the index for this.
func (h *HMags) GetIndex() indexes.Index {
	return h.IndexID
}

// GetIndexName returns the index name or page title where this image is
// hosted.
func (h *HMags) GetIndexName() string {
	return h.IndexName
}
