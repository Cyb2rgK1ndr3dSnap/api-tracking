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
	Created_date   time.Time `json:"created_date,omitempty"`
	LastUpdate     time.Time `json:"last_update,omitempty"`
	ExpirationDate time.Time `json:"expiration_date"`
}

type CreateShipping struct {
	IDUser         int       `json:"id_user"`
	Email          string    `json:"email" binding:"required"`
	ShippingNumber string    `json:"shipping_number" binding:"required"`
	Weight         float64   `json:"weight" binding:"required,gt=0"`
	Amount         float64   `json:"amount" binding:"required,gt=0"`
	Quantity       int       `json:"quantity" binding:"required,gt=0"`
	Status         int       `json:"-"`
	ExpirationDate time.Time `json:"expiration_date" binding:"required"`
}

type ReadShipping struct {
	IDUser         int    `form:"-"`
	IDRole         int    `form:"-"`
	IDShipping     int    `form:"id_shipping,omitempty"`
	Email          string `form:"email,omitempty"`
	ShippingNumber string `form:"shipping_number,omitempty"`
}

type UpdateShipping struct {
	IDShipping     int       `json:"id_shipping" binding:"required"`
	IDUser         int       `json:"id_user,omitempty"`
	Email          string    `json:"email,omitempty"`
	ShippingNumber string    `json:"shipping_number,omitempty"`
	Weight         float64   `json:"weight,omitempty" binding:"gt=0"`
	Amount         float64   `json:"amount,omitempty"  binding:"gt=0"`
	Quantity       int       `json:"quantity,omitempty" binding:"gt=0"`
	Status         int       `json:"status,omitempty"`
	ExpirationDate time.Time `json:"expiration_date,omitempty"`
}

type CloseShipping struct {
	IDShipping int `json:"id_shipping" binding:"required"`
}

type QuantityShipping struct {
	IDUser int `json:"-"`
	Status int `form:"status"`
}
