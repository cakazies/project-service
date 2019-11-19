package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"os"
)

var secretKey = os.Getenv("SECRET_SERVICE")
var message = os.Getenv("MESSAGE_SERVICE")

func Compare(messageQ string) bool {
	mac := hmac.New(sha256.New, []byte(secretKey))
	mac.Write([]byte(message))
	expectedMAC := mac.Sum(nil)

	mac1 := hmac.New(sha256.New, []byte(secretKey))
	mac1.Write([]byte(messageQ))
	messageMAC := mac1.Sum(nil)
	return hmac.Equal([]byte(messageMAC), expectedMAC)
}
