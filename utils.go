package salsa

import (
	"fmt"
	"io"
	"net/textproto"
	"os"
	"path/filepath"

	"github.com/gabriel-vasile/mimetype"
)

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
