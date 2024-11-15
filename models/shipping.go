package models

import "time"

type Shipping struct {
	IDShipping     int       `json:"id_shipping,omitempty"`
	IDUser         int       `json:"id_user"`
	Email          string    `json:"email"`
	ShippingNumber string    `json:"shipping_number"`
	Weight         float64   `json:"weight"`
	Amount         float64   `json:"amount"`
	Quantity       int       `json:"quantity"`
	Status         int       `json:"status"`
	CreationDate   time.Time `json:"creation_date,omitempty"`
	LastUpdate     time.Time `json:"last_update,omitempty"`
	ExpirationDate time.Time `json:"expiration_date"`
}

type CreateShipping struct {
	IDUser         int       `json:"id_user"`
	Email          string    `json:"email" binding:"required"`
	ShippingNumber string    `json:"shipping_number" binding:"required"`
	Weight         float64   `json:"weight" binding:"required"`
	Amount         float64   `json:"amount" binding:"required"`
	Quantity       int       `json:"quantity" binding:"required"`
	Status         int       `json:"status" binding:"required"`
	ExpirationDate time.Time `json:"expiration_date" binding:"required"`
}

type ReadShipping struct {
	Email          string `form:"email,omitempty"`
	ShippingNumber string `form:"shipping_number,omitempty"`
}

type UpdateShipping struct {
	IDUser         int       `json:"id_user"`
	Email          string    `json:"email"`
	ShippingNumber string    `json:"shipping_number"`
	Weight         float64   `json:"weight"`
	Amount         float64   `json:"amount"`
	Quantity       int       `json:"quantity"`
	Status         int       `json:"status"`
	ExpirationDate time.Time `json:"expiration_date"`
}
