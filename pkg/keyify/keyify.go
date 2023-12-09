package keyify

import (
	"crypto/rand"
	"encoding/base64"
)

func Keyify(length int, prefix ...string) (string, error) {
	randomBytes := make([]byte, length)

	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	var finalPrefix string
	if len(prefix) > 0 {
		finalPrefix = prefix[0]
	}

	key := finalPrefix + base64.URLEncoding.EncodeToString(randomBytes)[:length]

	return key, nil
}
