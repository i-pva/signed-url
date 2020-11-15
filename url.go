package go_url

import (
	"net/url"
	"time"
)

type Url struct {
	url.URL
}

//Create a signed route URL for a path route.
func (u *Url) SignedRoute() string {
	return ""
}

// Create a temporary signed route URL for a path route.
func (u *Url) TemporarySignedRoute(expiration time.Duration) string {
	return ""
}
