package services

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"github.com/Cyb2rgK1ndr3dSnap/api-tracking/models"
)

func CreateTransaction(createT models.CreateTransaction, tx *sql.Tx) error {
	_, err := tx.Exec("INSERT INTO transactions (id_user, id_shipping, id_transaction_type, transaction_amount) VALUES ($1, $2, $3, $4)",
		createT.IDUser, createT.IDShipping, createT.IDTransactionType, createT.Amount)
	if err != nil {
		return err
	}
	return nil
}

func UpdateTransaction(updateT models.UpdateTransaction, tx *sql.Tx) error {
	var queryParts []string
	var args []interface{}
	var counter int = 1

	if updateT.IDUser != 0 {
		queryParts = append(queryParts, "id_user = $"+strconv.Itoa(counter))
		args = append(args, updateT.IDUser)
		counter++
	}

	if updateT.Amount != 0 {
		queryParts = append(queryParts, "transaction_amount = $"+strconv.Itoa(counter))
		args = append(args, updateT.Amount)
		counter++
	}

	query := "UPDATE transactions SET " + strings.Join(queryParts, ", ") + " WHERE id_shipping = $" + strconv.Itoa(counter)
	args = append(args, updateT.IDShipping)

	_, err := tx.Exec(query, args...)
	if err != nil {
		return err
	}

	return nil
}

func BalanceTransaction(userID int, db *sql.DB) (float64, error) {
	var balance float64
	err := db.QueryRow("SELECT COALESCE(SUM(transaction_amount), 0) AS balance FROM transactions WHERE id_user = $1",
		userID).Scan(&balance)
	fmt.Println(err)
	if err != nil {
		return 0, err
	}
	return balance, nil
}
