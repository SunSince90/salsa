package salsa

import (
	"context"
	"fmt"
	"net/http"
	"path/filepath"
	"time"

	strip "github.com/grokify/html-strip-tags-go"
)

const (
	defaultBaseURL string = "https://saucenao.com/search.php"
)

// Client is an interface for communicating with Saucenao and querying for
// source for a local or remote image .
type Client interface {
	// GetSauceForFile tries to get source for a local file.
	//
	// The response from Saucenao is returned *only* if the provided request
	// options are valid. Otherwise the response is always present.
	// Users should *always* check for errors before accessing the response.
	// See examples for this.
	GetSauceForFile(context.Context, string, ...RequestOption) (*Response, error)
	// GetSauceForURL tries to get source for a remote picture.
	//
	// The response from Saucenao is returned *only* if the provided request
	// options are valid. Otherwise the response is always present.
	// Users should *always* check for errors before accessing the response.
	// See examples for this.
	GetSauceForURL(context.Context, string, ...RequestOption) (*Response, error)
}

// SauceClient is an implementation of the Client interface for connecting to
// Saucenao and query for source to a picture.
type SauceClient struct {
	apiKey   string
	defaults *RequestOptions
}

// NewSauceClient returns an instance for the Client.
//
// An API key *must* be provided in order to have a valid client, otherwise an
// error will be returned and the client returned will be nil.
// Optionally a list of options can be provided which will be included with
// each call. An error will be returned if provided options are not valid and
// a nil client will be returned.
func NewSauceClient(apiKey string, defaultOptions ...RequestOption) (*SauceClient, error) {
	if apiKey == "" {
		return nil, fmt.Errorf("missing api key")
	}

	def := &RequestOptions{
		apiKey:        apiKey,
		ResultsNumber: defaultResultsNumber,
		TestMode:      defaultTestMode,
		DB:            defaultIndex,
		HttpClient: &http.Client{
			Timeout: time.Minute,
		},
		MinimumSimilarity: defaultMinSimilarity,
		SafetyLevel:       ShowAll,
	}

	for _, opt := range defaultOptions {
		opt(def)
	}
	if err := validateRequestOptions(def); err != nil {
		return nil, fmt.Errorf("could not create valid client: %w", err)
	}

	return &SauceClient{
		apiKey:   apiKey,
		defaults: def,
	}, nil
}

// GetSauceForFile tries to get source for a local file.
//
// The response from Saucenao is returned *only* if the provided request
// options are valid. Otherwise the response is always present.
// Users should *always* check for errors before accessing the response.
// See examples for this.
func (c *SauceClient) GetSauceForFile(ctx context.Context, imagepath string, opts ...RequestOption) (*Response, error) {
	allowedExtensions := map[string]bool{
		".jpg":  true,
		".png":  true,
		".jpeg": true,
		".webp": true,
		".bmp":  true,
		".gif":  true,
	}

	ext := filepath.Ext(imagepath)
	if _, exists := allowedExtensions[ext]; !exists {
		return nil, &sauceError{msg: "file extension not allowed", err: ErrOptions}
	}

	options := c.defaults.clone()

	for _, opt := range opts {
		opt(options)
	}

	if err := validateRequestOptions(options); err != nil {
		return nil, fmt.Errorf("could not request for sauce: %w", err)
	}

	options.file = &imagepath

	sauceResp, err := makeRequest(options)
	if err != nil {
		return nil, fmt.Errorf("request for sauce was unsucessful: %w", err)
	}

	if sauceResp.Header.Status != 0 {
		if sauceResp.Header.Status > 0 {
			err = &sauceError{strip.StripTags(sauceResp.Header.Message), ErrServerSide}
		} else {
			err = &sauceError{strip.StripTags(sauceResp.Header.Message), ErrClientSide}
		}
	} else {
		err = nil
	}

	return newResponse(*sauceResp), err
}

// NewSauceClient returns an instance for the Client.
//
// An API key *must* be provided in order to have a valid client, otherwise an
// error will be returned and the client returned will be nil.
// Optionally a list of options can be provided which will be included with
// each call. An error will be returned if provided options are not valid and
// a nil client will be returned.
func (c *SauceClient) GetSauceForURL(ctx context.Context, url string, opts ...RequestOption) (*Response, error) {
	options := c.defaults.clone()

	for _, opt := range opts {
		opt(options)
	}

	if err := validateRequestOptions(options); err != nil {
		return nil, fmt.Errorf("could not request for sauce: %w", err)
	}

	options.url = &url

	sauceResp, err := makeRequest(options)
	if err != nil {
		return nil, fmt.Errorf("request for sauce was unsucessful: %w", err)
	}

	return newResponse(*sauceResp), nil
}
