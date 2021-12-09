package sauces

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
