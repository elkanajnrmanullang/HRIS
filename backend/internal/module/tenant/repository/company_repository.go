package repository

import (
	"context"
	"errors"

	"celestara.com/hris-api/internal/module/tenant/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// CompanyRepository
type CompanyRepository interface {
	Create(ctx context.Context, company *entity.Company) error
	GetByID(ctx context.Context, id uuid.UUID) (*entity.Company, error)
	Update(ctx context.Context, company *entity.Company) error
}

// companyRepository
type companyRepository struct {
	db *gorm.DB
}

// NewCompanyRepository
func NewCompanyRepository(db *gorm.DB) CompanyRepository {
	return &companyRepository{
		db: db,
	}
}

// Create
func (r *companyRepository) Create(ctx context.Context, company *entity.Company) error {
	return r.db.WithContext(ctx).Create(company).Error
}

// GetByID
func (r *companyRepository) GetByID(ctx context.Context, id uuid.UUID) (*entity.Company, error) {
	var company entity.Company
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&company).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &company, nil
}

// Update
func (r *companyRepository) Update(ctx context.Context, company *entity.Company) error {
	return r.db.WithContext(ctx).Save(company).Error
}