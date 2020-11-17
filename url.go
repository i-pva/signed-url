package go_url

import (
	"fmt"
	"net/url"
	"strconv"
	"time"
)

type Url struct {
	Path string
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
		u.Set("expires", strconv.Itoa(int(delay)))
	}
	u.Set("signature", hash([]byte(u.String())))
	return u.String()
}

func (u *Url) Set(key, value string) {
	if u.Values == nil {
		u.Values = url.Values{}
	}
	u.Values.Set(key, value)
}

func (u *Url) Add(key, value string) {
	if u.Values == nil {
		u.Values = url.Values{}
	}
	u.Values.Add(key, value)
}

func (u *Url) String() string {
	return fmt.Sprintf("%s?%s", u.Path, u.Values.Encode())
}
