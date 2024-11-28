package services

import (
	"database/sql"
	"fmt"

	"github.com/Cyb2rgK1ndr3dSnap/api-tracking/models"
)

func CreateShipping(createS models.CreateShipping, tx *sql.Tx) (int, error) {
	var idShipping int
	err := tx.QueryRow("INSERT INTO shippings (id_user, shipping_number, weight, amount, quantity, status, expiration_date) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id_shipping",
		createS.IDUser, createS.ShippingNumber, createS.Weight, createS.Amount, createS.Quantity, createS.Status, createS.ExpirationDate).Scan(&idShipping)
	if err != nil {
		return 0, err
	}
	return idShipping, nil
}

func UpdateShipping(updateS models.UpdateShipping, tx *sql.Tx) error {
	_, err := tx.Exec("UPDATE shippings SET shipping_number = $1, weight = $2, amount = $3, quantity = $4, status = $5, expiration_date = $6, id_user = $7 WHERE id_shipping = $8",
		updateS.ShippingNumber, updateS.Weight, updateS.Amount, updateS.Quantity, updateS.Status, updateS.ExpirationDate, updateS.IDUser, updateS.IDShipping)
	if err != nil {
		return err
	}
	return nil
}

func SearchShipping(readS models.ReadShipping, db *sql.DB) (*sql.Rows, error) {
	query := `SELECT s.*, u.email
              FROM shippings s
              JOIN users u ON s.id_user = u.id_user
              WHERE 1=1`

	var args []interface{}
	argIndex := 1 // Contador dinámico para los índices de los parámetros

	// Condición para Email
	if readS.Email != "" {
		query += fmt.Sprintf(" AND u.email = $%d", argIndex)
		args = append(args, readS.Email)
		argIndex++
	}

	// Condición para ShippingNumber
	if readS.ShippingNumber != "" {
		query += fmt.Sprintf(" AND s.shipping_number = $%d", argIndex)
		args = append(args, readS.ShippingNumber)
		argIndex++
	}

	// Condición para IDShipping
	if readS.IDShipping != 0 {
		query += fmt.Sprintf(" AND s.id_shipping = $%d", argIndex)
		args = append(args, readS.IDShipping)
		argIndex++
	}

	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func StatusShipping(idShipping int, db *sql.DB) (int, error) {
	var status int
	err := db.QueryRow("SELECT status FROM Shippings WHERE id_shipping = $1", idShipping).Scan(&status)
	if err != nil {
		return 0, err
	}
	return status, nil
}

func QuantityShipping(quantityShipping models.QuantityShipping, db *sql.DB) (int, error) {
	var quantity int
	err := db.QueryRow("SELECT COALESCE(SUM(quantity),0) AS quantity FROM shippings WHERE id_user = $1 and status = $2",
		quantityShipping.IDUser, quantityShipping.Status).Scan(&quantity)
	fmt.Println(err)
	if err != nil {
		return 0, err
	}
	return quantity, nil
}
