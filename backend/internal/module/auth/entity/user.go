package entity

import (
	"celestara.com/hris-api/internal/shared/model"
)

// User
type User struct {
	model.TenantBaseModel
	Email        string `gorm:"type:varchar(255);not null;uniqueIndex:idx_company_email" json:"email"`
	PasswordHash string `gorm:"type:varchar(255);not null" json:"-"`
	Role         string `gorm:"type:varchar(50);not null;default:'employee'" json:"role"`
	IsActive     bool   `gorm:"default:true" json:"is_active"`
}