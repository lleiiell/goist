package goist

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/google/uuid"
	"strings"
	"time"
)

// UuidSimple uuid without special characters
func UuidSimple() string {
	escaper := strings.NewReplacer("-", "90", "_", "91")

	return escaper.Replace(Uuid22())
}

// Uuid22 uuid of length 22
func Uuid22() string {
	id := Uuid32()

	bt, _ := hex.DecodeString(id)

	return base64.RawURLEncoding.EncodeToString(bt)
}

// Uuid32 uuid without dash
func Uuid32() string {
	return strings.Replace(uuid.New().String(), "-", "", -1)
}

// RandStr random string
// n custom length
func RandStr(n int) string {
	lowerLetterRunes := []rune("0123456789abcdefghijklmnopqrstuvwxyz")

	b := make([]rune, n)
	for i := range b {
		b[i] = lowerLetterRunes[Rand().Intn(len(lowerLetterRunes))]
		time.Sleep(1 * time.Nanosecond)
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

// Token generate token
func Token() string {
	return MD5(fmt.Sprintf("%d%d", Rand().Int63(), time.Now().UnixNano()))
}
