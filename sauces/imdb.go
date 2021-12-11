package sauces

import "github.com/SunSince90/salsa/indexes"

// IMDb is a source for a picture found from imdb.com.
type IMDb struct {
	// SauceHeader is the header for this source.
	*SauceHeader

	// ExternalURLs is a list of urls where the source for this picture is
	// hosted.
	ExternalURLs []string
	// ID of this movie/show.
	ID string
	// Source for the picture.
	Source string
	// Part...
	// TODO
	Part string
	// Year (or years range) when this was aired.
	Year string
	// EstimatedTime is the estimated duration of this movie or episode.
	EstimatedTime string
}

// GetSimilarity returns the similarity percentage between the result
// and the provided image.
func (i *IMDb) GetSimilarity() float64 {
	return i.Similarity
}

// GetThumbnail returns a URL with the thumbnail of this image.
func (i *IMDb) GetThumbnail() string {
	return i.Thumbnail
}

// IsHidden returns true if this result is hidden from the results.
func (i *IMDb) IsHidden() bool {
	return i.Hidden
}

// GetIndex returns the index for this.
func (i *IMDb) GetIndex() indexes.Index {
	return i.IndexID
}

// GetIndexName returns the index name or page title where this image is
// hosted.
func (i *IMDb) GetIndexName() string {
	return i.IndexName
}
