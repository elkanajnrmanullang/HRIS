package entity

import (
	"time"

	"celestara.com/hris-api/internal/shared/model"
	"github.com/google/uuid"
)

// Employee
type Employee struct {
	model.TenantBaseModel
	UserID     uuid.UUID  `gorm:"type:uuid;uniqueIndex;not null" json:"user_id"`
	NIK        string     `gorm:"type:varchar(50);not null;uniqueIndex:idx_company_nik" json:"nik"`
	FirstName  string     `gorm:"type:varchar(100);not null" json:"first_name"`
	LastName   string     `gorm:"type:varchar(100)" json:"last_name"`
	Email      string     `gorm:"type:varchar(255);not null;uniqueIndex:idx_company_emp_email" json:"email"`
	Phone      string     `gorm:"type:varchar(50)" json:"phone"`
	JoinDate   time.Time  `gorm:"type:date;not null" json:"join_date"`
	ResignDate *time.Time `gorm:"type:date" json:"resign_date"`
	JobTitle   string     `gorm:"type:varchar(100);not null" json:"job_title"`
	Department string     `gorm:"type:varchar(100);not null" json:"department"`
}