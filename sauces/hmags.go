package sauces

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
