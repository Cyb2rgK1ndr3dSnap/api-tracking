package models

import "time"

type Shipping struct {
	IDShipping     int       `json:"id_shipping,omitempty"`
	IDUser         int       `json:"id_user"`
	ShippingNumber string    `json:"shipping_number"`
	Weight         float64   `json:"weight"`
	Value          float64   `json:"value"`
	Status         int       `json:"status"`
	CreationDate   time.Time `json:"creation_date,omitempty"`
	LastUpdate     time.Time `json:"last_update,omitempty"`
	ExpirationDate time.Time `json:"expiration_date"`
}
