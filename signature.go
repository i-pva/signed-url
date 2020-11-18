package url

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

var Key []byte

// hash for generate signature from given url
func hash(url []byte) string {
	mac := hmac.New(sha256.New, Key)
	mac.Write(url)
	byteArray := mac.Sum(nil)

	return hex.EncodeToString(byteArray)
}
