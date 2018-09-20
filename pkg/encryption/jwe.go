package encryption

import "gopkg.in/square/go-jose.v2"

func DecryptJWE(words string, key []byte) (*string, error) {
	jwe, err := jose.ParseEncrypted(words)
	if err != nil {
		return nil, err
	}
	res, err := jwe.Decrypt(key)
	if err != nil {
		return nil, err
	}
	str := string(res)
	return &str, nil
}
