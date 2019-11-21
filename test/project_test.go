package test

import (
	"encoding/json"
	"fmt"
	"net/http"
	ctr "project-service/controllers"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/local/testify/assert"
)

var (
	projectCtr = ctr.ProjectCtr{}
	token      = ctr.GenerateToken()
)

type (
	testCase struct {
		name         string
		expectedData string
		expectedCode int
		path         string
		query        string
		token        string
		username     string
		password     string
		email        string
		requestData  string
		router       func(*gin.Context)
	}

	Response struct {
		Response Rest                   `json:"response"`
		Data     map[string]interface{} `json:"data,omitempty"`
		// Data Data `json:"data,omitempty"`
	}

	Data struct {
		Id int `json:"id"`
	}

	Rest struct {
		Message string `json:"message,omitempty"`
		Code    int    `json:"code,omitempty"`
	}
)

func init() {
	// Load ENV
	loadEnv()
}

func TestToken(t *testing.T) {
	tasks := []testCase{
		{
			name:         "Get Token with Auth valid",
			expectedData: "Success",
			expectedCode: http.StatusOK,
			path:         "/api/v1/token",
			query:        "",
			token:        token,
			router:       ctr.GetToken,
			requestData:  "",
		},
		{
			name:         "Get Token with Auth nil",
			expectedData: "Token is Empty",
			expectedCode: http.StatusNotFound,
			path:         "/api/v1/token",
			query:        "",
			router:       ctr.GetToken,
			requestData:  "",
			token:        "",
		},
		{
			name:         "Get Token with Auth invalid",
			expectedData: "This Token Invalid",
			expectedCode: http.StatusUnauthorized,
			path:         "/api/v1/token",
			query:        "",
			router:       ctr.GetToken,
			requestData:  "",
			token:        "s" + token,
		},
	}

	for _, tc := range tasks {
		t.Run(tc.name, func(t *testing.T) {
			buf := RequestPOST(tc)
			var respData Response
			if err := json.Unmarshal(buf, &respData); err != nil {
				t.Error("Can not parsing response testing. Error :", err)
			}

			assert.Equal(t, tc.expectedCode, respData.Response.Code, "Expedted Code is Wrong")
			getData := fmt.Sprintf("%v", respData.Response.Message)
			assert.Equal(t, tc.expectedData, getData, "Expedted Data is Wrong")
		})
	}
}

func GetDetail(t *testing.T) {
	tasks := []testCase{
		{
			name:         "Get Detail Product with ID params",
			expectedData: "Success",
			expectedCode: http.StatusOK,
			path:         "/api/v1/project/4",
			query:        "",
			token:        token,
			router:       ctr.GetToken,
			requestData:  "",
		},
		{
			name:         "Get Detail Product with ID nil",
			expectedData: "Token is Empty",
			expectedCode: http.StatusNotFound,
			path:         "/api/v1/project/",
			query:        "",
			router:       ctr.GetToken,
			requestData:  "",
			token:        token,
		},
		{
			name:         "Get Detail Product with params not integer",
			expectedData: "This Token Invalid",
			expectedCode: http.StatusUnauthorized,
			path:         "/api/v1/token",
			query:        "",
			router:       ctr.GetToken,
			requestData:  "",
			token:        token,
		},
	}

	for _, tc := range tasks {
		t.Run(tc.name, func(t *testing.T) {
			buf := RequestPOST(tc)
			var respData Response
			if err := json.Unmarshal(buf, &respData); err != nil {
				t.Error("Can not parsing response testing. Error :", err)
			}

			assert.Equal(t, tc.expectedCode, respData.Response.Code, "Expedted Code is Wrong")
			getData := fmt.Sprintf("%v", respData.Response.Message)
			assert.Equal(t, tc.expectedData, getData, "Expedted Data is Wrong")
		})
	}
}
