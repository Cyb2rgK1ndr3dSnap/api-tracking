package controllers

import (
	"database/sql"

	"github.com/Cyb2rgK1ndr3dSnap/api-tracking/models"
	"github.com/Cyb2rgK1ndr3dSnap/api-tracking/services"
	"github.com/Cyb2rgK1ndr3dSnap/api-tracking/utils"
	"github.com/gin-gonic/gin"
)

// @Summary Creación de paquete
// @Description Realiza el guardado del paquete que se quiere crear en la BD y una transacción se abre
// @Tags Shipping
// @Security JWT
// @Accept json
// @Produce application/json
// @Param shipping body models.CreateShipping true "crea paquete y transacción"
// @Success 200 {object} models.SuccessMessage "mensaje de éxito"
// @Failure 400 {object} models.ErrorMessage "Error en los datos proporcionados"
// @Router /shipping [post]
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

// @Summary Busqueda de paquetes
// @Description Realiza la busqueda de pedidos del usuario por defecto o email, id de paquete o número de paquete
// @Tags Shipping
// @Security JWT
// @Param email query string false "Correo electrónico del usuario a buscar"
// @Param shipping_number query int false "Código de paquete a buscar"
// @Param id_shipping query int false "Id de paquete a buscar"
// @Produce application/json
// @Success 200 {array} models.Shipping "Lista de paquetes encontrados"
// @Failure 400 {object} models.ErrorMessage "Error en los datos proporcionados"
// @Router /shipping [get]
func GetShipping(c *gin.Context) {
	db := c.MustGet("db").(*sql.DB)

	var Body models.ReadShipping

	err := c.ShouldBindQuery(&Body)
	if err != nil {
		c.JSON(400, gin.H{"message": "Please fill all the required data"})
		return
	}

	Body.IDUser = c.MustGet("userID").(int)
	Body.IDRole = c.MustGet("roleID").(int)

	rows, err := services.GetShipping(Body, db)
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

// @Summary Actualización de datos de paquete
// @Tags Shipping
// @Security JWT
// @Accept json
// @Produce application/json
// @Param shipping body models.UpdateShipping true "actualiza los datos del paquete"
// @Success 200 {object} models.SuccessMessage "mensaje de éxito"
// @Failure 400 {object} models.ErrorMessage "Error en el proceso"
// @Router /shipping [put]
func UpdateShipping(c *gin.Context) {
	db := c.MustGet("db").(*sql.DB)

	var Body models.UpdateShipping

	err := c.ShouldBindJSON(&Body)
	if err != nil {
		c.JSON(400, gin.H{"error": "Please complete all the required data"})
		return
	}

	currentStatus, err := services.StatusShippingByID(Body.IDShipping, db)
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
		tx.Rollback()
		c.JSON(400, gin.H{"error": "Error with update Shipping data"})
		return
	}

	transaction := models.UpdateTransaction{
		IDUser:     Body.IDUser,
		IDShipping: Body.IDShipping,
		Amount:     Body.Amount,
	}

	err = services.UpdateTransaction(transaction, tx)
	if err != nil {
		tx.Rollback()
		c.JSON(400, gin.H{"error": "Error with update Shipping data"})
		return
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		c.JSON(400, gin.H{"error": "Error with save data"})
		return
	}

	c.JSON(200, gin.H{"message": "Shipping updated successfully"})
}

// @Summary Cerrar paquete y creación de transacción
// @Tags Shipping
// @Security JWT
// @Accept json
// @Produce application/json
// @Param shipping body models.CloseShipping true "toma el id del paquete para cerrarlo y crear transacción de cierre"
// @Success 200 {object} models.SuccessMessage "mensaje de éxito"
// @Failure 400 {object} models.ErrorMessage "Error en el proceso"
// @Router /shipping/close [post]
func CloseShipping(c *gin.Context) {
	db := c.MustGet("db").(*sql.DB)

	var Body models.CloseShipping

	err := c.ShouldBindQuery(&Body)
	if err != nil {
		c.JSON(400, gin.H{"error": "Please fill all the required data"})
		return
	}

	shipping, err := services.GetShippingByID(Body.IDShipping, db)
	if err != nil {
		c.JSON(400, gin.H{"error": "Not exist data matching"})
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

	shippingU := models.UpdateShipping{
		IDShipping: shipping.IDShipping,
		Status:     1,
	}

	tx, err := db.Begin()
	if err != nil {
		c.JSON(400, gin.H{"error": "Error with server"})
		return
	}

	err = services.UpdateShipping(shippingU, tx)
	if err != nil {
		tx.Rollback()
		c.JSON(400, gin.H{"error": "Error with update Shipping data"})
		return
	}

	err = services.CreateTransaction(transaction, tx)
	if err != nil {
		tx.Rollback()
		c.JSON(400, gin.H{"error": "Error with save data"})
		return
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		c.JSON(400, gin.H{"error": "Error with save data"})
		return
	}

	rows, err := services.ReadDevices(shipping.IDUser, db)
	defer rows.Close()
	var devices = make([]string, 0, 10)
	for rows.Next() {
		var token string
		err := rows.Scan(
			&token,
		)
		if err != nil {
			c.JSON(400, gin.H{"error": "Error with database"})
			return
		}
		devices = append(devices, token)
	}

	go func(devices []string) {
		utils.SendPushNotification(devices, "Shipping closed", "Your shipping has been closed")
	}(devices)

	c.JSON(200, gin.H{"message": "Shipping Closed successfully"})
}
