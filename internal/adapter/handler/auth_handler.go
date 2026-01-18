package handler

import (
	"net/http"

	"GoWebapitest/config"
	"GoWebapitest/internal/core/service"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service *service.AuthService
}

func NewAuthHandler(s *service.AuthService) *AuthHandler {
	return &AuthHandler{service: s}
}

func (h *AuthHandler) Login(c *gin.Context) {

	config.EnableCORS(c.Writer)

	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	token, err := h.service.Login(body.Email, body.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error(), "status": false})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token,"status": true})
}

func (h *AuthHandler) GetAllUsers(c *gin.Context) {

	config.EnableCORS(c.Writer)

	users, err := h.service.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"users": users})
}