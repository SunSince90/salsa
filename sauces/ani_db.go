package sauces

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
