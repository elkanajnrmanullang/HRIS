package dto

import "github.com/google/uuid"

// RegisterRequest
type RegisterRequest struct {
	CompanyID uuid.UUID `json:"company_id" binding:"required"`
	Email     string    `json:"email" binding:"required,email"`
	Password  string    `json:"password" binding:"required,min=6"`
}

// LoginRequest
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}