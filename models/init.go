package models

import "github.com/jinzhu/gorm"

type (
	Project struct {
		Name       string `json:"name,omitempty"`
		Goal       string `json:"goal,omitempty"`
		Price      string `json:"price,omitempty"`
		Risk       string `json:"risk,omitempty"`
		Duration   string `json:"duration,omitempty"`
		Rate       string `json:"rate,omitempty"`
		Cbenefit   string `json:"cbenefit,omitempty"`
		Status     string `json:"status,omitempty"`
		Scheme     string `json:"scheme,omitempty"`
		Hidden     bool   `json:"hidden,omitempty"`
		CategoryId int    `json:"category_id,omitempty"`
		CreatedBy  int    `json:"created_by,omitempty"`
		gorm.Model
	}

	ProjectDetail struct {
		ProjectId    int    `json:"project_id,omitempty"`
		Description  string `json:"description,omitempty"`
		Address      string `json:"address,omitempty"`
		StartPeriod  string `json:"start_period,omitempty"`
		FailedDate   string `json:"failed_date,omitempty"`
		OngoingDate  string `json:"ongoing_date,omitempty"`
		FinishedDate string `json:"finished_date,omitempty"`
		CityId       int    `json:"city_id,omitempty"`
		ConfirmRdb   int    `json:"confirm_rdb,omitempty"`
		OverdueDate  string `json:"overdue_date,omitempty"`
		gorm.Model
	}

	ProjectTImeline struct {
		ProjectId   int    `json:"project_id,omitempty"`
		Date        string `json:"date,omitempty"`
		Title       string `json:"title,omitempty"`
		Description string `json:"description,omitempty"`
		ShortDesc   string `json:"short_desc,omitempty"`
		Photo       string `json:"photo,omitempty"`
		CreatedBy   int    `json:"created_by,omitempty"`
		gorm.Model
	}

	ProjectTimelineGallery struct {
		TimelineId int    `json:"timeline_id,omitempty"`
		ImageUrl   string `json:"image_url,omitempty"`
		gorm.Model
	}

	ProjectDocument struct {
		ProjectId   int    `json:"project_id,omitempty"`
		Title       string `json:"title,omitempty"`
		Description string `json:"description,omitempty"`
		Document    string `json:"document,omitempty"`
		CreatedBy   int    `json:"created_by,omitempty"`
		gorm.Model
	}

	ProjectLog struct {
		ProjectId int    `json:"project_id,omitempty"`
		Data      string `json:"data,omitempty"`
		UpdatedBy string `json:"updated_by,omitempty"`
		gorm.Model
	}

	ProjectGallery struct {
		ProjectId   int    `json:"project_id,omitempty"`
		ImagesUrl   string `json:"images_url,omitempty"`
		Description string `json:"description,omitempty"`
		Status      int    `json:"status,omitempty"`
		CreatedBy   int    `json:"created_by,omitempty"`
		gorm.Model
	}

	Category struct {
		Name        string `json:"name,omitempty"`
		Description string `json:"description,omitempty"`
		Icon        string `json:"icon,omitempty"`
		Icon_retina string `json:"icon_retina,omitempty"`
		CreatedBy   int    `json:"created_by,omitempty"`
		gorm.Model
	}

	CategoryTypes struct {
		CategoryId int    `json:"category_id,omitempty"`
		Name       string `json:"name,omitempty"`
		CreatedBy  int    `json:"created_by,omitempty"`
		gorm.Model
	}

	Rest struct {
		Message string `json:"message,omitempty"`
		Code    int    `json:"code,omitempty"`
		Count   int    `json:"count,omitempty"`
	}
)
