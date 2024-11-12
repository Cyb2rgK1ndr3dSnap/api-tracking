package security

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/Cyb2rgK1ndr3dSnap/api-tracking/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func CreateJWT(user models.User) (string, error) {
	expirationSeconds, err := strconv.Atoi(os.Getenv("JWT_EXPIRATION_IN_SECONDS"))
	if err != nil {
		expirationSeconds = 3600 // 1 hora por defecto
	}

	expiration := time.Second * time.Duration(expirationSeconds)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":      strconv.Itoa(int(user.IDUser)),
		"name":        user.FirstName,
		"lastname":    user.LastName,
		"email":       user.Email,
		"direction":   user.Direction,
		"phonenumber": user.PhoneNumber,
		"cedula":      user.CC,
		"roleID":      user.IDRole,
		"exp":         time.Now().Add(expiration).Unix(),
	})

	secret := []byte(os.Getenv("JWTSecret"))

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, err
}

func validateJWT(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JWTSecret")), nil
	})
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		token, err := validateJWT(tokenString)

		if err != nil || !token.Valid {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		//c.JSON(200, gin.H{"mess": claims})
		//str := claims["userID"].(string)

		c.Set("user_id", claims["userID"])
		//c.Set("username", claims.Username)
		//c.Set("role", claims.Role)

		c.Next()
	}
}
