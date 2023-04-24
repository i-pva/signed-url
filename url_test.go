package url

import (
	"net/http"
	"net/url"
	"testing"
	"time"
)

func init() {
	secretKey = []byte("secret-key")
}

func Test_InValid_HasValidURL(t *testing.T) {
	u := &url.URL{
		Scheme: "http",
		Host:   "test.com",
		Path:   "path",
	}

	v := url.Values{}
	v.Add("expires", "4494900544")
	v.Add("signature", "41d5c3a92c6ef94e80cb70c7dcda0859")

	u.RawQuery = v.Encode()

	if HasValidURL(&http.Request{URL: u}) {
		t.Errorf("HasValidURL() error = signature is valid")
	}
}

func Test_Valid_HasValidURL_With_Sign(t *testing.T) {
	u := &url.URL{
		Scheme: "http",
		Host:   "test.com",
		Path:   "path",
	}

	u, err := Sign(u)
	if err != nil {
		t.Errorf("Signed() error = %v", err)
	}

	if !HasValidURL(&http.Request{URL: u}) {
		t.Errorf("HasValidURL() error = signature is invalid")
	}
}

func Test_Valid_HasValidURL_With_SignTemporary(t *testing.T) {
	u := &url.URL{
		Scheme: "http",
		Host:   "test.com",
		Path:   "path",
	}

	u, err := SignTemporary(u, 1*time.Hour)
	if err != nil {
		t.Errorf("TemporarySigned() error = %v", err)
	}

	if !HasValidURL(&http.Request{URL: u}) {
		t.Errorf("HasValidURL() error = signature is invalid")
	}
}

func Test_InValid_HasValidURL_With_SignTemporary(t *testing.T) {
	u := &url.URL{
		Scheme: "http",
		Host:   "test.com",
		Path:   "path",
	}

	u, err := SignTemporary(u, -1*time.Hour)
	if err != nil {
		t.Errorf("TemporarySigned() error = %v", err)
	}

	if HasValidURL(&http.Request{URL: u}) {
		t.Errorf("HasValidURL() error = signature is invalid")
	}
}
