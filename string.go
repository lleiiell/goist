package goist

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"strings"
	"time"
	"unicode"
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

func IsAscii(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] > unicode.MaxASCII {
			return false
		}
	}
	return true
}

func IsJsonEmpty(data []byte) bool {
	var jsonData interface{}
	err := json.Unmarshal(data, &jsonData)
	if err != nil {
		// fmt.Println("Failed to unmarshal JSON:", err)
		return true
	}

	switch jsonDataType := jsonData.(type) {
	case map[string]interface{}:
		return len(jsonDataType) == 0
	case []interface{}:
		return len(jsonDataType) == 0
	default:
		return true
	}
}

func isAscii(s string) bool {
	for _, c := range s {
		if c > unicode.MaxASCII {
			return false
		}
	}
	return true
}

func stringDiff(a, b string) (same string, diffA string, diffB string) {

	ba := []rune(a)
	bb := []rune(b)

	i := 0
	for ; i < len(ba) && i < len(bb); i++ {
		if ba[i] != bb[i] {
			break
		}

		same += string(ba[i])

	}

	diffA = string(ba[i:])
	diffB = string(bb[i:])

	return
}
