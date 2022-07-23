package util

import (
	"crypto/md5"
	"encoding/hex"
)

// Md5 cz
func Md5(msg string) string {
	h := md5.New()
	h.Write([]byte(msg))
	return hex.EncodeToString(h.Sum(nil))
}
