package sauces

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
