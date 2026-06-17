package handler

import (
	"net/http"

	"celestara.com/hris-api/internal/module/tenant/dto"
	"celestara.com/hris-api/internal/module/tenant/service"
	"celestara.com/hris-api/internal/shared/response"
	"github.com/gin-gonic/gin"
)

// CompanyHandler
type CompanyHandler struct {
	service service.CompanyService
}

// NewCompanyHandler
func NewCompanyHandler(service service.CompanyService) *CompanyHandler {
	return &CompanyHandler{
		service: service,
	}
}

// Create
func (h *CompanyHandler) Create(c *gin.Context) {
	var req dto.CreateCompanyRequest

	// Validasi JSON input
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.Error("Invalid request payload", err.Error()))
		return
	}

	// Panggil Service
	company, err := h.service.CreateCompany(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Error("Failed to create company", err.Error()))
		return
	}

	// Response sukses
	c.JSON(http.StatusCreated, response.Success("Company created successfully", company))
}