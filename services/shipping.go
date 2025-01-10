package services

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

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
	// Inicializar las partes de la consulta y los valores
	var queryParts []string
	var args []interface{}
	var counter int = 1

	if updateS.ShippingNumber != "" {
		queryParts = append(queryParts, "shipping_number = $"+strconv.Itoa(counter))
		args = append(args, updateS.ShippingNumber)
		counter++
	}

	if updateS.Status != 0 {
		queryParts = append(queryParts, "status = $"+strconv.Itoa(counter))
		args = append(args, updateS.Status)
		counter++
	}

	if updateS.Weight != 0 {
		queryParts = append(queryParts, "weight = $"+strconv.Itoa(counter))
		args = append(args, updateS.Weight)
		counter++
	}

	if updateS.Amount != 0 {
		queryParts = append(queryParts, "amount = $"+strconv.Itoa(counter))
		args = append(args, updateS.Amount)
		counter++
	}

	if updateS.Quantity != 0 {
		queryParts = append(queryParts, "quantity = $"+strconv.Itoa(counter))
		args = append(args, updateS.Quantity)
		counter++
	}

	if !updateS.ExpirationDate.IsZero() {
		queryParts = append(queryParts, "expiration_date = $"+strconv.Itoa(counter))
		args = append(args, updateS.ExpirationDate)
		counter++
	}

	if updateS.IDUser != 0 {
		queryParts = append(queryParts, "id_user = $"+strconv.Itoa(counter))
		args = append(args, updateS.IDUser)
		counter++
	}

	query := "UPDATE shippings SET " + strings.Join(queryParts, ", ") + " WHERE id_shipping = $" + strconv.Itoa(counter)
	args = append(args, updateS.IDShipping)

	// Ejecutar la consulta
	_, err := tx.Exec(query, args...)
	if err != nil {
		return err
	}

	return nil
}

func GetShipping(readS models.ReadShippings, db *sql.DB) (*sql.Rows, error) {
	query := `SELECT s.id_shipping, s.id_user, s.shipping_number, s.weight, s.amount, s.quantity, s.status, s.last_update, s.expiration_date, u.email, u.username
			  FROM shippings s
			  JOIN users u ON s.id_user = u.id_user 
			  WHERE 1=1`

	var args []interface{}
	argIndex := 1

	// Aplicar filtros
	if readS.Email != "" && readS.IDRole == 1 {
		query += fmt.Sprintf(" AND u.email = $%d", argIndex)
		args = append(args, readS.Email)
		argIndex++
	} else if readS.ShippingNumber != "" {
		query += fmt.Sprintf(" AND s.shipping_number = $%d", argIndex)
		args = append(args, readS.ShippingNumber)
		argIndex++
	} else if !readS.Web {
		query += fmt.Sprintf(" AND s.id_user = $%d", argIndex)
		args = append(args, readS.IDUser)
		argIndex++
	}
	if readS.Status != 0 {
		query += fmt.Sprintf(" AND s.status = $%d", argIndex)
		args = append(args, readS.Status)
		argIndex++
	}
	if readS.Cursor != 0 {
		query += fmt.Sprintf(" AND s.id_shipping < $%d", argIndex)
		args = append(args, readS.Cursor)
		argIndex++
	}
	if readS.PageSize != 0 {
		// Ordenar y limitar los resultados
		query += fmt.Sprintf(" ORDER BY (CASE WHEN s.status = 1 THEN 1 ELSE 0 END) ASC, s.id_shipping DESC LIMIT $%d", argIndex)
		args = append(args, readS.PageSize)
		argIndex++
	}
	if readS.PageNumber != 0 {
		offSet := ((readS.PageNumber - 1) * readS.PageSize)
		query += fmt.Sprintf(" OFFSET $%d", argIndex)
		args = append(args, offSet)
		argIndex++
	}

	rows, err := db.Query(query, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return rows, nil
}

func GetShippingMaxPage(readS models.ReadShippings, db *sql.DB) (int, error) {
	var total int
	query := `SELECT COUNT(*) AS total
			  FROM shippings s
			  JOIN users u ON s.id_user = u.id_user 
			  WHERE 1=1`

	var args []interface{}
	argIndex := 1
	// Aplicar filtros
	if readS.Email != "" && readS.IDRole == 1 {
		query += fmt.Sprintf(" AND u.email = $%d", argIndex)
		args = append(args, readS.Email)
		argIndex++
	} else if readS.ShippingNumber != "" {
		query += fmt.Sprintf(" AND s.shipping_number = $%d", argIndex)
		args = append(args, readS.ShippingNumber)
		argIndex++
	} else if !readS.Web {
		query += fmt.Sprintf(" AND s.id_user = $%d", argIndex)
		args = append(args, readS.IDUser)
		argIndex++
	}
	if readS.Status != 0 {
		query += fmt.Sprintf(" AND s.status = $%d", argIndex)
		args = append(args, readS.Status)
		argIndex++
	}

	err := db.QueryRow(query, args...).Scan(&total)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		}
		return 0, err
	}

	return total, nil
}

func GetShippingByID(idShipping int, db *sql.DB) (*models.Shipping, error) {
	u := new(models.Shipping)
	err := db.QueryRow("SELECT id_shipping, id_user, shipping_number, amount, status FROM Shippings WHERE id_shipping = $1", idShipping).Scan(
		&u.IDShipping,
		&u.IDUser,
		&u.ShippingNumber,
		&u.Amount,
		&u.Status,
	)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func StatusShippingByID(idShipping int, db *sql.DB) (int, error) {
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
	if err != nil {
		return 0, err
	}
	return quantity, nil
}

func GetShippingTotal(db *sql.DB, readU models.ReadShipping) (int, error) {
	var total int

	query := `SELECT COUNT(*) AS total FROM shippings WHERE 1 = 1 `

	var args []interface{}
	argIndex := 1

	if readU.ShippingStatus != 0 && readU.IDRole == 1 {
		query += fmt.Sprintf(" AND status = $%d", argIndex)
		args = append(args, readU.ShippingStatus)
		argIndex++
	}

	err := db.QueryRow(query, args...).Scan(&total)
	if err != nil {
		return 0, err
	}

	return total, nil
}
