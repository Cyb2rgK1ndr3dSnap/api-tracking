package services

import (
	"database/sql"
	"fmt"

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
	rows, err := db.Query("SELECT * FROM users WHERE email = $1", email)
	if err != nil {
		return nil, err
	}

	u := new(models.User)
	for rows.Next() {
		u, err = scanRowsIntoUser(rows)

		if err != nil {
			fmt.Println("USER", err)
			return nil, err
		}
	}

	if u.IDUser == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return u, nil
}

func GetUserByID(id int, db *sql.DB) (*models.User, error) {
	rows, err := db.Query("SELECT * FROM users WHERE id_user = ?", id)
	if err != nil {
		return nil, err
	}

	u := new(models.User)
	for rows.Next() {
		u, err = scanRowsIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if u.IDUser == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return u, nil
}

func scanRowsIntoUser(rows *sql.Rows) (*models.User, error) {
	user := new(models.User)

	err := rows.Scan(
		&user.IDUser,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Direction,
		&user.PhoneNumber,
		&user.CC,
		&user.Password,
		&user.IDRole,
		&user.CreatedDate,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}
