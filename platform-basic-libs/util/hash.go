package util

import (
	"crypto"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
)

// Hash 散列函数 返回空值则为错误
func Hash(ht crypto.Hash, b []byte) []byte {
	switch ht {
	case crypto.MD5:
		h := md5.New()
		h.Write(b)
		return h.Sum(nil)
	case crypto.SHA1:
	case crypto.SHA256:
	}
	return nil
}

// HashHex can hash and encode to string
func HashHex(ht crypto.Hash, b []byte) string {
	return hex.EncodeToString(Hash(ht, b))
}

// MD5Hash md5
func MD5Hash(b []byte) []byte {
	h := md5.New()
	h.Write(b)
	return h.Sum(nil)
}

// MD5HexHash md5 and encode to string
func MD5HexHash(b []byte) string {
	h := md5.New()
	h.Write(b)
	return hex.EncodeToString(h.Sum(nil))
}

// SHA256HexHash sha256 and encode to string
func SHA256HexHash(b []byte) string {
	h := sha256.New()
	h.Write(b)
	return hex.EncodeToString(h.Sum(nil))
}

// SHA1HexHash sha1 and encode to string
func SHA1HexHash(b []byte) string {
	h := sha1.New()
	h.Write(b)
	return hex.EncodeToString(h.Sum(nil))
}

// HMACSHA1Hash hmac-sha1
func HMACSHA1Hash(src, key []byte) []byte {
	h := hmac.New(sha1.New, key)
	h.Write(src)
	return h.Sum(nil)
}

// HMACSHA1HexHash hmac-sha1 and encode to string
func HMACSHA1HexHash(src, key []byte) string {
	h := hmac.New(sha1.New, key)
	h.Write(src)
	return hex.EncodeToString(h.Sum(nil))
}

func HmacSha256(data string, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}
