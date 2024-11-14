package controllers

import (
	"database/sql"

	"github.com/Cyb2rgK1ndr3dSnap/api-tracking/models"
	"github.com/Cyb2rgK1ndr3dSnap/api-tracking/services"
	"github.com/gin-gonic/gin"
)

func CreateShipping(c *gin.Context) {
	db := c.MustGet("db").(*sql.DB)

	var Body models.CreateShipping

	err := c.ShouldBindJSON(&Body)
	if err != nil {
		c.JSON(400, gin.H{"message": "Please complete all the required data"})
		return
	}

	u, err := services.GetUserByEmail(Body.Email, db)
	if err != nil {
		c.JSON(400, gin.H{"message": "user with that email not exists"})
		return
	}

	Body.IDUser = u.IDUser

	err = services.CreateShipping(Body, db)
	if err != nil {
		c.JSON(400, gin.H{"message": "Error with Shipping data"})
		return
	}

	c.JSON(200, gin.H{"message": "Shipping created successfully"})
}

func UpdateShipping(c *gin.Context) {
	db := c.MustGet("db").(*sql.DB)

	var Body models.CreateShipping

	err := c.ShouldBindJSON(&Body)
	if err != nil {
		c.JSON(400, gin.H{"message": "user with that email not exists"})
		return
	}

	err = services.UpdateShipping(Body, db)
	if err != nil {
		c.JSON(400, gin.H{"message": "Error with update Shipping data"})
		return
	}

	c.JSON(200, gin.H{"message": "Shipping created successfully"})
}

func ReadShipping(c *gin.Context) {
	db := c.MustGet("db").(*sql.DB)

	var Body models.ReadShipping

	err := c.ShouldBindQuery(&Body)
	if err != nil {
		c.JSON(400, gin.H{"message": "user with that email not exists"})
		return
	}

	rows, err := services.ReadShipping(Body, db)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	defer rows.Close()

	var shippings []models.Shipping
	for rows.Next() {
		var shipping models.Shipping
		err := rows.Scan(
			&shipping.IDShipping,
			&shipping.IDUser,
			&shipping.ShippingNumber,
			&shipping.Weight,
			&shipping.Amount,
			&shipping.Quantity,
			&shipping.Status,
			&shipping.CreationDate,
			&shipping.LastUpdate,
			&shipping.ExpirationDate,
		)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		shippings = append(shippings, shipping)
	}

	if err = rows.Err(); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, shippings)
}
