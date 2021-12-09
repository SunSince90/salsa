package salsa

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
)

func makeRequest(opts *RequestOptions) (sr *sauceResponse, err error) {
	reqURL, _ := url.Parse(defaultBaseURL)
	reqURL.RawQuery = buildQueries(opts).Encode()

	var req *http.Request

	if opts.file != nil {
		req, err = createRequestForFile(reqURL.String(), *opts.file)
	} else {
		req, err = createRequestForURL(reqURL.String(), *opts.url)
	}

	if err != nil {
		return nil, fmt.Errorf("could not prepare request for sauce: %w", err)
	}

	res, err := opts.HttpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error while making request: %w", err)
	}

	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&sr); err != nil {
		return nil, fmt.Errorf("could not parse response from server: %w", err)
	}

	return
}

func createRequestForURL(u, sauceURL string) (req *http.Request, err error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	writer.WriteField("url", sauceURL)

	req, err = http.NewRequest(http.MethodPost, u, body)
	if err != nil {
		return
	}

	req.Header.Add("Content-Type", writer.FormDataContentType())
	return
}

func createRequestForFile(u, imagepath string) (req *http.Request, err error) {
	file, err := os.Open(imagepath)
	if err != nil {
		return
	}
	defer file.Close()

	header, err := createFormFile(file)
	if err != nil {
		return
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	if err = func() (err error) {
		// Wrapping this in a lambda so we can use defer and close the writer
		// on any condition
		defer writer.Close()

		part, err := writer.CreatePart(header)
		if err != nil {
			return
		}

		if _, err = io.Copy(part, file); err != nil {
			return
		}

		return
	}(); err != nil {
		return
	}

	req, err = http.NewRequest(http.MethodPost, u, body)
	if err != nil {
		return
	}

	req.Header.Add("Content-Type", writer.FormDataContentType())
	return
}
