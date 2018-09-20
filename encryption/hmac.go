package encryption

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

// CheckMACByString ...
func CheckMACByString(messageMac string, message []byte, key []byte) bool {
	mac := hmac.New(sha256.New, key)
	mac.Write(message)
	if hex.EncodeToString(mac.Sum(nil)) != messageMac {
		return false
	}
	return true
}

// DEPRECATED ...
func CheckMAC(message, messageMAC, key []byte) bool {
	mac := hmac.New(sha256.New, key)
	mac.Write(message)
	expectedMAC := mac.Sum(nil)
	return hmac.Equal(messageMAC, expectedMAC)
}
