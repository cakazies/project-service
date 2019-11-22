package controllers

import (
	"encoding/json"
	"fmt"
	"project-service/configs"
	rpc "project-service/grpc"
	"project-service/models"
	"strconv"

	"gopkg.in/go-playground/validator.v9"
)

func init() {
	configs.InitASCII()
	configs.Connect()
}

type (
	// GrpcRoute struct for set this function
	GrpcRoute struct{}

	sendMsql struct {
		userBY int
		data   models.ProjectAll
	}
)

// GetProjects function for get project
func (GrpcRoute) GetProjects(paging *rpc.Pagination) []byte {
	bniModel := models.ProjectModels{}
	orderby := paging.ShortBy + " " + paging.Shortvalue
	result, err := bniModel.GetProjects(paging.Query, paging.Limit, paging.Offset, orderby)
	if err != nil {
		fmt.Println(err)
	}
	data, err := json.Marshal(result)
	return data
}

// GetProject function for get detail per project and return byte
func (GrpcRoute) GetProject(id string) []byte {
	bniModel := models.ProjectModels{}
	result, err := bniModel.GetProject(id)
	if err != nil {
		fmt.Println(err)
	}
	data, err := json.Marshal(result)
	return data
}

// Create function for get detail per project and return byte
func (GrpcRoute) Create(data string) []byte {
	var all models.ProjectAll
	var resp models.Rest
	var msq sendMsql

	v := validator.New()
	err := json.Unmarshal([]byte(data), &all)
	projectModel := models.ProjectModels{}

	err = v.Struct(all)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			if e != nil {
				resp.Code = 400
				resp.Message = "Data `" + e.Field() + "` doesn't exist"
				data, _ := json.Marshal(resp)
				return data
			}
		}
	}

	// insert project and get ID project
	dataProject, err := projectModel.InsertProject(all.Project)
	if err != nil {

	}

	all.ProjectDetail.ProjectID = int(dataProject.ID)
	err = projectModel.InsertProjectDetail(all.ProjectDetail)

	msq.userBY = all.CreatedBy
	msq.data = all
	dataSend, err := json.Marshal(&msq)
	configs.Publish("PROJECT_MESSAGING", dataSend)

	resp.Code = 200
	resp.Message = "Success"

	return nil
}

// Edit function for Update data project
func (GrpcRoute) Edit(idproject string, data string) error {
	var all models.ProjectAll

	err := json.Unmarshal([]byte(data), &all)
	projectModel := models.ProjectModels{}

	if all.Hidden != true {
		all.Hidden = false
	}
	fmt.Println(all)
	id, _ := strconv.Atoi(idproject)
	// insert project and get ID project
	err = projectModel.UpdateProject(id, all.Project)
	if err != nil {
		return err
	}

	err = projectModel.UpdateProjectDetail(id, all.ProjectDetail)
	if err != nil {
		return err
	}

	return nil
}
