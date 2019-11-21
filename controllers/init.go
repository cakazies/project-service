package controllers

import (
	"crypto/rand"
	"encoding/base64"
	b64 "encoding/base64"
	"errors"
	"os"
	"project-service/models"
	"project-service/utils"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func GetToken(c *gin.Context) {
	var resp models.Rest

	result := GenerateToken()
	resp.Code = 200
	resp.Message = "Success"
	c.JSON(resp.Code, gin.H{"response": resp, "token": result})
	return
}

func GenerateToken() string {
	hmac := utils.Generate()
	encHmac := b64.StdEncoding.EncodeToString([]byte(hmac))
	hour, _ := strconv.Atoi(os.Getenv("TIME_EXP"))
	timein := time.Now().Local().Add(time.Hour * time.Duration(hour)).Format("2006-01-02")
	encTime := b64.StdEncoding.EncodeToString([]byte(timein))
	join := encHmac + "|" + generateRandom() + "|" + encTime + "|" + generateRandom()
	encJoin := b64.StdEncoding.EncodeToString([]byte(join))
	result := strings.ReplaceAll(encJoin, "=", "")
	return result
}

func generateRandom() string {
	randomBytes := make([]byte, 16)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(randomBytes)
}

// CheckingString function for checking params request string
func CheckingString(param string, c *gin.Context) (string, error) {
	url := c.Request.URL.String()
	if param == "" {
		param = url[19:29]
	}

	if param == "" {
		return "", errors.New("This Parameter is null")
	}
	return param, nil
}

// CheckingString function for checking params request int
func CheckingInt(param string) (int, error) {
	if param == "" {
		return 0, errors.New("This params doesn't exist")
	}

	result, err := strconv.Atoi(param)
	if err != nil {
		return 0, errors.New("This params not number")
	}

	if result <= 0 {
		return 0, errors.New("This params <= '0'")
	}

	return result, nil
}

// Response function for return error
func ErrorsResponse(code int, message string, c *gin.Context) {
	resp := map[string]string{"error": message, "code": strconv.Itoa(code), "status": "failed"}
	c.JSON(code, resp)
	c.Abort()
}

// this function for process limit and offset
func limitOffset(c *gin.Context) (string, string) {
	var limit string
	var offset string

	query := c.Request.URL.Query()
	if query["limit"] != nil {
		limits := query["limit"]
		_, err := strconv.Atoi(limits[0])
		limit = limits[0]
		if limits[0] == "" || err != nil {
			limit = "10"
		}
	} else {
		limit = "10"

	}

	if query["offset"] != nil {
		offsets := query["offset"]
		_, err := strconv.Atoi(offsets[0])
		offset = offsets[0]
		if offsets[0] == "" || err != nil {
			offset = "0"
		}
	} else {
		offset = "0"
	}

	return limit, offset
}

func orderBy(c *gin.Context) (string, string) {
	var shortBy string
	var shortValue string

	query := c.Request.URL.Query()
	if query["shortby"] != nil {
		short := query["shortby"]
		shortBy = short[0]
		value := strings.ToLower(shortBy)
		if value == "name" || value == "goal" || value == "start_period" {
			shortBy = value
		} else {
			shortBy = "id"
		}
	} else {
		shortBy = "id"
	}

	if query["shortvalue"] != nil {
		short := query["shortvalue"]
		shortValue = short[0]
		value := strings.ToLower(shortValue)
		if value == "desc" || value == "asc" {
			shortValue = strings.ToUpper(shortValue)
		} else {
			shortValue = "DESC"
		}
	} else {
		shortValue = "DESC"
	}
	return shortBy, shortValue
}

func where(c *gin.Context) string {
	var params string

	query := c.Request.URL.Query()
	if query["query"] != nil {
		short := query["query"]
		params = short[0]
	} else {
		params = ""
	}

	return params
}
