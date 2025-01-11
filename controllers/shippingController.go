package controllers

import (
	"context"
	"database/sql"
	"strconv"
	"sync"
	"time"

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
		c.JSON(400, gin.H{"error": "Please complete all the required data"})
		return
	}

	u, err := services.GetUserByUsername(Body.Username, db)
	if err != nil {
		c.JSON(400, gin.H{"error": "user with that username not exists"})
		return
	}

	Body.IDUser = u.IDUser
	Body.Status = 2

	tx, err := db.Begin()
	if err != nil {
		c.JSON(400, gin.H{"error": "Error with server"})
		return
	}

	idShipping, err := services.CreateShipping(Body, tx)
	if err != nil {
		tx.Rollback()
		c.JSON(400, gin.H{"error": "Error with save data"})
		return
	}

	transaction := models.CreateTransaction{
		IDUser:            Body.IDUser,
		IDShipping:        idShipping,
		IDTransactionType: 2,
		Amount:            Body.Amount,
	}

	err = services.CreateTransaction(transaction, tx)
	if err != nil {
		tx.Rollback()
		c.JSON(400, gin.H{"error": "Error with save data"})
		return
	}

	err = tx.Commit()
	if err != nil {
		c.JSON(400, gin.H{"error": "Error with save data"})
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

	var Body models.ReadShippings

	err := c.ShouldBindQuery(&Body)
	if err != nil {
		c.JSON(400, gin.H{"message": "Please fill all the required data" + err.Error()})
		return
	}

	Body.IDUser = c.MustGet("userID").(int)
	Body.IDRole = c.MustGet("roleID").(int)

	var response = make(gin.H)

	if Body.PageSize == 0 {
		Body.PageSize = 10
	}

	if Body.Web && Body.PageNumber == 1 {
		totalPages, err := services.GetShippingMaxPage(Body, db)
		if err != nil {
			c.JSON(400, gin.H{"error": "Not exist data matching"})
			return
		}

		response["total_pages"] = (totalPages + Body.PageSize - 1) / Body.PageSize
	}

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
			&shipping.Username,
			&shipping.Debt,
		)
		if err != nil {
			c.JSON(400, gin.H{"error": "Error with database" + err.Error()})
			return
		}
		shippings = append(shippings, shipping)
	}

	if Body.Web {
		response["data"] = shippings
		c.JSON(200, response)
	} else {
		c.JSON(200, shippings)
	}
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
	var response = make(gin.H)

	err := c.ShouldBindJSON(&Body)
	if err != nil {
		c.JSON(400, gin.H{"error": "Please complete all the required data"})
		return
	}

	if !Body.ExpirationDate.IsZero() {
		if Body.ExpirationDate.Before(time.Now()) {
			c.JSON(400, gin.H{"error": "The date is not correctly"})
			return
		}
	}

	shipping, err := services.GetShippingByID(Body.IDShipping, db)
	if err != nil {
		c.JSON(400, gin.H{"error": "Not exist data matching"})
		return
	}

	if shipping.Status == 1 {
		c.JSON(400, gin.H{"error": "Cannot make that process, that shipping is already closed"})
		return
	}

	if Body.Username != "" {
		u, err := services.GetUserByUsername(Body.Username, db)
		if err != nil {
			c.JSON(400, gin.H{"error": "user with that username not exists"})
			return
		}
		Body.IDUser = u.IDUser
		response["email"] = u.Email
	}

	tx, err := db.Begin()
	if err != nil {
		c.JSON(400, gin.H{"error": "Error with server"})
		return
	}

	err = services.UpdateShipping(Body, tx)
	if err != nil {
		tx.Rollback()
		c.JSON(400, gin.H{"error": "Error with update Shipping data"})
		return
	}

	if Body.Amount != 0 || Body.IDUser != 0 {
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
	}

	if Body.Status == 1 {
		transaction := models.CreateTransaction{
			IDUser:            shipping.IDUser,
			IDShipping:        shipping.IDShipping,
			IDTransactionType: 1,
			Amount:            (shipping.Amount * -1),
		}

		err = services.CreateTransaction(transaction, tx)
		if err != nil {
			tx.Rollback()
			c.JSON(400, gin.H{"error": "Error with save data"})
			return
		}
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		c.JSON(400, gin.H{"error": "Error with save data"})
		return
	}

	if Body.Status == 1 || Body.Status == 3 {
		// Notification
		message := "Your shipping " + shipping.ShippingNumber
		if Body.Status == 1 {
			message += " has been closed"
		} else if Body.Status == 3 {
			message += " is awaiting pickup"
		}

		data := map[string]string{
			"shippingNumber": shipping.ShippingNumber,
			"idUser":         strconv.Itoa(shipping.IDUser),
			"title":          "Shipping closed",
			"body":           message,
		}
		// End notification

		// Send notification to user device's
		devices, err := services.ReadDevices(shipping.IDUser, db)
		if err != nil {
			c.JSON(400, gin.H{"error": "Error with database"})
			return
		}

		// Si es difente de 0 entrar
		if len(devices) > 0 {
			var wg sync.WaitGroup
			var notificationErr error

			ctx := context.Background()

			wg.Add(1)
			go func(devices []string) {
				defer wg.Done()
				notificationErr = utils.SendNotification(devices, ctx, data)
			}(devices)

			wg.Wait()

			if notificationErr != nil {
				c.JSON(400, gin.H{"error": "Error with push notification"})
				return
			}
		}
		// End notification
	}

	response["message"] = "Shipping updated successfully"

	c.JSON(200, response)
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

	err := c.ShouldBindJSON(&Body)
	if err != nil {
		c.JSON(400, gin.H{"error": "Please fill all the required data" + err.Error()})
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
		IDTransactionType: 1,
		Amount:            ((shipping.Amount + shipping.Debt) * -1),
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

	devices, err := services.ReadDevices(shipping.IDUser, db)
	if err != nil {
		c.JSON(400, gin.H{"error": "Error with push notification: " + err.Error()})
		return
	}

	var wg sync.WaitGroup
	var notificationErr error

	ctx := context.Background()

	wg.Add(1)
	go func(devices []string) {
		defer wg.Done()
		notificationErr = utils.SendNotification(devices, ctx, nil)
	}(devices)

	wg.Wait()

	if notificationErr != nil {
		c.JSON(400, gin.H{"error": "Error with push notification" + notificationErr.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Shipping Closed successfully"})
}

func GetShippingsTotal(c *gin.Context) {
	db := c.MustGet("db").(*sql.DB)

	var Body models.ReadShipping

	err := c.ShouldBindJSON(&Body)
	if err != nil {
		c.JSON(400, gin.H{"error": "Please fill all the required data"})
		return
	}

	Body.IDRole = c.MustGet("roleID").(int)

	total, err := services.GetShippingTotal(db, Body)
	if err != nil {
		c.JSON(400, gin.H{"error": "Error with server" + err.Error()})
		return
	}

	c.JSON(200, gin.H{"total": total})
}
