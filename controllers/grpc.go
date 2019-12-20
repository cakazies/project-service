package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/cakazies/project-service/configs"
	rpc "github.com/cakazies/project-service/grpc"
	"github.com/cakazies/project-service/models"

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

	id, _ := strconv.Atoi(idproject)

	// check id project
	_, err = projectModel.GetProject(idproject)
	if err != nil && err.Error() == "record not found" {
		return fmt.Errorf("This ID project %s doesn't exist", idproject)
	}

	currentDate := time.Now().Format("2006-01-02 15:04:05")
	if all.Project.Status == "FINISHED" {
		all.ProjectDetail.FinishedDate = currentDate
	} else if all.Project.Status == "FAILED" {
		all.ProjectDetail.FailedDate = currentDate
	} else if all.Project.Status == "ON GOING" {
		all.ProjectDetail.OngoingDate = currentDate
	} else if all.Project.Status == "OVERDUE" {
		all.ProjectDetail.OverdueDate = currentDate
	}

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

// GetGallery function for get detail per project and return byte
func (GrpcRoute) GetGallery(projectID string) ([]byte, error) {
	projectModel := models.ProjectModels{}

	// check id project
	_, err := projectModel.GetProject(projectID)
	if err != nil && err.Error() == "record not found" {
		return nil, fmt.Errorf("This ID project %s doesn't exist", projectID)
	}

	result, err := projectModel.GetGallery(projectID)
	if err != nil {
		fmt.Println(err)
	}
	data, err := json.Marshal(result)
	return data, nil
}

// CreateGallery function for get detail per project and return byte
func (GrpcRoute) CreateGallery(data string) ([]byte, error) {
	var all models.ProjectGallery

	v := validator.New()
	err := json.Unmarshal([]byte(data), &all)
	projectModel := models.ProjectModels{}

	err = v.Struct(all)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			if e != nil {
				return nil, errors.New("Data `" + e.Field() + "` doesn't exist")
			}
		}
	}
	projectID := strconv.Itoa(all.ProjectID)
	// check id project
	_, err = projectModel.GetProject(projectID)
	if err != nil && err.Error() == "record not found" {
		return nil, fmt.Errorf("This ID project %s doesn't exist", projectID)
	}

	// insert project gallery
	err = projectModel.InsertGallery(all)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// UpdateGallery function for get detail per project and return byte
func (GrpcRoute) UpdateGallery(id int, data string) error {
	var all models.ProjectGallery

	err := json.Unmarshal([]byte(data), &all)
	projectModel := models.ProjectModels{}

	// insert project gallery
	err = projectModel.UpdateGallery(id, all)
	if err != nil {
		return err
	}

	return nil
}
