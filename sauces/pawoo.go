package sauces

import (
	"time"
)

// Pawoo is a source for a picture found from Pawoo.
type Pawoo struct {
	// SauceHeader is the header for this source.
	*SauceHeader

	// ExternalURLs is a list of urls where the source for this picture is
	// hosted.
	ExternalURLs []string
	// CreatedAt is the time when this was created.
	CreatedAt time.Time
	// ID of the post/page.
	ID int
	// UserAccount of the user that posted this.
	UserAccount string
	// Username of the user that posted this.
	Username string
	// UserDisplayName is the display name of the user that posted this.
	UserDisplayName string
}
