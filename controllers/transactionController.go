package controllers

import (
	"database/sql"

	"github.com/Cyb2rgK1ndr3dSnap/api-tracking/models"
	"github.com/Cyb2rgK1ndr3dSnap/api-tracking/services"
	"github.com/gin-gonic/gin"
)

func CreateTransaction(c *gin.Context) {
	db := c.MustGet("db").(*sql.DB)

	var Body models.CreateTransaction

	err := c.ShouldBindJSON(&Body)
	if err != nil {
		c.JSON(400, gin.H{"message": "Please complete all the required data"})
		return
	}

	tx, err := db.Begin()
	if err != nil {
		c.JSON(400, gin.H{"message": "Error with server"})
		return
	}

	err = services.CreateTransaction(Body, tx)
	if err != nil {
		tx.Rollback()
		c.JSON(400, gin.H{"message": "Error with save data"})
		return
	}

	c.JSON(200, gin.H{"message": "Transaction create successfully"})
}

func BalanceTransaction(c *gin.Context) {
	db := c.MustGet("db").(*sql.DB)

	var Body models.QuantityShipping

	err := c.ShouldBindQuery(&Body)
	if err != nil {
		c.JSON(400, gin.H{"message": "Please fill all the required data"})
		return
	}

	Body.IDUser = c.MustGet("userID").(int)

	balance, err := services.BalanceTransaction(Body.IDUser, db)
	if err != nil {
		c.JSON(400, gin.H{"error": "Error with server"})
		return
	}

	quantity, err := services.QuantityShipping(Body, db)
	if err != nil {
		c.JSON(400, gin.H{"error": "Error with server"})
		return
	}

	c.JSON(200, gin.H{"balance": balance, "quantity": quantity})
}
