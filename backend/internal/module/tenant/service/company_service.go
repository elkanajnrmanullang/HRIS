package service

import (
	"context"

	"celestara.com/hris-api/internal/module/tenant/dto"
	"celestara.com/hris-api/internal/module/tenant/entity"
	"celestara.com/hris-api/internal/module/tenant/repository"
)

// CompanyService
type CompanyService interface {
	CreateCompany(ctx context.Context, req dto.CreateCompanyRequest) (*entity.Company, error)
}

// companyService
type companyService struct {
	repo repository.CompanyRepository
}

// NewCompanyService
func NewCompanyService(repo repository.CompanyRepository) CompanyService {
	return &companyService{
		repo: repo,
	}
}

// CreateCompany
func (s *companyService) CreateCompany(ctx context.Context, req dto.CreateCompanyRequest) (*entity.Company, error) {
	company := &entity.Company{
		Name:         req.Name,
		Email:        req.Email,
		Phone:        req.Phone,
		Address:      req.Address,
		MaxEmployees: req.MaxEmployees,
		IsActive:     true, 
	}

	err := s.repo.Create(ctx, company)
	if err != nil {
		return nil, err
	}

	return company, nil
}