package sauces

// PortalGraphics is a source for a picture found from portalgraphics.net.
type PortalGraphics struct {
	// SauceHeader is the header for this source.
	*SauceHeader

	// ExternalURLs is a list of urls where the source for this picture is
	// hosted.
	ExternalURLs []string
	// ID of the post/page.
	ID int
	// Title of the post/page.
	Title string
	// MemberName is the name of the user that posted this.
	MemberName string
	// MemberID is the ID of the user that posted this.
	MemberID int
}
