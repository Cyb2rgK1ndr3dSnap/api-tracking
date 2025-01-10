package models

import "time"

type Shipping struct {
	IDShipping     int       `json:"id_shipping,omitempty"`
	IDUser         int       `json:"id_user,omitempty"`
	Email          string    `json:"email,omitempty"`
	Username       string    `json:"username,omitempty"`
	ShippingNumber string    `json:"shipping_number,omitempty"`
	Weight         float64   `json:"weight,omitempty"`
	Amount         float64   `json:"amount,omitempty"`
	Quantity       int       `json:"quantity,omitempty"`
	Status         int       `json:"status,omitempty"`
	Created_date   time.Time `json:"created_date,omitempty"`
	LastUpdate     time.Time `json:"last_update,omitempty"`
	ExpirationDate time.Time `json:"expiration_date,omitempty"`
}

type CreateShipping struct {
	IDUser         int       `json:"id_user"`
	Username       string    `json:"username" binding:"required"`
	ShippingNumber string    `json:"shipping_number" binding:"required"`
	Weight         float64   `json:"weight" binding:"required,gt=0"`
	Amount         float64   `json:"amount" binding:"required,gt=0"`
	Quantity       int       `json:"quantity" binding:"required,gt=0"`
	Status         int       `json:"-"`
	ExpirationDate time.Time `json:"expiration_date" binding:"required"`
}

type ReadShippings struct {
	IDUser         int    `form:"-"`
	IDRole         int    `form:"-"`
	Web            bool   `form:"web,omitempty"`
	IDShipping     int    `form:"id_shipping,omitempty"`
	Email          string `form:"email,omitempty"`
	Status         int    `form:"status,omitempty"`
	ShippingNumber string `form:"shipping_number,omitempty"`
	Cursor         int    `form:"cursor,omitempty"`
	PageNumber     int    `form:"page,omitempty" binding:"omitempty,gt=0"`
	PageSize       int    `form:"size" binding:"lte=100"`
}

type ReadShipping struct {
	IDUser         int `form:"-"`
	IDRole         int `form:"-"`
	ShippingStatus int `json:"status,omitempty"`
}

type UpdateShipping struct {
	IDShipping     int       `json:"id_shipping" binding:"required"`
	IDUser         int       `json:"-"`
	Username       string    `json:"username,omitempty"`
	ShippingNumber string    `json:"shipping_number,omitempty"`
	Weight         float64   `json:"weight,omitempty" binding:"omitempty,gt=0"`
	Amount         float64   `json:"amount,omitempty"  binding:"omitempty,gt=0"`
	Quantity       int       `json:"quantity,omitempty" binding:"omitempty,gt=0"`
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
