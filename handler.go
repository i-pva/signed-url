package go_url

import "net/http"

// Handler verify the url in the request, and return error as necessary.
func Handler() http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		// verify url
	})
}
