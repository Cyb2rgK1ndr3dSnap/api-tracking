package models

type Locker struct {
	IDLocker     int    `json:"id_locker,omitempty"`
	PhoneNumber1 string `json:"phone_number_1,omitempty"`
	PhoneNumber2 string `json:"phone_number_2,omitempty"`
	Name         string `json:"name,omitempty"`
	Direction1   string `json:"direction_1,omitempty"`
	Direction2   string `json:"direction_2,omitempty"`
	City         string `json:"city,omitempty"`
	State        string `json:"state,omitempty"`
	PostalCode   string `json:"postal_code,omitempty"`
}
