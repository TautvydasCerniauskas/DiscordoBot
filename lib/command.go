package lib

import (
	"crypto/md5"
	"encoding/hex"
)

// MD5Hash return md5 hash of a string
func MD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
