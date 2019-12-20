package models

import (
	"fmt"

	"github.com/cakazies/project-service/configs"
)

var (
	// PROJECT for table project
	PROJECT = "projects"
	// TblProjectDetail for table project detail
	TblProjectDetail = "project_details"
)

type (
	// ProjectModels struct for this class
	ProjectModels struct{}
)

// GetProjects function for get all project and some parameter like limit offset orderby and the others
func (b *ProjectModels) GetProjects(addwhere string, limit string, offset string, orderby string) ([]ProjectAll, error) {
	result := make([]ProjectAll, 0)
	join := fmt.Sprintf("join %s on %s.id = %s.project_id ", TblProjectDetail, PROJECT, TblProjectDetail)
	selects := fmt.Sprintf("%s.*,%s.*", PROJECT, TblProjectDetail)
	order := fmt.Sprintf("%s.%s", PROJECT, orderby)
	where := `(name LIKE "%` + addwhere + `%" OR description LIKE "%` + addwhere + `%" OR address LIKE "%` + addwhere + `%")`

	err := configs.GetDB.Table(PROJECT).Select(selects).Joins(join).Where(where).Limit(limit).Offset(offset).Order(order).Find(&result).Error
	return result, err
}

// GetProject function for get detail project
func (b *ProjectModels) GetProject(id string) (ProjectAll, error) {
	var result ProjectAll
	join := fmt.Sprintf("join %s on %s.id = %s.project_id ", TblProjectDetail, PROJECT, TblProjectDetail)
	where := fmt.Sprintf("%s.id = ?", PROJECT)
	selects := fmt.Sprintf("%s.*,%s.*", PROJECT, TblProjectDetail)

	err := configs.GetDB.Table(PROJECT).Select(selects).Joins(join).Where(where, id).Find(&result).Error
	return result, err
}

// InsertProject function for insert project
func (b *ProjectModels) InsertProject(data Project) (Project, error) {
	var result Project
	err := configs.GetDB.Create(&data).Scan(&result).Error
	return result, err
}

// UpdateProject function for update project
func (b *ProjectModels) UpdateProject(id int, data Project) error {
	err := configs.GetDB.Table(PROJECT).Where("id = ?", id).Updates(&data).Error
	return err
}

// InsertProjectDetail function for insert project detail
func (b *ProjectModels) InsertProjectDetail(data ProjectDetail) error {
	err := configs.GetDB.Create(&data).Error
	return err
}

// UpdateProjectDetail function for update project detail
func (b *ProjectModels) UpdateProjectDetail(id int, data ProjectDetail) error {
	err := configs.GetDB.Table(TblProjectDetail).Where("project_id = ?", id).Updates(&data).Error
	return err
}

// GetGallery function for get detail project
func (b *ProjectModels) GetGallery(projectID string) ([]ProjectGallery, error) {
	result := make([]ProjectGallery, 0)

	err := configs.GetDB.Model(&result).Where("project_id = ?", projectID).Find(&result).Error
	return result, err
}

// InsertGallery function for insert project
func (b *ProjectModels) InsertGallery(data ProjectGallery) error {
	err := configs.GetDB.Create(&data).Error
	return err
}

// UpdateGallery function for update project detail
func (b *ProjectModels) UpdateGallery(id int, data ProjectGallery) error {
	err := configs.GetDB.Model(&data).Where("id = ?", id).Updates(&data).Error
	return err
}
