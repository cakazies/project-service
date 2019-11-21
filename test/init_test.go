package test

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	mw "project-service/middleware"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// this function about testing request POST in this Service
func RequestPOST(tc testCase) []byte {
	resp := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	jsonStr := []byte(tc.requestData)
	c, r := gin.CreateTestContext(resp)
	c.Request, _ = http.NewRequest(http.MethodPost, tc.path, bytes.NewBuffer(jsonStr))
	c.Request.Header.Set("token", tc.token)
	r.Use(mw.CheckMiddleware)
	r.POST(tc.path, tc.router)
	r.ServeHTTP(resp, c.Request)

	return resp.Body.Bytes()
}

func loadEnv() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file error : ", err)
	}
}
