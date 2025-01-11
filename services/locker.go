package services

import (
	"database/sql"

	"github.com/Cyb2rgK1ndr3dSnap/api-tracking/models"
)

func GetLocker(db *sql.DB) (*models.Locker, error) {
	l := new(models.Locker)
	err := db.QueryRow(`SELECT * FROM LOCKERS`).Scan(
		&l.IDLocker,
		&l.PhoneNumber1,
		&l.PhoneNumber2,
		&l.Name,
		&l.Direction1,
		&l.Direction2,
		&l.City,
		&l.State,
		&l.PostalCode,
	)

	if err != nil {
		return nil, err
	}

	return l, nil
}
