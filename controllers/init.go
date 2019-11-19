package controllers

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

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
