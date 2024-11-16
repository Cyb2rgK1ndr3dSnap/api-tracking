package controllers

import (
	"database/sql"

	"github.com/Cyb2rgK1ndr3dSnap/api-tracking/models"
	"github.com/Cyb2rgK1ndr3dSnap/api-tracking/security"
	"github.com/Cyb2rgK1ndr3dSnap/api-tracking/services"
	"github.com/gin-gonic/gin"
)

// Cambiar mensaje en cada status 400
func RegisterUser(c *gin.Context) {
	db := c.MustGet("db").(*sql.DB)

	var registerUser models.RegisterUser

	err := c.ShouldBindJSON(&registerUser)
	if err != nil {
		c.JSON(400, gin.H{"error": "Please fill all the required data"})
		return
	}

	_, err = services.GetUserByEmail(registerUser.Email, db)
	if err == nil {
		c.JSON(400, gin.H{"error": "user with email already exists"})
		return
	}

	if registerUser.Password != registerUser.ConfirmPassword {
		c.JSON(400, gin.H{"error": "Passwords do not match"})
		return
	}

	hashedPassword, err := security.HashPassword(registerUser.Password)
	if err != nil {
		c.JSON(400, gin.H{"error": "Error with server"})
		return
	}

	err = services.CreateUser(registerUser, hashedPassword, db)

	if err != nil {
		c.JSON(400, gin.H{"error": "Error with the creation of user"})
		return
	}

	c.JSON(200, gin.H{"message": "User created successfully"})
}

func LoginUser(c *gin.Context) {
	db := c.MustGet("db").(*sql.DB)

	var loginUser models.LoginUser

	err := c.ShouldBindJSON(&loginUser)
	if err != nil {
		c.JSON(400, gin.H{"error": "Please fill all the required data"})
		return
	}

	u, err := services.GetUserByEmail(loginUser.Email, db)
	if err != nil {
		c.JSON(400, gin.H{"error": "user with that email not exists"})
		return
	}

	if !security.CheckPasswordHash(loginUser.Password, u.Password) {
		c.JSON(400, gin.H{"error": "invalid email or password"})
		return
	}

	token, err := security.CreateJWT(*u)
	if err != nil {
		c.JSON(400, gin.H{"error": "Interal error"})
		return
	}

	u.Password = ""
	u.Token = token

	c.JSON(200, u)
}
