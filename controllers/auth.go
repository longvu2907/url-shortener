package controllers

import (
	"net/http"
	"url-shortener/models"
	"url-shortener/utils/token"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterResponse struct {
	Message string `json:"message"`
}

// @Summary Register
// @Description Register new account
// @Tags auth
// @Accept json
// @Produce json
// @Param data body RegisterRequest true "Register data"
// @Success 200 {object} RegisterResponse
// @Router /auth/register [post]
func Register(c *gin.Context) {
	var input RegisterRequest

	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := models.User{
		Username: input.Username,
		Password: input.Password,
	}
	_, err := u.CreateUser()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, RegisterResponse{Message: "registration success"})
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func loginCheck(username, password string) (string, error) {
	user, err := models.GetUserByUsername(username)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", err
	}

	token, err := token.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

// @Summary Login
// @Description Login your account
// @Tags auth
// @Accept json
// @Produce json
// @Param data body LoginRequest true "Login data"
// @Success 200 {object} LoginResponse "Login response"
// @Router /auth/login [post]
func Login(c *gin.Context) {
	var input LoginRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := models.User{}
	u.Username = input.Username
	u.Password = input.Password
	token, err := loginCheck(u.Username, u.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
		return
	}

	c.JSON(http.StatusOK, LoginResponse{Token: token})
}
