package services

import (
	"database/sql"

	"github.com/Cyb2rgK1ndr3dSnap/api-tracking/models"
)

func CreateShipping(createS models.CreateShipping, db *sql.DB) error {
	_, err := db.Exec("INSERT INTO shippings (id_user, shipping_number, weight, amount, quantity, status, expiration_date) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		createS.IDUser, createS.ShippingNumber, createS.Weight, createS.Amount, createS.Quantity, createS.Status, createS.ExpirationDate)
	if err != nil {
		return err
	}
	return nil
}

func UpdateShipping(updateS models.UpdateShipping, db *sql.DB) error {
	_, err := db.Exec("UPDATE shippings SET shipping_number = $1, weight = $2, amount = $3, quantity = $4, status = $5, expiration_date = $6, id_user = $7 WHERE shipping_number = $1",
		updateS.ShippingNumber, updateS.Weight, updateS.Amount, updateS.Quantity, updateS.Status, updateS.ExpirationDate, updateS.IDUser)
	if err != nil {
		return err
	}
	return nil
}

func ReadShipping(selectS models.ReadShipping, db *sql.DB) (*sql.Rows, error) {

	if selectS.Email != "" {
		rows, err := db.Query("SELECT s.*,u.email FROM Shippings s JOIN Users u ON s.id_user = u.id_user WHERE u.email = $1", selectS.Email)
		if err != nil {
			return nil, err
		}
		return rows, nil
	} else {
		rows, err := db.Query("SELECT s.*,u.email FROM shippings s JOIN Users u ON s.id_user = u.id_user WHERE s.shipping_number = $1", selectS.ShippingNumber)
		if err != nil {
			return nil, err
		}
		return rows, nil
	}
}
