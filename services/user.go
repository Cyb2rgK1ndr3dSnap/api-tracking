package services

import (
	"database/sql"
	"fmt"

	"github.com/Cyb2rgK1ndr3dSnap/api-tracking/models"
)

func RegisterUser(registerUser models.RegisterUser, hashedPassword string, db *sql.DB) error {
	_, err := db.Exec("INSERT INTO users (firstname, lastname, username, email, direction, phone_number, id_role, cc, password) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)",
		registerUser.FirstName, registerUser.LastName, registerUser.UserName, registerUser.Email, registerUser.Direction, registerUser.PhoneNumber, registerUser.Role, registerUser.CC, hashedPassword)
	if err != nil {
		return err
	}
	return nil
}

func GetUserByUsername(username string, db *sql.DB) (*models.User, error) {
	u := new(models.User)
	err := db.QueryRow("SELECT id_user,firstname,lastname,email,direction,phone_number,cc,password,id_role,username FROM users WHERE username = $1", username).Scan(
		&u.IDUser,
		&u.FirstName,
		&u.LastName,
		&u.Email,
		&u.Direction,
		&u.PhoneNumber,
		&u.CC,
		&u.Password,
		&u.IDRole,
		&u.UserName,
	)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func GetUserByEmail(email string, db *sql.DB) (*models.User, error) {
	u := new(models.User)
	err := db.QueryRow("SELECT id_user,firstname,lastname,email,direction,phone_number,cc,password,id_role,username FROM users WHERE email = $1", email).Scan(
		&u.IDUser,
		&u.FirstName,
		&u.LastName,
		&u.Email,
		&u.Direction,
		&u.PhoneNumber,
		&u.CC,
		&u.Password,
		&u.IDRole,
		&u.UserName,
	)
	if err != nil {
		/*if err == sql.ErrNoRows {
			return nil, nil
		}*/
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

func GetUserTotal(db *sql.DB, readU models.ReadUser) (int, error) {
	var total int

	query := `SELECT COUNT(*) AS total FROM users WHERE 1 = 1 `

	var args []interface{}
	argIndex := 1

	if readU.UserType != 0 && readU.IDRole == 1 {
		query += fmt.Sprintf(" AND id_role = $%d", argIndex)
		args = append(args, readU.UserType)
		argIndex++
	}

	err := db.QueryRow(query, args...).Scan(&total)
	if err != nil {
		return 0, err
	}

	return total, nil
}
