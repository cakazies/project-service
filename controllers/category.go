package controllers

import (
	"fmt"
	"project-service/models"

	"github.com/gin-gonic/gin"
)

type (
	CategoryCtrl struct{}
)

func (cg *CategoryCtrl) GetCategories(c *gin.Context) {
	var resp models.Rest

	query := c.Request.URL.Query()
	params := ""

	for k, v := range query {
		if k == "rdbAccount" || k == "email" {
			params = params + fmt.Sprintf(" AND %s = '%s' ", k, v[0])
		}
		if k == "createdAt" {
			// params = params + fmt.Sprintf(" AND DATE_FORMAT(`%s`, '%Y-%m-%D') = %s ", k, v[0])
			params = params + " AND DATE_FORMAT(" + k + ", '%Y-%m-%d') = '" + v[0] + "' "
		}
	}

	limit, offset := limitOffset(c)
	catModel := models.CatgoryModel{}
	result, err := catModel.GetCategories(params, limit, offset)
	if err != nil {
		ErrorsResponse(404, err.Error(), c)
		return
	}

	resp.Code = 200
	resp.Message = "Success"
	c.JSON(resp.Code, gin.H{"response": resp, "data": result})
	return
}

func (cg *CategoryCtrl) GetCategory(c *gin.Context) {
	var resp models.Rest

	param := c.Param("id")
	id, err := CheckingInt(param)
	if err != nil {
		ErrorsResponse(404, err.Error(), c)
		return
	}

	catModel := models.CatgoryModel{}
	result, err := catModel.GetCategory(id)
	if err != nil {
		ErrorsResponse(404, err.Error(), c)
		return
	}

	resp.Code = 200
	resp.Message = "Success"
	c.JSON(resp.Code, gin.H{"response": resp, "data": result})
	return
}
