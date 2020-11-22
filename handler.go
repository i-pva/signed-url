package url

import (
	"errors"
	"net/http"
	"strconv"
	"time"
)

var ErrInvalidSignature = errors.New("invalid signature")

// Handler verify the url from given request, and return error as necessary.
func Handler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if hasValidSignature(request) {
			h.ServeHTTP(writer, request)
			return
		}
		http.Error(writer, ErrInvalidSignature.Error(), http.StatusForbidden)
	})
}

func hasValidSignature(request *http.Request) bool {
	return hasCorrectSignature(request) && signatureHasNotExpired(request)
}

// Determine if the signature from the given request matches the URL.
func hasCorrectSignature(request *http.Request) bool {
	url := request.URL

	values := url.Query()
	signature := values.Get("signature")
	values.Del("signature")
	url.RawQuery = values.Encode()

	return signature == hash([]byte(url.String()))
}

// Determine if the expires timestamp from the given request is not from the past.
func signatureHasNotExpired(request *http.Request) bool {
	url := request.URL
	expires := url.Query().Get("expires")
	unix, _ := strconv.Atoi(expires)
	if unix == 0 {
		return true
	}

	return int(time.Now().Unix()) < unix
}
