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
	Body.Status = 2

	tx, err := db.Begin()
	if err != nil {
		c.JSON(400, gin.H{"message": "Error with server"})
		return
	}

	idShipping, err := services.CreateShipping(Body, tx)
	if err != nil {
		tx.Rollback()
		c.JSON(400, gin.H{"message": "Error with save data"})
		return
	}

	transaction := models.CreateTransaction{
		IDUser:            Body.IDUser,
		IDShipping:        idShipping,
		IDTransactionType: 1,
		Amount:            Body.Amount,
	}

	err = services.CreateTransaction(transaction, tx)
	if err != nil {
		tx.Rollback()
		c.JSON(400, gin.H{"message": "Error with save data"})
		return
	}

	err = tx.Commit()
	if err != nil {
		c.JSON(400, gin.H{"message": "Error with save data"})
		return
	}

	c.JSON(200, gin.H{"message": "Shipping created successfully"})
}

func SearchShipping(c *gin.Context) {
	db := c.MustGet("db").(*sql.DB)

	var Body models.ReadShipping

	err := c.ShouldBindQuery(&Body)
	if err != nil {
		c.JSON(400, gin.H{"message": "Please fill all the required data"})
		return
	}

	rows, err := services.SearchShipping(Body, db)
	if err != nil {
		c.JSON(400, gin.H{"error": "Not exist data matching"})
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
			&shipping.Created_date,
			&shipping.LastUpdate,
			&shipping.ExpirationDate,
			&shipping.Email,
		)
		if err != nil {
			c.JSON(400, gin.H{"error": "Error with database"})
			return
		}
		shippings = append(shippings, shipping)
	}

	/*if err = rows.Err(); err != nil {
		c.JSON(400, gin.H{"error": "Error with database"})
		return
	}*/

	c.JSON(200, shippings)
}

func UpdateShipping(c *gin.Context) {
	db := c.MustGet("db").(*sql.DB)

	var Body models.UpdateShipping

	err := c.ShouldBindJSON(&Body)
	if err != nil {
		c.JSON(400, gin.H{"error": "Please complete all the required data"})
		return
	}

	currentStatus, err := services.StatusShipping(Body.IDShipping, db)
	if err != nil {
		c.JSON(400, gin.H{"error": "Shipping not found"})
		return
	}

	if currentStatus == 1 {
		c.JSON(400, gin.H{"error": "Cannot modify a shipping that is already closed"})
		return
	}

	u, err := services.GetUserByEmail(Body.Email, db)
	if err != nil {
		c.JSON(400, gin.H{"error": "user with that email not exists"})
		return
	}

	Body.IDUser = u.IDUser

	tx, err := db.Begin()
	if err != nil {
		c.JSON(400, gin.H{"message": "Error with server"})
		return
	}

	err = services.UpdateShipping(Body, tx)
	if err != nil {
		c.JSON(400, gin.H{"error": "Error with update Shipping data"})
		return
	}

	transaction := models.UpdateTransaction{
		IDUser:            Body.IDUser,
		IDShipping:        Body.IDShipping,
		IDTransactionType: Body.IDTransactionType,
		Amount:            Body.Amount,
	}

	err = services.UpdateTransaction(transaction, tx)
	if err != nil {
		c.JSON(400, gin.H{"error": "Error with update Shipping data"})
		return
	}

	err = tx.Commit()
	if err != nil {
		c.JSON(400, gin.H{"message": "Error with save data"})
		return
	}

	c.JSON(200, gin.H{"message": "Shipping updated successfully"})
}

func CloseShipping(c *gin.Context) {
	db := c.MustGet("db").(*sql.DB)

	var Body models.ReadShipping

	err := c.ShouldBindQuery(&Body)
	if err != nil {
		c.JSON(400, gin.H{"message": "Please fill all the required data"})
		return
	}

	rows, err := services.SearchShipping(Body, db)
	if err != nil {
		c.JSON(400, gin.H{"error": "Not exist data matching"})
		return
	}

	defer rows.Close()

	var shipping models.Shipping

	if rows.Next() {
		err := rows.Scan(
			&shipping.IDShipping,
			&shipping.IDUser,
			&shipping.ShippingNumber,
			&shipping.Weight,
			&shipping.Amount,
			&shipping.Quantity,
			&shipping.Status,
			&shipping.Created_date,
			&shipping.LastUpdate,
			&shipping.ExpirationDate,
			&shipping.Email,
		)
		if err != nil {
			c.JSON(400, gin.H{"error": "Error with database"})
			return
		}
	} else {
		c.JSON(404, gin.H{"error": "No shipping data found"})
		return
	}

	if shipping.Status == 1 {
		c.JSON(400, gin.H{"error": "Cannot create transaction, that shipping is already closed"})
		return
	}

	transaction := models.CreateTransaction{
		IDUser:            shipping.IDUser,
		IDShipping:        shipping.IDShipping,
		IDTransactionType: 2,
		Amount:            (shipping.Amount * -1),
	}

	tx, err := db.Begin()
	if err != nil {
		c.JSON(400, gin.H{"message": "Error with server"})
		return
	}

	err = services.CreateTransaction(transaction, tx)
	if err != nil {
		tx.Rollback()
		c.JSON(400, gin.H{"message": "Error with save data"})
		return
	}
	err = tx.Commit()
	if err != nil {
		c.JSON(400, gin.H{"message": "Error with save data"})
		return
	}

	c.JSON(200, gin.H{"message": "Shipping Closed successfully"})
}
