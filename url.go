package url

import (
	"errors"
	"net/url"
	"strconv"
	"time"
)

// Create a signed url URL for a path route.
func Signed(rawurl string) (string, error) {
	return sign(rawurl, 0)
}

// Create a temporary signed url URL for a path route.
func TemporarySigned(rawurl string, expiration time.Duration) (string, error) {
	return sign(rawurl, expiration)
}

func hasValidSignature() bool {
	return false
}

func hasCorrectSignature() bool {
	return false
}

func signatureHasNotExpired() bool {
	return false
}

func sign(rawurl string, expiration time.Duration) (string, error) {
	u, err := url.Parse(rawurl)
	if err != nil {
		return rawurl, err
	}

	values := u.Query()

	signature := values.Get("signature")
	if len(signature) != 0 {
		return rawurl, errors.New("'signature' is a reserved parameter when generating signed url. Please rename your url parameter")
	}

	if expiration != 0 {
		delay := time.Now().Add(expiration).Unix()
		values.Set("expires", strconv.Itoa(int(delay)))
	}
	values.Set("signature", hash([]byte(u.String())))

	u.RawQuery = values.Encode()
	return u.String(), nil
}
