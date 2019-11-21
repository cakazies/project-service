package controllers

import (
	"project-service/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
)

type (
	ProjectCtr struct{}
)

func (p *ProjectCtr) Create(c *gin.Context) {
	var all models.ProjectAll
	var resp models.Rest

	v := validator.New()
	c.BindJSON(&all)
	projectModel := models.ProjectModels{}

	err := v.Struct(all)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			if e != nil {
				resp.Code = 400
				resp.Message = "Data `" + e.Field() + "` doesn't exist"
				c.JSON(resp.Code, gin.H{"response": resp})
				return
			}
		}
	}

	// insert project and get ID project
	dataProject, err := projectModel.InsertProject(all.Project)
	if err != nil {
		ErrorsResponse(404, err.Error(), c)
		return
	}

	all.ProjectDetail.ProjectID = int(dataProject.ID)
	err = projectModel.InsertProjectDetail(all.ProjectDetail)
	if err != nil {
		ErrorsResponse(404, err.Error(), c)
		return
	}

	resp.Code = 200
	resp.Message = "Success"
	c.JSON(resp.Code, gin.H{"response": resp, "project_id": dataProject.ID})
	return
}

func (p *ProjectCtr) Update(c *gin.Context) {
	var all models.ProjectAll
	var resp models.Rest

	c.BindJSON(&all)
	id, _ := strconv.Atoi(c.Param("id"))
	projectModel := models.ProjectModels{}

	// insert project and get ID project
	err := projectModel.UpdateProject((uint)(id), all.Project)
	if err != nil {
		ErrorsResponse(404, err.Error(), c)
		return
	}

	err = projectModel.UpdateProjectDetail(id, all.ProjectDetail)
	if err != nil {
		ErrorsResponse(404, err.Error(), c)
		return
	}

	resp.Code = 200
	resp.Message = "Success"
	c.JSON(resp.Code, gin.H{"response": resp, "project_id": id})
	return
}

func (p *ProjectCtr) GetProjects(c *gin.Context) {
	var resp models.Rest

	limit, offset := limitOffset(c)
	shortby, shortValue := orderBy(c)
	addwhere := where(c)
	orderBy := shortby + " " + shortValue

	bniModel := models.ProjectModels{}
	result, err := bniModel.GetProjects(addwhere, limit, offset, orderBy)
	if err != nil {
		ErrorsResponse(404, err.Error(), c)
		return
	}

	resp.Code = 200
	resp.Message = "Success"
	c.JSON(resp.Code, gin.H{"response": resp, "data": result})
	return
}

func (p *ProjectCtr) GetProject(c *gin.Context) {
	var resp models.Rest

	id := c.Param("id")
	id, err := CheckingString(id, c)
	if err != nil {
		ErrorsResponse(404, err.Error(), c)
		return
	}

	bniModel := models.ProjectModels{}
	result, err := bniModel.GetProject(id)
	if err != nil {
		ErrorsResponse(404, err.Error(), c)
		return
	}

	resp.Code = 200
	resp.Message = "Success"
	c.JSON(resp.Code, gin.H{"response": resp, "data": result})
	return
}
