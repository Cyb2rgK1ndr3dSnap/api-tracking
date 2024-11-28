package models

import "time"

type Transaction struct {
	IDTransaction     int       `json:"id_transaction"`
	IDUser            int       `json:"id_user"`
	IDShipping        int       `json:"id_shipping"`
	IDTransactionType int       `json:"id_transaction_type"`
	TransactionAmount float64   `json:"transaction_amount"`
	CreationDate      time.Time `json:"creation_date"`
}

type CreateTransaction struct {
	IDUser            int     `json:"id_user" binding:"required"`
	IDShipping        int     `json:"id_shipping" binding:"required"`
	IDTransactionType int     `json:"id_transaction_type" binding:"required"`
	Amount            float64 `json:"amount" binding:"required"`
}

type UpdateTransaction struct {
	IDUser            int     `json:"id_user" binding:"required"`
	IDShipping        int     `json:"id_shipping" binding:"required"`
	IDTransactionType int     `json:"id_transaction_type" binding:"required"`
	Amount            float64 `json:"amount" binding:"required"`
}

type BalanceTransaction struct {
	IDUser int `json:"id_user" binding:"required"`
}
