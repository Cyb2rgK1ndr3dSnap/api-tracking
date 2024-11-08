package controllers

import (
	"database/sql"

	"github.com/Cyb2rgK1ndr3dSnap/api-tracking/models"
	"github.com/gin-gonic/gin"
)

func CreateShipping(c *gin.Context) {
	db := c.MustGet("db").(*sql.DB)
	var Body models.Shipping
	err := c.ShouldBindJSON(&Body)

	//currentTime := time.Now()
	//formattedTime := currentTime.Format("2006-01-02 15:04:05")

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	_, err = db.Exec("INSERT INTO shippings (id_user, shipping_number, weight, value, status, expiration_date) VALUES ($1, $2, $3, $4, $5, $6)",
		Body.IDShipping, Body.ShippingNumber, Body.Weight, Body.Value, Body.Status, Body.ExpirationDate)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Shipping created successfully"})
}
