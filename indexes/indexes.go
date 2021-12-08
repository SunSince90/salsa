package indexes

import "fmt"

// Index defines an index in the database of websites where to look for source.
type Index int

const (
	Unknown Index = iota - 1
	HMags
	HAnime1
	HCG
	DDBObjects
	DDBSamples
	Pixiv
	PixivHistorical
	Anime1
	SeigaIllustration
	Danbooru
	Drawr
	Nijie
	Yandere
	Animeop
	IMDb
	Shutterstock
	FAKKU
	reserved
	NHentai
	Market2D
	Medibang
	AniDB
	HAnime
	Movies
	Shows
	Gelbooru
	Konachan
	Sankaku
	AnimePictures
	E621
	IdolComplex
	BcyIllust
	BcyCosplay
	PortalGraphics
	DeviantArt
	Pawoo
	Madokami
	Mangadex
	EHentai
	ArtStation
	FurAffinity
	Twitter
	FurryNetwork
	ALL Index = 999
)

var allowedIndexes = map[Index]bool{
	HMags:             true,
	HAnime1:           false,
	HCG:               false,
	DDBObjects:        false,
	DDBSamples:        false,
	Pixiv:             true,
	PixivHistorical:   false,
	Anime1:            false,
	SeigaIllustration: true,
	Danbooru:          true,
	Drawr:             false,
	Nijie:             false,
	Yandere:           false,
	Animeop:           false,
	IMDb:              true,
	Shutterstock:      false,
	FAKKU:             false,
	NHentai:           false,
	Market2D:          false,
	Medibang:          false,
	AniDB:             true,
	Movies:            true,
	Shows:             false,
	Gelbooru:          true,
	Konachan:          false,
	Sankaku:           true,
	AnimePictures:     false,
	E621:              true,
	IdolComplex:       false,
	BcyIllust:         false,
	BcyCosplay:        false,
	PortalGraphics:    true,
	DeviantArt:        true,
	Pawoo:             true,
	Madokami:          false,
	Mangadex:          false,
	EHentai:           true,
	ArtStation:        true,
	FurAffinity:       true,
	Twitter:           false,
	FurryNetwork:      false,
	ALL:               true,
}

// IsIndexValid returns an error if the provided index is not valid or not
// supported. It returns nil otherwise.
func IsIndexValid(index Index) error {
	if val := allowedIndexes[index]; !val {
		return fmt.Errorf("provided index is either invalid or not supported")
	}

	return nil
}
