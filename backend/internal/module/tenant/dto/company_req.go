package dto

// CreateCompanyRequest
type CreateCompanyRequest struct {
	Name         string `json:"name" binding:"required,max=255"`
	Email        string `json:"email" binding:"required,email"`
	Phone        string `json:"phone" binding:"omitempty,max=50"`
	Address      string `json:"address" binding:"omitempty"`
	MaxEmployees int    `json:"max_employees" binding:"required,min=1"`
}