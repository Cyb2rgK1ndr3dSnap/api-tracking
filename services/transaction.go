package services

import (
	"database/sql"
	"fmt"

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
	_, err := tx.Exec("UPDATE transactions SET id_user=$1, transaction_amount=$2 WHERE id_shipping = $3",
		updateT.IDUser, updateT.Amount, updateT.IDShipping)
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
