package url

import (
	"errors"
	"net/url"
	"strconv"
	"time"
)

// Create a signed url URL.
func Signed(u *url.URL) (*url.URL, error) {
	return sign(u, 0)
}

// Create a temporary signed url URL.
func TemporarySigned(u *url.URL, expiration time.Duration) (*url.URL, error) {
	return sign(u, expiration)
}

// Sing given url
func sign(u *url.URL, expiration time.Duration) (*url.URL, error) {

	values := u.Query()

	signature := values.Get("signature")
	if len(signature) != 0 {
		return u, errors.New("'signature' is a reserved parameter when generating signed url. Please rename your url parameter")
	}

	if expiration != 0 {
		delay := time.Now().Add(expiration).Unix()
		values.Set("expires", strconv.Itoa(int(delay)))
	}
	values.Set("signature", hash([]byte(u.String())))

	u.RawQuery = values.Encode()
	return u, nil
}
