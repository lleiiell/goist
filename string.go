package goist

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

func RandStr(n int) string {
	lowerLetterRunes := []rune("0123456789abcdefghijklmnopqrstuvwxyz")

	b := make([]rune, n)
	for i := range b {
		b[i] = lowerLetterRunes[Rand().Intn(len(lowerLetterRunes))]
	}

	return string(b)
}

func IsStrEmpty(str string) bool {
	return strings.TrimSpace(str) == ""
}

func MD5(str string) string {
	c := md5.New()
	c.Write([]byte(str))

	return hex.EncodeToString(c.Sum(nil))
}
