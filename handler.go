package url

import (
	"errors"
	"net/http"
)

var ErrInvalidSignature = errors.New("invalid signature")

// Handler verify the url from given request, and return error as necessary.
func Handler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if HasValidURL(request) {
			h.ServeHTTP(writer, request)
			return
		}
		http.Error(writer, ErrInvalidSignature.Error(), http.StatusForbidden)
	})
}
