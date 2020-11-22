package url

import (
	"net/http"
	"net/http/httptest"
	"net/url"

	"testing"
)

var testUrl = &url.URL{
	Path:     "/test",
	RawQuery: "a=b",
}

func TestHandlerOK(t *testing.T) {
	signed, err := Signed(testUrl)
	if err != nil {
		t.Error(err)
	}

	h := Handler(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
	}))

	server := httptest.NewServer(h)

	req, err := http.NewRequest("GET", server.URL+signed.String(), nil)
	if err != nil {
		t.Error(err)
	}

	res, err := server.Client().Do(req)
	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Error(err)
	}
}
