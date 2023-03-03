package goist

import (
    "crypto/md5"
    "encoding/hex"
)

func MD5(str string) string {
    c := md5.New()
    c.Write([]byte(str))

    return hex.EncodeToString(c.Sum(nil))
}
