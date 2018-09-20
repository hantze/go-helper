package encryption

import (
	"crypto/sha256"
	"encoding/hex"
)

// HashSHA ...
func HashSHA(data string) string {
	hash := sha256.New()
	hash.Write([]byte(data))
	md := hash.Sum(nil)
	return hex.EncodeToString(md)
}
