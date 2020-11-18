package url

import (
	"errors"
	"net/http"
)

var ErrInvalidSignature = errors.New("invalid signature")

// Handler verify the url in the request, and return error as necessary.
func Handler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if hasValidSignature() {
			h.ServeHTTP(writer, request)
			return
		}

		http.Error(writer, ErrInvalidSignature.Error(), http.StatusForbidden)
	})
}
