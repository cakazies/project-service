package models

import "github.com/crowdeco/project-service/configs"

type (
	// CatgoryModel struct for this setting class
	CatgoryModel struct{}
)

// GetCategories function for get categories
func (c *CatgoryModel) GetCategories(addWhere string, limit string, offset string) ([]Category, error) {
	result := make([]Category, 0)

	err := configs.GetDB.Model(&result).Where(addWhere).Limit(limit).Offset(offset).Find(&result).Error
	return result, err
}

// GetCategory function for get category
func (c *CatgoryModel) GetCategory(id int) (Category, error) {
	var result Category
	err := configs.GetDB.Model(&result).Where("id = ?", id).Find(&result).Error
	return result, err
}
