package salsa

import (
	"bytes"
	"mime/multipart"
	"net/http"
)

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
