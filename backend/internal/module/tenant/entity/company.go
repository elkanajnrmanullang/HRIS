package entity

import (
	"celestara.com/hris-api/internal/shared/model"
)

// Company
type Company struct {
	model.BaseModel
	Name         string `gorm:"type:varchar(255);not null" json:"name"`
	Email        string `gorm:"type:varchar(255);uniqueIndex;not null" json:"email"`
	Phone        string `gorm:"type:varchar(50)" json:"phone"`
	Address      string `gorm:"type:text" json:"address"`
	LogoURL      string `gorm:"type:varchar(255)" json:"logo_url"`
	IsActive     bool   `gorm:"default:true" json:"is_active"`
	MaxEmployees int    `gorm:"default:50;not null" json:"max_employees"`
}