package test

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"testing"

	project "github.com/cakazies/project-service/grpc"
)

var (
	pro = serviceProject()
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
		// router       func(*gin.Context)
	}

	Response struct {
		Response Rest                   `json:"response"`
		Data     map[string]interface{} `json:"data,omitempty"`
		// Data Data `json:"data,omitempty"`
	}

	Data struct {
		ID int `json:"id"`
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
			name:         "Get Token BNI with Auth valid",
			expectedData: "Bearer",
			expectedCode: http.StatusOK,
			path:         "/api/v1/token",
			query:        "",
			token:        os.Getenv("TEST_TOKEN"),
			requestData:  "10",
		},
		// {
		// 	name:         "Get Token BNI with Auth nil",
		// 	expectedData: "<nil>",
		// 	expectedCode: http.StatusUnauthorized,
		// 	path:         "/api/v1/token",
		// 	query:        "",
		// 	requestData:  "",
		// 	token:        "",
		// },
		// {
		// 	name:         "Get Token BNI with Auth wrong",
		// 	expectedData: "<nil>",
		// 	expectedCode: http.StatusUnauthorized,
		// 	path:         "/api/v1/token",
		// 	query:        "",
		// 	requestData:  "",
		// 	token:        "s" + os.Getenv("TEST_TOKEN"),
		// },
	}

	for _, tc := range tasks {
		t.Run(tc.name, func(t *testing.T) {
			var result project.Project

			idProject, err := strconv.Atoi(tc.requestData)
			result.Id = int32(idProject)
			res3, err := pro.Detail(context.Background(), &result)
			if err != nil {
				log.Println(err.Error())
			}
			err = json.Unmarshal([]byte(res3.Data), &result)
			fmt.Println(err)
			fmt.Println(result)

			// fmt.Println(tc.name)
			// buf := RequestPOST(tc)
			// var respData Response
			// if err := json.Unmarshal(buf, &respData); err != nil {
			// 	t.Error("Can not parsing response testing. Error :", err)
			// }

			// assert.Equal(t, tc.expectedCode, "207", "Expedted Code is Wrong")
			// getData := fmt.Sprintf("%v", respData.Data["token_type"])
			// assert.Equal(t, tc.expectedData, getData, "Expedted Data is Wrong")
		})
	}
}
