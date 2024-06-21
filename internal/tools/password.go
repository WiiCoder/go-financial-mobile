package tools

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

func CheckPassword(content, encrypted string) bool {
	return strings.EqualFold(MD5Encode(content), encrypted)
}

func MD5Encode(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
