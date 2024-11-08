package models

import "time"

type User struct {
	IDUser      int       `json:"id_user,omitempty"`
	Name        string    `json:"name,omitempty"`
	Lastname    string    `json:"lastname,omitempty"`
	Email       string    `json:"email,omitempty"`
	Direction   string    `json:"direction,omitempty"`
	PhoneNumber string    `json:"phone_number,omitempty"`
	CC          string    `json:"cc,omitempty"`
	Password    string    `json:"password,omitempty"`
	CreatedDate time.Time `json:"created_date,omitempty"`
}
