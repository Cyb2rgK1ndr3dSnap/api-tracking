package security

import (
	"os"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func GetBcryptCost() int {
	costStr := os.Getenv("BCRYPT_COST")
	if costStr == "" {
		return bcrypt.DefaultCost // Default cost if not set
	}
	cost, err := strconv.Atoi(costStr)
	if err != nil || cost < bcrypt.MinCost || cost > bcrypt.MaxCost {
		return bcrypt.DefaultCost // Default cost if invalid
	}
	return cost
}

func HashPassword(password string) (string, error) {
	cost := GetBcryptCost()
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
