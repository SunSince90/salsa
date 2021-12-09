package salsa

import (
	"time"

	"github.com/SunSince90/salsa/sauces"
)

type result struct {
	Header *resultHeader `json:"header"`
	Data   resultData    `json:"data"`
}

func (r *result) deviantArt() *sauces.DeviantArt {
	return &sauces.DeviantArt{
		SauceHeader:  r.Header.toSauceHeader(),
		ExternalURLs: r.Data.toStringSlice("ext_urls"),
		Title:        r.Data.toString("title"),
		ID:           r.Data.toIntegerID("da_id"),
		AuthorName:   r.Data.toString("author_name"),
		AuthorURL:    r.Data.toString("author_url"),
	}
}

func (r *result) eHentai() *sauces.EHentai {
	return &sauces.EHentai{
		SauceHeader:  r.Header.toSauceHeader(),
		Source:       r.Data.toString("source"),
		Creators:     r.Data.toStringSlice("creator"),
		EnglishName:  r.Data.toString("eng_name"),
		JapaneseName: r.Data.toString("jp_name"),
	}
}

func (r *result) artStation() *sauces.ArtStation {
	return &sauces.ArtStation{
		SauceHeader:  r.Header.toSauceHeader(),
		ExternalURLs: r.Data.toStringSlice("ext_urls"),
		Title:        r.Data.toString("title"),
		Project:      r.Data.toString("as_project"),
		AuthorName:   r.Data.toString("author_name"),
		AuthorURL:    r.Data.toString("author_url"),
	}
}

func (r *result) pixiv() *sauces.Pixiv {
	return &sauces.Pixiv{
		SauceHeader:  r.Header.toSauceHeader(),
		ExternalURLs: r.Data.toStringSlice("ext_urls"),
		Title:        r.Data.toString("title"),
		ID:           r.Data.toIntegerID("pixiv_id"),
		MemberName:   r.Data.toString("member_name"),
		MemberID:     r.Data.toIntegerID("member_id"),
	}
}

func (r *result) aniDB() *sauces.AniDB {
	return &sauces.AniDB{
		SauceHeader:   r.Header.toSauceHeader(),
		ExternalURLs:  r.Data.toStringSlice("ext_urls"),
		Source:        r.Data.toString("source"),
		ID:            r.Data.toIntegerID("anidb_aid"),
		Part:          r.Data.toString("part"),
		Year:          r.Data.toString("year"),
		EstimatedTime: r.Data.toString("est_time"),
	}
}

func (r *result) pawoo() *sauces.Pawoo {
	return &sauces.Pawoo{
		SauceHeader:     r.Header.toSauceHeader(),
		ExternalURLs:    r.Data.toStringSlice("ext_urls"),
		CreatedAt:       r.Data.toTime("created_at", time.RFC3339),
		ID:              r.Data.toIntegerID("pawoo_id"),
		UserAccount:     r.Data.toString("pawoo_user_acct"),
		Username:        r.Data.toString("pawoo_user_username"),
		UserDisplayName: r.Data.toString("pawoo_user_display_name"),
	}
}

func (r *result) gelbooru() *sauces.Gelbooru {
	return &sauces.Gelbooru{
		SauceHeader:  r.Header.toSauceHeader(),
		ExternalURLs: r.Data.toStringSlice("ext_urls"),
		ID:           r.Data.toIntegerID("gelbooru_id"),
		Creator:      r.Data.toString("creator"),
		Material:     r.Data.toString("material"),
		Characters:   r.Data.toString("characters"),
		Source:       r.Data.toString("source"),
	}
}

func (r *result) danbooru() *sauces.Danbooru {
	return &sauces.Danbooru{
		SauceHeader:  r.Header.toSauceHeader(),
		ExternalURLs: r.Data.toStringSlice("ext_urls"),
		ID:           r.Data.toIntegerID("danbooru_id"),
		Creator:      r.Data.toString("creator"),
		Material:     r.Data.toString("material"),
		Characters:   r.Data.toString("characters"),
		Source:       r.Data.toString("source"),
	}
}

func (r *result) e621() *sauces.E621 {
	return &sauces.E621{
		SauceHeader:  r.Header.toSauceHeader(),
		ExternalURLs: r.Data.toStringSlice("ext_urls"),
		ID:           r.Data.toIntegerID("e621_id"),
		Creator:      r.Data.toString("creator"),
		Material:     r.Data.toString("material"),
		Characters:   r.Data.toString("characters"),
		Source:       r.Data.toString("source"),
	}
}

func (r *result) portalGraphics() *sauces.PortalGraphics {
	return &sauces.PortalGraphics{
		SauceHeader:  r.Header.toSauceHeader(),
		ExternalURLs: r.Data.toStringSlice("ext_urls"),
		ID:           r.Data.toIntegerID("pg_id"),
		Title:        r.Data.toString("title"),
		MemberName:   r.Data.toString("member_name"),
		MemberID:     r.Data.toIntegerID("member_id"),
	}
}

func (r *result) sankaku() *sauces.Sankaku {
	return &sauces.Sankaku{
		SauceHeader:  r.Header.toSauceHeader(),
		ExternalURLs: r.Data.toStringSlice("ext_urls"),
		ID:           r.Data.toIntegerID("sankaku_id"),
		Creator:      r.Data.toString("creator"),
		Material:     r.Data.toString("material"),
		Characters:   r.Data.toString("characters"),
		Source:       r.Data.toString("source"),
	}
}

func (r *result) furAffinity() *sauces.FurAffinity {
	return &sauces.FurAffinity{
		SauceHeader:  r.Header.toSauceHeader(),
		ExternalURLs: r.Data.toStringSlice("ext_urls"),
		ID:           r.Data.toIntegerID("fa_id"),
		AuthorName:   r.Data.toString("author_name"),
		AuthorURL:    r.Data.toString("author_url"),
	}
}

func (r *result) seigaIllustration() *sauces.SeigaIllustration {
	return &sauces.SeigaIllustration{
		SauceHeader:  r.Header.toSauceHeader(),
		ExternalURLs: r.Data.toStringSlice("ext_urls"),
		ID:           r.Data.toIntegerID("seiga_id"),
		MemberName:   r.Data.toString("member_name"),
		MemberID:     r.Data.toIntegerID("member_id"),
	}
}

func (r *result) hMags() *sauces.HMags {
	return &sauces.HMags{
		SauceHeader:  r.Header.toSauceHeader(),
		ExternalURLs: r.Data.toStringSlice("ext_urls"),
		Title:        r.Data.toString("title"),
		Part:         r.Data.toString("part"),
		Date:         r.Data.toString("date"),
	}
}

func (r *result) iMDb() *sauces.IMDb {
	return &sauces.IMDb{
		SauceHeader:   r.Header.toSauceHeader(),
		ExternalURLs:  r.Data.toStringSlice("ext_urls"),
		ID:            r.Data.toString("imdb_id"),
		Part:          r.Data.toString("creator"),
		Year:          r.Data.toString("year"),
		EstimatedTime: r.Data.toString("est_time"),
	}
}

type resultData map[string]interface{}

func (r resultData) toString(key string) string {
	val, exists := r[key]
	if !exists {
		return ""
	}

	switch val := val.(type) {
	case string:
		return val
	default:
		return ""
	}
}

func (r resultData) toStringSlice(key string) (values []string) {
	val, exists := r[key]
	if !exists {
		return []string{}
	}

	switch val := val.(type) {
	case []interface{}:
		for i := 0; i < len(val); i++ {
			values = append(values, val[i].(string))
		}
		return
	case interface{}:
		values = append(values, val.(string))
	default:
		values = []string{}
	}

	return
}

func (r resultData) toIntegerID(key string) int {
	val, exists := r[key]
	if !exists {
		return -1
	}

	switch val := val.(type) {
	case string:
		return atoi(&val)
	case int:
		return val
	case float64:
		// We're talking IDs here: pretty sure the compiler set them
		// as float64 to be as sure as possible, but we know better in
		// this case, and that is actually an int.
		return int(val)
	default:
		return -1
	}
}

func (r resultData) toTime(key, format string) time.Time {
	defaultTime := time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC)
	val, exists := r[key]
	if !exists {
		return defaultTime
	}

	switch val := val.(type) {
	case string:
		t, err := time.Parse(format, val)
		if err != nil {
			return defaultTime
		}
		return t
		// TODO: format timestamp
	default:
		return defaultTime
	}
}
