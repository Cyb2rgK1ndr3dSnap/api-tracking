package security

import (
	"net/http"
	"os"
	"time"

	"github.com/Cyb2rgK1ndr3dSnap/api-tracking/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	UserID      int    `json:"userID"`
	RoleID      int    `json:"roleID"`
	TokenDevice string `json:"tokenDevice"`
	jwt.RegisteredClaims
}

func CreateJWT(user models.User, TokenDevice string, expirationSeconds int) (string, error) {
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
		"tokenDevice": TokenDevice,
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
		var err error
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			tokenString, err = c.Cookie("jwt_token")

			if err != nil {
				c.Redirect(http.StatusFound, "/")
				c.Abort()
				return
			}
		}

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
		c.Set("tokenDevice", claims.TokenDevice)

		c.Next()
	}
}
