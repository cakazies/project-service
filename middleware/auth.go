package middleware

import (
	b64 "encoding/base64"
	"errors"
	"strings"
	"time"

	"github.com/cakazies/project-service/models"
	"github.com/cakazies/project-service/utils"

	"github.com/gin-gonic/gin"
)

// CheckMiddleware function for checking middleware
func CheckMiddleware(c *gin.Context) {
	noAuthPath := []string{"/api/v1/token"}
	requestPath := c.Request.URL.Path
	for _, path := range noAuthPath {
		if path == requestPath {
			c.Next()
			return
		}
	}

	token := c.Request.Header.Get("token")
	if token == "" {
		respondWithError(c, 404, "Token is Empty")
		return
	}

	ok, err := generate(token)
	if err != nil || !ok {
		respondWithError(c, 401, err.Error())
		return
	}

	c.Next()
}

// respondWithError function for response in middleware if error
func respondWithError(c *gin.Context, code int, message string) {
	var resp models.Rest
	resp.Code = code
	resp.Message = message
	c.AbortWithStatusJSON(code, gin.H{"response": resp})
}

func generate(token string) (bool, error) {
	decode, err := b64.StdEncoding.DecodeString(token)
	if err != nil {
		return false, errors.New("This Token Invalid")
	}
	data := strings.Split(string(decode), "|")
	encHmac := data[0]
	encTime := data[2]

	hmac, _ := b64.StdEncoding.DecodeString(encHmac)
	times, _ := b64.StdEncoding.DecodeString(encTime)

	myDate, _ := time.Parse("2006-01-02 15:04", (string(times) + " 00:00"))
	diff := myDate.Sub(time.Now())

	if diff < 0 {
		return false, errors.New("This Token Expired")
	}

	ok := utils.Compare(hmac)
	if !ok {
		return false, errors.New("This Token Invalid")
	}

	return true, nil
}
