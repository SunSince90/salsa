package sauces

import "github.com/asimpleidea/salsa/indexes"

// AniDB is a source for a picture found from anidb.net.
type AniDB struct {
	// SauceHeader is the header for this source.
	*SauceHeader

	// ExternalURLs is a list of urls where the source for this picture is
	// hosted.
	ExternalURLs []string
	// Source for this image.
	Source string
	// ID of the post/page where this source was found on.
	ID int
	// Part...
	// TODO
	Part string
	// Year when this anime was running.
	// Can contain multiple years, e.g.: 2012 - 2020.
	Year string
	// EstimatedTime defines the average length of an episode/movie.
	// TODO: inaccurate description?
	EstimatedTime string
}

// GetSimilarity returns the similarity percentage between the result
// and the provided image.
func (a *AniDB) GetSimilarity() float64 {
	return a.Similarity
}

// GetThumbnail returns a URL with the thumbnail of this image.
func (a *AniDB) GetThumbnail() string {
	return a.Thumbnail
}

// IsHidden returns true if this result is hidden from the results.
func (a *AniDB) IsHidden() bool {
	return a.Hidden
}

// GetIndex returns the index for this.
func (a *AniDB) GetIndex() indexes.Index {
	return a.IndexID
}

// GetIndexName returns the index name or page title where this image is
// hosted.
func (a *AniDB) GetIndexName() string {
	return a.IndexName
}
