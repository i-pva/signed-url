package url

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"strconv"
	"time"
)

var SecretKey []byte

// Determine if the given request has a valid signature.
func HasValidURL(request *http.Request) bool {
	return hasCorrectSignature(request) && signatureHasNotExpired(request)
}

// Determine if the signature from the given request matches the URL.
func hasCorrectSignature(request *http.Request) bool {
	url := request.URL

	values := url.Query()
	signature := values.Get("signature")
	values.Del("signature")
	url.RawQuery = values.Encode()

	return signature == hash([]byte(url.String()))
}

// Determine if the expires timestamp from the given request is not from the past.
func signatureHasNotExpired(request *http.Request) bool {
	url := request.URL
	expires := url.Query().Get("expires")
	unix, _ := strconv.Atoi(expires)
	if unix == 0 {
		return true
	}

	return int(time.Now().Unix()) < unix
}

// hash for generate signature from given url
func hash(url []byte) string {
	mac := hmac.New(sha256.New, SecretKey)
	mac.Write(url)
	byteArray := mac.Sum(nil)

	return hex.EncodeToString(byteArray)
}
