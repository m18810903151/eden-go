package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5(str string) string {
	bytes := []byte(str)
	hash := md5.Sum(bytes)
	return hex.EncodeToString(hash[:])
}

func MD5Salt(str string, salt string) string {
	bytes := []byte(str + salt)
	hash := md5.Sum(bytes)
	return hex.EncodeToString(hash[:])
}
