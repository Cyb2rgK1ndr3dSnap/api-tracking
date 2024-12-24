package models

import "time"

type User struct {
	IDUser      int       `json:"id_user,omitempty"`
	FirstName   string    `json:"firstname,omitempty"`
	LastName    string    `json:"lastname,omitempty"`
	UserName    string    `json:"username,omitempty"`
	Email       string    `json:"email,omitempty"`
	Direction   string    `json:"direction,omitempty"`
	PhoneNumber string    `json:"phone_number,omitempty"`
	CC          string    `json:"cc,omitempty"`
	Password    string    `json:"-"`
	IDRole      int       `json:"id_role,omitempty"`
	CreatedDate time.Time `json:"created_date,omitempty"`
	Token       string    `json:"token,omitempty"`
}

type RegisterUser struct {
	FirstName       string `json:"firstname" binding:"required"`
	LastName        string `json:"lastname,omitempty"`
	UserName        string `json:"username" binding:"required"`
	Email           string `json:"email" binding:"required,email"`
	Direction       string `json:"direction,omitempty"`
	PhoneNumber     string `json:"phonenumber" binding:"required"`
	CC              string `json:"cc" binding:"required"`
	Password        string `json:"password" binding:"required,min=8"`
	ConfirmPassword string `json:"passwordc" binding:"required,eqfield=Password"`
	Role            int    `json:"-"`
}

type LoginUser struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Token    string `json:"token,omitempty"`
}
