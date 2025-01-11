package controllers

import (
	"os"

	"github.com/gin-gonic/gin"
)

func GetBusinessInformation(c *gin.Context) {

	supportNumber := os.Getenv("SUPPORT_NUMBER")
	accountBankNumber := os.Getenv("ACCOUNT_BANKNUMBER")

	c.JSON(200, gin.H{
		"supportNumber":     supportNumber,
		"accountBankNumber": accountBankNumber,
	})
}
