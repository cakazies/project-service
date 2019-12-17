package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"os"
)

var secretKey = os.Getenv("SECRET_SERVICE")
var message = os.Getenv("MESSAGE_SERVICE")

// Compare function for compare hmac
func Compare(messageQ []byte) bool {
	mac := hmac.New(sha256.New, []byte(secretKey))
	mac.Write([]byte(message))
	expectedMAC := mac.Sum(nil)
	return hmac.Equal(messageQ, expectedMAC)
}

// Generate function hmac
func Generate() string {
	mac := hmac.New(sha256.New, []byte(secretKey))
	mac.Write([]byte(message))
	expectedMAC := mac.Sum(nil)
	return string(expectedMAC)
}
