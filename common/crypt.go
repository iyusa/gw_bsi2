package common

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

// Sha256 calculate Sha256 to string
func Sha256(src []byte) string {
	sum := sha256.Sum256(src)
	return fmt.Sprintf("%x", sum)
}

// Hmac256 enc
func Hmac256(src, secret []byte) string {
	h := hmac.New(sha256.New, secret)
	h.Write(src)
	digest := h.Sum(nil)
	return fmt.Sprintf("%x", digest) // bca
}

// DigestHMAC256 return base64 hash signature
// ref: https://www.php2golang.com/method/function.hash-hmac.html
func DigestHMAC256(ciphertext, key []byte) string {
	mac := hmac.New(sha256.New, key)
	mac.Write([]byte(ciphertext))
	hmac := mac.Sum(nil)
	return base64.StdEncoding.EncodeToString(hmac)
}

func CreateSignature(source string) string {
	h := sha256.New()
	h.Write([]byte(source))
	return hex.EncodeToString(h.Sum(nil))
}
