package sauces

import "github.com/SunSince90/salsa/indexes"

// SauceHeader contains the header for the source, with additional information
// about the source, such as the similarity, thumbnail etc.
type SauceHeader struct {
	// Similarity is the percentage of similarity between the provided image
	// an the found source.
	Similarity float64
	// Thumbnail for the source.
	Thumbnail string
	// IndexID is the ID of the index for this source.
	// TODO: delete this on future versions?
	IndexID indexes.Index
	// IndexName is the name of the website's page where the source was found.
	// TODO: inaccurate description?
	IndexName string
	// Dupes is the number of duplicates found for this index.
	// TODO: inaccurate description?
	Dupes int
	// Hidden specifies whether this source should be hidden because of safety
	// rules.
	// TODO: inaccurate description?
	Hidden bool
}
