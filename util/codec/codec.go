package codec

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// Md5实现1
func Md5Hex(in string) string {
	if in == "" {
		return ""
	}
	hash := md5.New()
	hash.Write([]byte(in))
	return hex.EncodeToString(hash.Sum(nil))
}

// Md5实现2
func Md5Hex2(in string) string {
	if in == "" {
		return ""
	}

	return fmt.Sprintf("%x", md5.Sum([]byte(in)))
}

// Sha1实现
func Sha1Hex(in string) string {
	if in == "" {
		return ""
	}
	hash := sha1.New()
	hash.Write([]byte(in))
	return hex.EncodeToString(hash.Sum(nil))
}

// Sha256实现
func Sha256Hex(in string) string {
	if in == "" {
		return ""
	}
	hash := sha256.New()
	hash.Write([]byte(in))
	return hex.EncodeToString(hash.Sum(nil))
}
