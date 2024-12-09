package controllers

import (
	"database/sql"

	"github.com/Cyb2rgK1ndr3dSnap/api-tracking/models"
	"github.com/Cyb2rgK1ndr3dSnap/api-tracking/services"
	"github.com/gin-gonic/gin"
)

// @Summary Creación de transacción
// @Description Realiza el guardado de la transacción que se quiere crear en la BD
// @Tags Transaction
// @Security JWT
// @Accept json
// @Produce application/json
// @Param shipping body models.CreateTransaction true "crea transacción"
// @Success 200 {object} models.SuccessMessage "mensaje de éxito"
// @Failure 400 {object} models.ErrorMessage "Error en los datos proporcionados"
// @Router /shipping [post]
func CreateTransaction(c *gin.Context) {
	db := c.MustGet("db").(*sql.DB)

	var Body models.CreateTransaction

	err := c.ShouldBindJSON(&Body)
	if err != nil {
		c.JSON(400, gin.H{"error": "Please complete all the required data"})
		return
	}

	tx, err := db.Begin()
	if err != nil {
		c.JSON(400, gin.H{"error": "Error with server"})
		return
	}

	err = services.CreateTransaction(Body, tx)
	if err != nil {
		tx.Rollback()
		c.JSON(400, gin.H{"error": "Error with save data"})
		return
	}

	c.JSON(200, gin.H{"message": "Transaction create successfully"})
}

// @Summary Estado de cuenta de usuario
// @Tags Transaction
// @Security JWT
// @Produce application/json
// @Success 200 {object} models.BalanceTransaction "ESTADO DE SALDO Y PAQUETES QUE TIENE"
// @Failure 400 {object} models.ErrorMessage "Error en los datos proporcionados"
// @Router /transaction [get]
func BalanceTransaction(c *gin.Context) {
	db := c.MustGet("db").(*sql.DB)

	var Body models.QuantityShipping
	var balance models.BalanceTransaction
	var err error

	Body.IDUser = c.MustGet("userID").(int)
	Body.Status = 2

	balance.Balance, err = services.BalanceTransaction(Body.IDUser, db)
	if err != nil {
		c.JSON(400, gin.H{"error": "Error with server"})
		return
	}

	balance.Quantity, err = services.QuantityShipping(Body, db)
	if err != nil {
		c.JSON(400, gin.H{"error": "Error with server"})
		return
	}

	c.JSON(200, balance)
}
