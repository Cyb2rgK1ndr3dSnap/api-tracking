package services

import (
	"database/sql"

	"github.com/Cyb2rgK1ndr3dSnap/api-tracking/models"
)

func RegisterDevice(registerToken models.RegisterToken, db *sql.DB) error {
	_, err := db.Exec("INSERT INTO devicestokens (id_user, token) VALUES ($1, $2)",
		registerToken.IDUser, registerToken.Token)
	if err != nil {
		return err
	}
	return nil
}

func ReadDevices(IDUser int, db *sql.DB) (*sql.Rows, error) {
	rows, err := db.Query("SELECT token FROM devicestokens WHERE id_user = $1", IDUser)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return rows, nil
}
