package models

import "project-service/configs"

type (
	CatgoryModel struct{}
)

func (c *CatgoryModel) GetCategories(addWhere string, limit string, offset string) ([]Category, error) {
	result := make([]Category, 0)

	err := configs.GetDB.Model(&result).Where(addWhere).Limit(limit).Offset(offset).Find(&result).Error
	return result, err
}

func (c *CatgoryModel) GetCategory(id int) (Category, error) {
	var result Category
	err := configs.GetDB.Model(&result).Where("id = ?", id).Find(&result).Error
	return result, err
}
