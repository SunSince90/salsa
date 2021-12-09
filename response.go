package salsa

import (
	"github.com/SunSince90/salsa/indexes"
	"github.com/SunSince90/salsa/sauces"
)

// AccountType defines the type of the account of the Saucenao user making
// the request.
type AccountType int

const (
	// Unknown account type.
	Unknown AccountType = iota - 1
	// AnonymousAccount defines that the user is not logged.
	AnonymousAccount
	// NormalAccount specifies that the user is logged.
	NormalAccount
	// PremiumAccount specifies that the user is logged and has a premium
	// subscription.
	PremiumAccount
)

type result struct {
	Header *resultHeader          `json:"header"`
	Data   map[string]interface{} `json:"data"`
}

type sauceResponse struct {
	Header  *responseHeader `json:"header"`
	Results []result        `json:"results"`
}

type resultHeader struct {
	Similarity string        `json:"similarity"`
	Thumbnail  string        `json:"thumbnail"`
	IndexID    indexes.Index `json:"index_id"`
	IndexName  string        `json:"index_name"`
	Dupes      int           `json:"dupes"`
	Hidden     int           `json:"hidden"`
}

func (r *resultHeader) toSauceHeader() *sauces.SauceHeader {
	return &sauces.SauceHeader{
		Similarity: atof64(&r.Similarity),
		Thumbnail:  r.Thumbnail,
		IndexID:    r.IndexID,
		IndexName:  r.IndexName,
		Dupes:      r.Dupes,
		Hidden:     r.Hidden == 1,
	}
}

// Response is the response sent by Saucenao and parsed and cleaned by salsa.
type Response struct {
	// UserID is the ID of the user logged to Saucenao.
	UserID int
	// AccountType defines the type of the account of the logged user.
	AccountType AccountType
	// ShortLimit defines the limit of requests that the user can perform
	// every 30 seconds. (This may change in future: refer to Saucenao account)
	// section.
	ShortLimit int
	// LongLimit defines the limit of requests that the user can perform
	// every 24 hours. (This may change in future: refer to Saucenao account)
	// section.
	LongLimit int
	// ShortRemaining defines how many requests the user can perform before
	// stopping for -- at least -- 30 seconds. Check this before performing any
	// other requests.
	ShortRemaining int
	// LongRemaining defines how many requests the user can perform before
	// stopping for -- at least -- 24 hours.
	LongRemaining int
	// ResultsRequested defines how many results were requested by the user.
	ResultsRequested int
	// ResultsReturned defines the number of results returned by the request.
	ResultsReturned int
	// SearchDepth...
	// TODO
	SearchDepth float64
	// MinimumSimilarity defines the minimum similarity found.
	// TODO: or is it the one requested?
	MinimumSimilarity float64
	// QueryImage is a thumbnail for the requested image.
	QueryImage string
	// QueryImageDisplay...
	// TODO
	QueryImageDisplay string
	// Results is an iterator with all the results returned by Saucenao.
	Results *ResultsIterator

	header *responseHeader
}

func newResponse(r sauceResponse) *Response {
	return &Response{
		header:            r.Header,
		UserID:            atoi(&r.Header.UserID),
		AccountType:       AccountType(atoi(&r.Header.AccountType)),
		ShortLimit:        atoi(&r.Header.ShortLimit),
		LongLimit:         atoi(&r.Header.LongLimit),
		ShortRemaining:    r.Header.ShortRemaining,
		LongRemaining:     r.Header.LongRemaining,
		ResultsRequested:  r.Header.ResultsRequested,
		ResultsReturned:   r.Header.ResultsReturned,
		SearchDepth:       atof64(&r.Header.SearchDepth),
		MinimumSimilarity: r.Header.MinimumSimilarity,
		QueryImage:        r.Header.QueryImage,
		QueryImageDisplay: r.Header.QueryImageDisplay,
		Results:           &ResultsIterator{},
	}
}

// ResultsIterator is an iterator with all the results returned by Saucenao.
type ResultsIterator struct {
	// TODO: implement me
}

// Next returns the next result from the list of results returned by Saucenao.
// You should always check the error before checking the item.
//
// Cast the source to the appropriate structure by checking the second returned
// parameter.
// TODO: make examples.
func (r *ResultsIterator) Next() (interface{}, indexes.Index, error) {
	// TODO: implement me
	return nil, indexes.Unknown, nil
}

type responseHeader struct {
	UserID            string  `json:"user_id"`
	AccountType       string  `json:"account_type"`
	ShortLimit        string  `json:"short_limit"`
	LongLimit         string  `json:"long_limit"`
	LongRemaining     int     `json:"long_remaining"`
	ShortRemaining    int     `json:"short_remaining"`
	Status            int     `json:"status"`
	ResultsRequested  int     `json:"results_requested"`
	ResultsReturned   int     `json:"results_returned"`
	SearchDepth       string  `json:"search_depth"`
	MinimumSimilarity float64 `json:"minimum_similarity"`
	QueryImageDisplay string  `json:"query_image_display"`
	QueryImage        string  `json:"query_image"`
	Message           string  `json:"message"`
}

type responseIndexData struct {
	Status   int           `json:"status"`
	ParentID indexes.Index `json:"parent_id"`
	ID       int           `json:"id"`
	Results  int           `json:"results"`
}
