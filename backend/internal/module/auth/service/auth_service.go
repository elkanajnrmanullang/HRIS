package service

import (
	"context"
	"errors"

	"celestara.com/hris-api/internal/module/auth/dto"
	"celestara.com/hris-api/internal/module/auth/entity"
	"celestara.com/hris-api/internal/module/auth/repository"
	"celestara.com/hris-api/internal/shared/utils"
)

type AuthService interface {
	Register(ctx context.Context, req dto.RegisterRequest) (*entity.User, error)
	Login(ctx context.Context, req dto.LoginRequest) (string, error)
}

type authService struct {
	repo repository.UserRepository
}

// NewAuthService constructor
func NewAuthService(repo repository.UserRepository) AuthService {
	return &authService{
		repo: repo,
	}
}

func (s *authService) Register(ctx context.Context, req dto.RegisterRequest) (*entity.User, error) {
	existingUser, err := s.repo.FindByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New("email is already registered")
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	user := &entity.User{
		Email:        req.Email,
		PasswordHash: hashedPassword,
		Role:         "tenant_admin",
		IsActive:     true,
	}
	user.CompanyID = req.CompanyID

	err = s.repo.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *authService) Login(ctx context.Context, req dto.LoginRequest) (string, error) {
	user, err := s.repo.FindByEmail(ctx, req.Email)
	if err != nil || user == nil {
		return "", errors.New("invalid credentials")
	}

	// Verifikasi password hash
	if !utils.CheckPasswordHash(req.Password, user.PasswordHash) {
		return "", errors.New("invalid credentials")
	}

	// Buat JWT
	token, err := utils.GenerateToken(user.ID, user.Role)
	if err != nil {
		return "", err
	}

	return token, nil
}