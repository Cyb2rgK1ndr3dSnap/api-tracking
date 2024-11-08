package models

import "time"

type Debt struct {
	IDDebt       int       `json:"id_debt,omitempty"`
	IDUser       int       `json:"id_user,omitempty"`
	IDShipping   int       `json:"id_shipping,omitempty"`
	Value        float64   `json:"value,omitempty"`
	Status       int       `json:"status,omitempty"`
	CreationDate time.Time `json:"creation_date,omitempty"`
	LastUpdate   time.Time `json:"last_update,omitempty"`
}
