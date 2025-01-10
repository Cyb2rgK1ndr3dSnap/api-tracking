package controllers

import (
	"database/sql"
	"os"
	"strconv"

	"github.com/Cyb2rgK1ndr3dSnap/api-tracking/models"
	"github.com/Cyb2rgK1ndr3dSnap/api-tracking/security"
	"github.com/Cyb2rgK1ndr3dSnap/api-tracking/services"
	"github.com/gin-gonic/gin"
)

// @Summary Registro de usuario
// @Description Realiza el guardado del usuario en la BD
// @Tags User
// @Accept json
// @Produce application/json
// @Param User body models.RegisterUser true "crea usuario"
// @Success 200 {object} models.SuccessMessage "mensaje de éxito"
// @Failure 400 {object} models.ErrorMessage "Error en los datos proporcionados"
// @Router /user [post]
func RegisterUser(c *gin.Context) {
	db := c.MustGet("db").(*sql.DB)

	var registerUser models.RegisterUser

	err := c.ShouldBindJSON(&registerUser)
	if err != nil {
		c.JSON(400, gin.H{"error": "Please fill all the required data"})
		return
	}

	u, err := services.GetUserByUsername(registerUser.UserName, db)
	if err != nil && err != sql.ErrNoRows {
		c.JSON(400, gin.H{"error": "Error with server"})
		return
	}

	if u != nil {
		c.JSON(400, gin.H{"error": "username already exists"})
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

	registerUser.Role = 2

	err = services.RegisterUser(registerUser, hashedPassword, db)
	if err != nil {
		c.JSON(400, gin.H{"error": "Error with the register of user"})
		return
	}

	c.JSON(200, gin.H{"message": "Registered user"})
}

// @Summary Inicio de sesión de usuario
// @Description Realiza el guardado del usuario en la BD
// @Tags User
// @Accept json
// @Produce application/json
// @Param User body models.LoginUser true "Email y contraseña"
// @Success 200 {object} models.User "mensaje de éxito"
// @Failure 400 {object} models.ErrorMessage "Error en los datos proporcionados"
// @Router /user/login [post]
func LoginUser(c *gin.Context) {
	db := c.MustGet("db").(*sql.DB)

	var u *models.User
	var loginUser models.LoginUser

	err := c.ShouldBindJSON(&loginUser)
	if err != nil {
		c.JSON(400, gin.H{"error": "Please fill all the required data"})
		return
	}

	u, err = services.GetUserByUsername(loginUser.Email, db)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid user or password"})
		return
	}

	if !security.CheckPasswordHash(loginUser.Password, u.Password) {
		c.JSON(400, gin.H{"error": "invalid user or password"})
		return
	}

	expirationSeconds, err := strconv.Atoi(os.Getenv("JWT_EXPIRATION_IN_SECONDS"))
	if err != nil {
		expirationSeconds = 3600 // 1 hora por defecto
	}

	token, err := security.CreateJWT(*u, loginUser.Token, expirationSeconds)
	if err != nil {
		c.JSON(400, gin.H{"error": "Interal error"})
		return
	}

	c.SetCookie("jwt_token", token, expirationSeconds, "/", "", false, true)

	if loginUser.Token != "" {
		RegisterToken := models.RegisterToken{
			IDUser: u.IDUser,
			Token:  loginUser.Token,
		}

		services.RegisterDevice(RegisterToken, db)
	}

	u.Password = ""
	u.Token = token

	c.JSON(200, u)
}

func LogoutUser(c *gin.Context) {
	db := c.MustGet("db").(*sql.DB)

	token, _ := c.Get("tokenDevice")

	services.DeleteDevice(token, db)
}

func GetUsersTotal(c *gin.Context) {
	db := c.MustGet("db").(*sql.DB)

	var Body models.ReadUser

	err := c.ShouldBindJSON(&Body)
	if err != nil {
		c.JSON(400, gin.H{"error": "Please fill all the required data"})
		return
	}

	Body.IDRole = c.MustGet("roleID").(int)

	total, err := services.GetUserTotal(db, Body)
	if err != nil {
		c.JSON(400, gin.H{"error": "Error with server" + err.Error()})
		return
	}

	c.JSON(200, gin.H{"total": total})
}
