package security

import (
	"os"
	"strconv"
	"time"

	"github.com/Cyb2rgK1ndr3dSnap/api-tracking/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	UserID int `json:"userID"`
	RoleID int `json:"roleID"`
	jwt.RegisteredClaims
}

func CreateJWT(user models.User) (string, error) {
	expirationSeconds, err := strconv.Atoi(os.Getenv("JWT_EXPIRATION_IN_SECONDS"))
	if err != nil {
		expirationSeconds = 3600 // 1 hora por defecto
	}

	expiration := time.Second * time.Duration(expirationSeconds)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":      user.IDUser,
		"name":        user.FirstName,
		"lastname":    user.LastName,
		"email":       user.Email,
		"direction":   user.Direction,
		"phonenumber": user.PhoneNumber,
		"cedula":      user.CC,
		"roleID":      user.IDRole,
		"exp":         time.Now().Add(expiration).Unix(),
	})

	secret := []byte(os.Getenv("JWTSECRET"))

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, err
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		claims := &CustomClaims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWTSECRET")), nil
		})

		if err != nil || !token.Valid {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		c.Set("userID", claims.UserID)
		c.Set("roleID", claims.RoleID)

		c.Next()
	}
}
