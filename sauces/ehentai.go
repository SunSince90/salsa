package sauces

import "github.com/SunSince90/salsa/indexes"

// EHentai is a source for a picture found from e-hentai.org.
type EHentai struct {
	// SauceHeader is the header for this source.
	*SauceHeader

	// Source of this picture.
	Source string
	// Creators is a list of creators for this picture.
	Creators []string
	// EnglishName of the gallery.
	EnglishName string
	// JapaneseName of the gallery.
	JapaneseName string
}

// GetSimilarity returns the similarity percentage between the result
// and the provided image.
func (e *EHentai) GetSimilarity() float64 {
	return e.Similarity
}

// GetThumbnail returns a URL with the thumbnail of this image.
func (e *EHentai) GetThumbnail() string {
	return e.Thumbnail
}

// IsHidden returns true if this result is hidden from the results.
func (e *EHentai) IsHidden() bool {
	return e.Hidden
}

// GetIndex returns the index for this.
func (e *EHentai) GetIndex() indexes.Index {
	return e.IndexID
}

// GetIndexName returns the index name or page title where this image is
// hosted.
func (e *EHentai) GetIndexName() string {
	return e.IndexName
}
