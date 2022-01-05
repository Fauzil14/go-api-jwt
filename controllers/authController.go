package controllers

import (
	"fmt"
	"go-api-jwt/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// inputan login
type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// inputan register
type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

// Login User godoc
// @Summary Login as an User
// @Description Logging in to get jwt token for authorization
// @Tags Auth
// @Param Body body LoginInput true "the body to login"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /login [post]
func Login(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	usr := models.User{}

	usr.Username = input.Username
	usr.Password = input.Password

	token, err := models.LoginCheck(usr.Username, usr.Password, db)

	if err != nil {
		fmt.Println(err)
		// gin.H == gin.(map[string]interface{})
		c.JSON(http.StatusBadRequest, gin.H{"error": "incorrect username or password"})
		return
	}

	// data yang akan di kirim ke front end
	user := map[string]string{
		"username": usr.Username,
		"email":    usr.Email,
	}

	c.JSON(http.StatusOK, gin.H{"message": "login success", "user": user, "token": token})
}

// Register User godoc
// @Summary Register new User
// @Description Register user from public access
// @Tags Auth
// @Param Body body RegisterInput true "the body to Register"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /register [post]
func Register(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var input RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	usr := models.User{}

	usr.Username = input.Username
	usr.Password = input.Password
	usr.Email = input.Email

	// save user with method in user instance
	_, err := usr.SaveUser(db)

	if err != nil {
		fmt.Println(err)
		// gin.H == gin.(map[string]interface{})
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := map[string]string{
		"username": usr.Username,
		"email":    usr.Email,
	}

	c.JSON(http.StatusOK, gin.H{"message": "registeration success", "user": user})
}
