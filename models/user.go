package models

import "time"

type User struct {
	IDUser      int       `json:"id_user,omitempty"`
	FirstName   string    `json:"firstname,omitempty"`
	LastName    string    `json:"lastname,omitempty"`
	Email       string    `json:"email,omitempty"`
	Direction   string    `json:"direction,omitempty"`
	PhoneNumber string    `json:"phone_number,omitempty"`
	CC          string    `json:"cc,omitempty"`
	Password    string    `json:"password,omitempty"`
	IDRole      int       `json:"id_role,omitempty"`
	CreatedDate time.Time `json:"created_date,omitempty"`
	Token       string    `json:"token,omitempty"`
}

type RegisterUser struct {
	FirstName       string `json:"firstname" binding:"required"`
	LastName        string `json:"lastname,omitempty"`
	Email           string `json:"email" binding:"required"`
	Direction       string `json:"direction,omitempty"`
	PhoneNumber     string `json:"phone_number" binding:"required"`
	CC              string `json:"cc,omitempty"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" binding:"required"`
	Role            int    `json:"role" binding:"required"`
}

type LoginUser struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
