package controllers

import (
	"fmt"

	"github.com/Cyb2rgK1ndr3dSnap/api-tracking/models"
	"github.com/Cyb2rgK1ndr3dSnap/api-tracking/security"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	//db := c.MustGet("db").(*sql.DB)

	var Body models.User

	err := c.ShouldBindJSON(&Body)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := security.HashPassword(Body.Password)
	if err != nil {
		fmt.Println("Error hashing password:", err)
		return
	}
	fmt.Println("Hashed password:", hashedPassword)
	c.JSON(200, gin.H{"message": "User created successfully"})
	return
	/*_, err = db.Exec("INSERT INTO users (id_user, name, lastname, email, direction, phone_number, cc, password) VALUES ($1, $2, $3, $4, $5, $6, $7,$8)",
		Body.IDUser, Body.Name, Body.Lastname, Body.Email, Body.Direction, Body.PhoneNumber, Body.CC, hashedPassword)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}*/
}
