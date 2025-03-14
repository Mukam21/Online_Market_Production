package handlers

import (
	"net/http"

	"github.com/Mukam21/Online_Market_Production/internal/jwt"
	"github.com/Mukam21/Online_Market_Production/internal/models"
	"github.com/Mukam21/Online_Market_Production/internal/services"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	Service services.UserService
}

func NewUserHandler(service services.UserService) *UserHandler {
	return &UserHandler{Service: service}
}

// @Summary Create a new user
// @Description Register a new user
// @Accept  json
// @Produce  json
// @Param user body models.User true "User Information"
// @Success 201 {object} gin.H{"message": "User created"}
// @Failure 400 {object} gin.H{"error": "Bad request"}
// @Router /api/auth/register [post]
func (h *UserHandler) Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if err := h.Service.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "User registered"})
}

// @Summary User login
// @Description User login with email and password
// @Accept  json
// @Produce  json
// @Param credentials body models.LoginCredentials true "User Credentials"
// @Success 200 {object} gin.H{"token": "JWT token"}
// @Failure 401 {object} gin.H{"error": "Invalid email or password"}
// @Router /api/auth/login [post]
func (h *UserHandler) Login(c *gin.Context) {
	var credentials models.User
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	user, err := h.Service.FindByEmail(credentials.Email)
	if err != nil || user.Password != credentials.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Генерация токена с email и userID
	token, err := jwt.GenerateToken(user.Email, user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
