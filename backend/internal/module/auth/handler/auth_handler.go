package handler

import (
	"net/http"

	"celestara.com/hris-api/internal/module/auth/dto"
	"celestara.com/hris-api/internal/module/auth/service"
	"celestara.com/hris-api/internal/shared/response"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service service.AuthService
}

func NewAuthHandler(service service.AuthService) *AuthHandler {
	return &AuthHandler{
		service: service,
	}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.Error("Invalid request payload", err.Error()))
		return
	}

	user, err := h.service.Register(c.Request.Context(), req)
	if err != nil {
		if err.Error() == "email is already registered" {
			c.JSON(http.StatusConflict, response.Error("Registration failed", err.Error()))
			return
		}
		c.JSON(http.StatusInternalServerError, response.Error("Failed to register user", err.Error()))
		return
	}

	c.JSON(http.StatusCreated, response.Success("User registered successfully", user))
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest

	// Validasi input JSON
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.Error("Invalid request payload", err.Error()))
		return
	}

	token, err := h.service.Login(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, response.Error("Login failed", err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.Success("Login successful", gin.H{"token": token}))
}