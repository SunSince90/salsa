package sauces

// Danbooru is a source for a picture found from danbooru.donmai.us.
type Danbooru struct {
	// SauceHeader is the header for this source.
	*SauceHeader

	// ExternalURLs is a list of urls where the source for this picture is
	// hosted.
	ExternalURLs []string
	// ID of picture/post.
	ID int
	// Creator is a space-separated list of creators for this picture.
	Creator string
	// Material...
	// TODO
	Material string
	// Characters is a space-separated list of supposed characters included in
	// this picture.
	Characters string
	// Source for this picture. Can be blank.
	Source string
}
