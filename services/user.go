package services

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

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

func GetUsers(readU models.ReadUser, db *sql.DB) ([]*models.User, error) {
	query := `SELECT id_user, firstname, lastname, COALESCE(username, '') AS username, email, direction, phone_number, cc, id_role 
			  FROM users
			  WHERE 1=1`

	var args []interface{}
	argIndex := 1

	if readU.UserType != 0 {
		query += fmt.Sprintf(" AND id_role = $%d", argIndex)
		args = append(args, readU.UserType)
		argIndex++
	}
	if readU.UserName != "" {
		query += fmt.Sprintf(" AND username = $%d", argIndex)
		args = append(args, readU.UserName)
		argIndex++
	}
	if readU.Email != "" {
		query += fmt.Sprintf(" AND email = $%d", argIndex)
		args = append(args, readU.Email)
		argIndex++
	}

	query += ` ORDER BY id_user DESC`

	rows, err := db.Query(query, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	defer rows.Close()

	var users []*models.User
	for rows.Next() {
		user := new(models.User)
		err := rows.Scan(
			&user.IDUser,
			&user.FirstName,
			&user.LastName,
			&user.UserName,
			&user.Email,
			&user.Direction,
			&user.PhoneNumber,
			&user.CC,
			&user.IDRole,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func UpdateUser(updateU models.UpdateUser, tx *sql.Tx) error {
	// Inicializar las partes de la consulta y los valores
	var queryParts []string
	var args []interface{}
	var counter int = 1

	if updateU.FirstName != "" {
		queryParts = append(queryParts, "firstname = $"+strconv.Itoa(counter))
		args = append(args, updateU.FirstName)
		counter++
	}

	if updateU.LastName != "" {
		queryParts = append(queryParts, "lastname = $"+strconv.Itoa(counter))
		args = append(args, updateU.LastName)
		counter++
	}

	if updateU.UserName != "" {
		queryParts = append(queryParts, "username = $"+strconv.Itoa(counter))
		args = append(args, updateU.UserName)
		counter++
	}

	if updateU.Email != "" {
		queryParts = append(queryParts, "email = $"+strconv.Itoa(counter))
		args = append(args, updateU.Email)
		counter++
	}

	if updateU.Direction != "" {
		queryParts = append(queryParts, "direction = $"+strconv.Itoa(counter))
		args = append(args, updateU.Direction)
		counter++
	}

	if updateU.PhoneNumber != "" {
		queryParts = append(queryParts, "phone_number = $"+strconv.Itoa(counter))
		args = append(args, updateU.PhoneNumber)
		counter++
	}

	if updateU.CC != "" {
		queryParts = append(queryParts, "cc = $"+strconv.Itoa(counter))
		args = append(args, updateU.CC)
		counter++
	}

	if updateU.IDRole != 0 {
		queryParts = append(queryParts, "id_role = $"+strconv.Itoa(counter))
		args = append(args, updateU.IDRole)
		counter++
	}

	query := "UPDATE users SET " + strings.Join(queryParts, ", ") + " WHERE id_user = $" + strconv.Itoa(counter)
	args = append(args, updateU.IDUser)

	// Ejecutar la consulta
	_, err := tx.Exec(query, args...)
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

func GetUserTotal(readU models.ReadUser, db *sql.DB) (int, error) {
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
