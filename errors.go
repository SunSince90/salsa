package salsa

import "errors"

var (
	// ErrOptions represents an error that occurs when the options provided by
	// the user contains invalid values.
	ErrOptions = errors.New("provided options are not valid")
)

type sauceError struct {
	msg string
	err error
}

func (se *sauceError) Error() string {
	return se.msg
}

func (se *sauceError) Unwrap() error {
	return se.err
}
