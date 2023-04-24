package url

import (
	"net/http"
	"net/url"
	"testing"
)

func TestHasValidURL_WithExpiredSignature(t *testing.T) {
	SetSecretKey([]byte("mySecretKey"))

	// Create a URL with an expired signature
	u, _ := url.Parse("https://example.com/path?key=value&signature=23a2bc5a5dc71f0f91061d55ebe666a39e87a29f9812b3053a3d07e99f5b5a8a&expires=1619664800")
	req := &http.Request{URL: u}

	// Verify that the function returns false
	if HasValidURL(req) {
		t.Error("expected false, got true")
	}
}
