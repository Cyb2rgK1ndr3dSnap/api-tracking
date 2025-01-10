package services

import (
	"database/sql"

	"github.com/Cyb2rgK1ndr3dSnap/api-tracking/models"
)

func GetStatus(db *sql.DB) ([]models.Status, error) {
	query := `SELECT * FROM status`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var status []models.Status
	for rows.Next() {
		var st models.Status
		err := rows.Scan(
			&st.IDStatus,
			&st.Name,
		)
		if err != nil {
			return nil, err
		}
		status = append(status, st)
	}

	return status, nil
}
