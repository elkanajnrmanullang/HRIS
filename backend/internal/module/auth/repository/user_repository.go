package repository

import (
	"context"
	"errors"

	"celestara.com/hris-api/internal/module/auth/entity"
	"gorm.io/gorm"
)

// UserRepository
type UserRepository interface {
	Create(ctx context.Context, user *entity.User) error
	FindByEmail(ctx context.Context, email string) (*entity.User, error)
}

// userRepository
type userRepository struct {
	db *gorm.DB
}

// NewUserRepository
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

// Create
func (r *userRepository) Create(ctx context.Context, user *entity.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}

// FindByEmail
func (r *userRepository) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	var user entity.User
	err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil 
		}
		return nil, err
	}
	return &user, nil
}