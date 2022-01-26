package salsa

import (
	"net/http"

	"github.com/asimpleidea/salsa/indexes"
)

// RequestOptions contains options to include with the request.
type RequestOptions struct {
	apiKey string
	file   *string
	url    *string

	// ResultsNumber specifies the maximum number of results to return.
	// Default is 6.
	ResultsNumber int
	// Test mode specifies whether to enable test mode. Read Saucenao
	// documentation to learn more about this. Default is false.
	TestMode bool
	// DB specifies the index where to search for. Default is ALL.
	DB indexes.Index
	// HttpClient to use. Use this in case you want to have custom http
	// options.
	HttpClient *http.Client
	// MinimumSimilarity specifies the minimum similarity for the results.
	// Default is 55.
	MinimumSimilarity int
	// SafetyLevel specifies whether to filter explicit or suspected explicit
	// results. Default level is ShowAll.
	SafetyLevel SafetyLevel
}

func (r *RequestOptions) clone() *RequestOptions {
	return &RequestOptions{
		apiKey: r.apiKey,
		file:   r.file,
		url:    r.url,

		ResultsNumber:     r.ResultsNumber,
		TestMode:          r.TestMode,
		DB:                r.DB,
		HttpClient:        r.HttpClient,
		MinimumSimilarity: r.MinimumSimilarity,
		SafetyLevel:       r.SafetyLevel,
	}
}

// SafetyLevel is the level of the
type SafetyLevel int

const (
	// ShowAll specifies the all results should be returned.
	ShowAll SafetyLevel = iota
	// HideExpectedExplicit tells Saucenao to hide results that are expected
	// to be explicit.
	HideExpectedExplicit
	// HideExpectedAndSuspectedExplicit is exactly like HideExpectedExplicit
	// but it also hides results that are suspected to be explicit.
	HideExpectedAndSuspectedExplicit
	// Safest only shows safe results.
	Safest
)

// RequestOption adds an option to the request.
type RequestOption func(*RequestOptions)

// WithResultsNumber sets the maximum number of the results.
func WithResultsNumber(num int) RequestOption {
	return func(ro *RequestOptions) {
		ro.ResultsNumber = num
	}
}

// WithTestMode sets test mode to true. Use this only when testing.
func WithTestMode() RequestOption {
	return func(ro *RequestOptions) {
		ro.TestMode = true
	}
}

// WithIndex modifies the index to look for.
func WithIndex(index indexes.Index) RequestOption {
	return func(ro *RequestOptions) {
		ro.DB = index
	}
}

// WithHTTPClient sets a custom HTTP Client instead of the new one.
func WithHTTPClient(client *http.Client) RequestOption {
	return func(ro *RequestOptions) {
		ro.HttpClient = client
	}
}

// WithMinimumSimilarity tells Saucenao to filter results that have a
// similarity lower than the one provided.
func WithMinimumSimilarity(minimum int) RequestOption {
	return func(ro *RequestOptions) {
		ro.MinimumSimilarity = minimum
	}
}

// WithSafetyLevel sets the safety level for this request.
func WithSafetyLevel(mode SafetyLevel) RequestOption {
	return func(ro *RequestOptions) {
		ro.SafetyLevel = mode
	}
}
