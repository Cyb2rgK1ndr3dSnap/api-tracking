package services

import (
	"database/sql"

	"github.com/Cyb2rgK1ndr3dSnap/api-tracking/models"
)

func CreateUser(registerUser models.RegisterUser, hashedPassword string, db *sql.DB) error {
	_, err := db.Exec("INSERT INTO users (firstname, lastname, email, direction, phone_number, id_role, cc, password) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
		registerUser.FirstName, registerUser.LastName, registerUser.Email, registerUser.Direction, registerUser.PhoneNumber, registerUser.Role, registerUser.CC, hashedPassword)
	if err != nil {
		return err
	}

	return nil
}

func GetUserByEmail(email string, db *sql.DB) (*models.User, error) {
	u := new(models.User)
	err := db.QueryRow("SELECT * FROM users WHERE email = $1", email).Scan(
		&u.IDUser,
		&u.FirstName,
		&u.LastName,
		&u.Email,
		&u.Direction,
		&u.PhoneNumber,
		&u.CC,
		&u.Password,
		&u.IDRole,
		&u.CreatedDate,
	)
	if err != nil {
		return nil, err
	}

	if u.IDUser == 0 {
		return nil, err
	}

	return u, nil
}

func GetUserByID(id int, db *sql.DB) (*models.User, error) {
	u := new(models.User)
	err := db.QueryRow("SELECT * FROM users WHERE id_user = ?", id).Scan(
		&u.IDUser,
		&u.FirstName,
		&u.LastName,
		&u.Email,
		&u.Direction,
		&u.PhoneNumber,
		&u.CC,
		&u.Password,
		&u.IDRole,
		&u.CreatedDate,
	)
	if err != nil {
		return nil, err
	}

	if u.IDUser == 0 {
		return nil, err
	}

	return u, nil
}
