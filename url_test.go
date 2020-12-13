package url

import (
	"net/http"
	"net/url"
	"testing"
	"time"
)

func Test_InValid_HasValidURL(t *testing.T) {
	u := &url.URL{
		Scheme: "http",
		Host:   "test.com",
		Path:   "path",
	}

	SecretKey = []byte("secret-key")

	v := url.Values{}
	v.Add("expires", "4494900544")
	v.Add("signature", "41d5c3a92c6ef94e80cb70c7dcda0859")

	u.RawQuery = v.Encode()

	if HasValidURL(&http.Request{URL: u}) {
		t.Errorf("HasValidURL() error = signature is valid")
	}
}

func Test_Valid_HasValidURL_With_Signed(t *testing.T) {
	u := &url.URL{
		Scheme: "http",
		Host:   "test.com",
		Path:   "path",
	}

	SecretKey = []byte("secret-key")

	u, err := Signed(u)
	if err != nil {
		t.Errorf("Signed() error = %v", err)
	}

	if !HasValidURL(&http.Request{URL: u}) {
		t.Errorf("HasValidURL() error = signature is invalid")
	}
}

func Test_Valid_HasValidURL_With_TemporarySigned(t *testing.T) {
	u := &url.URL{
		Scheme: "http",
		Host:   "test.com",
		Path:   "path",
	}

	SecretKey = []byte("secret-key")

	u, err := TemporarySigned(u, 1*time.Hour)
	if err != nil {
		t.Errorf("TemporarySigned() error = %v", err)
	}

	if !HasValidURL(&http.Request{URL: u}) {
		t.Errorf("HasValidURL() error = signature is invalid")
	}
}

func Test_InValid_HasValidURL_With_TemporarySigned(t *testing.T) {
	u := &url.URL{
		Scheme: "http",
		Host:   "test.com",
		Path:   "path",
	}

	SecretKey = []byte("secret-key")

	u, err := TemporarySigned(u, -1*time.Hour)
	if err != nil {
		t.Errorf("TemporarySigned() error = %v", err)
	}

	if HasValidURL(&http.Request{URL: u}) {
		t.Errorf("HasValidURL() error = signature is invalid")
	}
}
