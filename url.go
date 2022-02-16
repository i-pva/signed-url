package url

import (
	"errors"
	"net/url"
	"strconv"
	"time"
)

// ErrSignatureExists represents a failure when signature already exists in url.
var ErrSignatureExists = errors.New("'signature' is a reserved parameter when generating signed url. Please rename your url parameter")

// Sign sings the given url and returns a signed url with error.
func Sign(u *url.URL) (*url.URL, error) {
	return sign(u, 0)
}

// SignTemporary sings the given url with expiration and returns a signed url with error.
func SignTemporary(u *url.URL, expiration time.Duration) (*url.URL, error) {
	return sign(u, expiration)
}

func sign(u *url.URL, expiration time.Duration) (*url.URL, error) {

	values := u.Query()

	signature := values.Get("signature")
	if len(signature) != 0 {
		return u, ErrSignatureExists
	}

	if expiration != 0 {
		delay := time.Now().Add(expiration).Unix()
		values.Set("expires", strconv.Itoa(int(delay)))
		u.RawQuery = values.Encode()
	}
	values.Set("signature", hash([]byte(u.String())))

	u.RawQuery = values.Encode()
	return u, nil
}
