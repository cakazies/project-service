package models

import (
	"github.com/jinzhu/gorm"
)

type (
	// Project struct for table project
	Project struct {
		Name       string `json:"name" validate:"required"`
		Goal       string `json:"goal" validate:"required"`
		Price      string `json:"price" validate:"required"`
		Risk       string `json:"risk" validate:"required"`
		Duration   string `json:"duration" validate:"required"`
		Rate       string `json:"rate" validate:"required"`
		Cbenefit   string `json:"cbenefit" validate:"required"`
		Status     string `json:"status" validate:"required"`
		Scheme     string `json:"scheme" validate:"required"`
		Hidden     bool   `json:"hidden" validate:"required"`
		CategoryID int    `json:"category_id" validate:"required"`
		CreatedBy  int    `json:"created_by" validate:"required"`
		gorm.Model
	}

	// ProjectDetail struct for table project detail
	ProjectDetail struct {
		ProjectID    int    `json:"project_id,omitempty"`
		Description  string `json:"description" validate:"required"`
		Address      string `json:"address"`
		StartPeriod  string `json:"start_period" validate:"required"`
		FailedDate   string `json:"failed_date,omitempty"`
		OngoingDate  string `json:"ongoing_date,omitempty"`
		FinishedDate string `json:"finished_date,omitempty"`
		CityID       int    `json:"city_id" validate:"required"`
		ConfirmRdb   int    `json:"confirm_rdb,omitempty"`
		OverdueDate  string `json:"overdue_date,omitempty"`
		gorm.Model
	}

	// ProjectTImeline struct for table project timeline
	ProjectTImeline struct {
		ProjectID   int    `json:"project_id,omitempty"`
		Date        string `json:"date,omitempty"`
		Title       string `json:"title,omitempty"`
		Description string `json:"description,omitempty"`
		ShortDesc   string `json:"short_desc,omitempty"`
		Photo       string `json:"photo,omitempty"`
		CreatedBy   int    `json:"created_by,omitempty"`
		gorm.Model
	}

	// ProjectTimelineGallery struct for table project timeline gallery
	ProjectTimelineGallery struct {
		TimelineID int    `json:"timeline_id,omitempty"`
		ImageURL   string `json:"image_url,omitempty"`
		gorm.Model
	}

	// ProjectDocument struct for table project document
	ProjectDocument struct {
		ProjectID   int    `json:"project_id,omitempty"`
		Title       string `json:"title,omitempty"`
		Description string `json:"description,omitempty"`
		Document    string `json:"document,omitempty"`
		CreatedBy   int    `json:"created_by,omitempty"`
		gorm.Model
	}

	// ProjectLog struct for table project log
	ProjectLog struct {
		ProjectID int    `json:"project_id,omitempty"`
		Data      string `json:"data,omitempty"`
		UpdatedBy string `json:"updated_by,omitempty"`
		gorm.Model
	}

	// ProjectGallery struct for table project gallery
	ProjectGallery struct {
		ProjectID   int    `json:"project_id,omitempty"`
		ImagesURL   string `json:"images_url,omitempty"`
		Description string `json:"description,omitempty"`
		Status      int    `json:"status,omitempty"`
		CreatedBy   int    `json:"created_by,omitempty"`
		gorm.Model
	}

	// Category struct for table category
	Category struct {
		Name        string `json:"name,omitempty"`
		Description string `json:"description,omitempty"`
		Icon        string `json:"icon,omitempty"`
		IconRetina  string `json:"icon_retina,omitempty"`
		CreatedBy   int    `json:"created_by,omitempty"`
		gorm.Model
	}

	// CategoryTypes struct for table category_types
	CategoryTypes struct {
		CategoryID int    `json:"category_id,omitempty"`
		Name       string `json:"name,omitempty"`
		CreatedBy  int    `json:"created_by,omitempty"`
		gorm.Model
	}

	// Rest struct for response
	Rest struct {
		Message string `json:"message,omitempty"`
		Code    int    `json:"code,omitempty"`
		Count   int    `json:"count,omitempty"`
	}

	// ProjectAll struct for getall project
	ProjectAll struct {
		Project
		ProjectDetail
	}
)
