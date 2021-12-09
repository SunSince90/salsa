package salsa

import (
	"fmt"
	"io"
	"net/textproto"
	"net/url"
	"os"
	"path/filepath"
	"strconv"

	"github.com/SunSince90/salsa/indexes"
	"github.com/gabriel-vasile/mimetype"
)

const (
	defaultTestMode      bool          = false
	defaultIndex         indexes.Index = indexes.ALL
	defaultResultsNumber int           = 6
	defaultMinSimilarity int           = 55
	defaultSafetyLevel   SafetyLevel   = ShowAll
)

const (
	queryApiKey     string = "api_key"
	queryOutputType string = "output_type"
	queryOutputJSON string = "2"
	queryNumRes     string = "numres"
	queryDb         string = "db"
	queryMinSim     string = "minsim"
	queryTestMode   string = "testmode"
	queryHide       string = "hide"
)

func validateRequestOptions(opts *RequestOptions) error {
	if opts.ResultsNumber == 0 {
		return &sauceError{msg: "must request at least one result", err: ErrOptions}
	}

	if err := indexes.IsIndexValid(opts.DB); err != nil {
		return err
	}

	if opts.MinimumSimilarity <= 0 || opts.MinimumSimilarity > 99 {
		return &sauceError{msg: "minimum similarity value is not valid", err: ErrOptions}
	}

	if opts.SafetyLevel > Safest {
		return &sauceError{msg: "invalid safety level provided", err: ErrOptions}
	}

	return nil
}

func buildQueries(opts *RequestOptions) url.Values {
	queries := (&url.URL{}).Query()

	queries.Add(queryApiKey, opts.apiKey)
	queries.Add(queryOutputType, queryOutputJSON)
	queries.Add(queryNumRes, strconv.Itoa(opts.ResultsNumber))
	queries.Add(queryDb, strconv.Itoa(int(opts.DB)))
	queries.Add(queryMinSim, strconv.Itoa(opts.MinimumSimilarity))
	if opts.TestMode {
		queries.Add(queryTestMode, "1")
	}
	if opts.SafetyLevel != ShowAll {
		queries.Add(queryHide, strconv.Itoa(int(opts.SafetyLevel)))
	}

	return queries
}

func createFormFile(file *os.File) (textproto.MIMEHeader, error) {
	defer file.Seek(0, io.SeekStart)
	h := textproto.MIMEHeader{}

	contentType, err := mimetype.DetectReader(file)
	if err != nil {
		return h, err
	}

	fname := filepath.Base(file.Name())
	formData := fmt.Sprintf(`form-data; name="file"; filename="%s"`, fname)
	h.Set("Content-Disposition", formData)
	h.Set("Content-Type", contentType.String())

	return h, nil
}
