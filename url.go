package go_url

import (
	"net/url"
	"strconv"
	"time"
)

type Url struct {
	url.URL
	url.Values
}

// Create a signed url URL for a path route.
func (u *Url) Signed() string {
	return u.sign(0)
}

// Create a temporary signed url URL for a path route.
func (u *Url) TemporarySigned(expiration time.Duration) string {
	return u.sign(expiration)
}

func (u *Url) sign(expiration time.Duration) string {
	if expiration != 0 {
		delay := time.Now().Add(expiration).Unix()
		u.setQueryParameter("expires", strconv.Itoa(int(delay)))
	}
	u.setQueryParameter("signature", hash([]byte(u.String())))
	return u.String()
}

func (u *Url) SetQueryParameter(key, value string) {
	u.setQueryParameter(key, value)
}

func (u *Url) setQueryParameter(key, value string) {
	if u.Values == nil {
		u.Values = url.Values{}
	}
	u.Values.Set(key, value)
	u.RawQuery = u.Values.Encode()
}
