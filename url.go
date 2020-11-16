package go_url

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"github.com/google/uuid"
	"net/url"
	"strconv"
	"time"
)

type KeyResolver func() string

type Url struct {
	url.URL
	url.Values
	KeyResolver KeyResolver
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
	if u.KeyResolver == nil {
		u.KeyResolver = func() string {
			return uuid.New().String()
		}
	}

	if expiration != 0 {
		delay := time.Now().Add(expiration).Unix()
		u.Values.Set("expires", strconv.Itoa(int(delay)))
	}
	u.Values.Set("signature", hash(u.bytes(), []byte(u.KeyResolver())))
	u.RawQuery = u.Values.Encode()
	return u.String()
}

// Func bytes returns bytes of url
func (u *Url) bytes() []byte {
	return []byte(u.String())
}

func hash(url, key []byte) string {
	mac := hmac.New(sha256.New, key)
	mac.Write(url)
	byteArray := mac.Sum(nil)

	return hex.EncodeToString(byteArray)
}
