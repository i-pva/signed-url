package url

import (
	"net/url"
	"testing"
	"time"
)

func TestSign(t *testing.T) {
	u, err := url.Parse("https://example.com/path?key=value")
	if err != nil {
		t.Error(err)
	}

	signedURL, err := sign(u, 0)
	if err != nil {
		t.Error(err)
	}

	// Verify that the "signature" parameter was added to the query string
	if signedURL.Query().Get("signature") == "" {
		t.Error("signature parameter not added to query string")
	}

	// Verify that the "key" parameter is in s query string
	if signedURL.Query().Get("key") == "" {
		t.Error("key parameter was removed from query string")
	}
}

func TestSignTemporary(t *testing.T) {
	u, err := url.Parse("https://example.com/path?key=value")
	if err != nil {
		t.Error(err)
	}

	expiration := time.Hour
	signedURL, err := sign(u, expiration)
	if err != nil {
		t.Error(err)
	}

	// Verify that the "expires" parameter was added to the query string
	if signedURL.Query().Get("expires") == "" {
		t.Error("expires parameter not added to query string")
	}

	// Verify that the "signature" parameter was added to the query string
	if signedURL.Query().Get("signature") == "" {
		t.Error("signature parameter not added to query string")
	}
}

func TestSign_WithExistingSignatureParameter(t *testing.T) {
	u, err := url.Parse("https://example.com/path?key=value&signature=abc")
	if err != nil {
		t.Error(err)
	}

	_, err = sign(u, 0)

	// Verify that the function returns an error when the "signature" parameter already exists in the query string
	if err != ErrSignatureExists {
		t.Error("expected ErrSignatureExists, got", err)
	}
}
