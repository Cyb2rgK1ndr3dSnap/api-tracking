package tasks

import (
	"database/sql"
	"log"
	"time"
)

func StartScheduler(db *sql.DB) {
	go func() {
		for {
			now := time.Now()
			nextRun := time.Date(
				now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location(),
			)
			if now.After(nextRun) {
				nextRun = nextRun.Add(24 * time.Hour)
			}

			waitDuration := time.Until(nextRun)
			log.Printf("La próxima tarea se ejecutará en: %v\n", waitDuration)

			time.Sleep(waitDuration)

			GenerateLateFeeTransactions(db)
		}
	}()
}

func GenerateLateFeeTransactions(db *sql.DB) {
	//
	query := `
		INSERT INTO transactions (id_user, id_shipping, id_transaction_type, transaction_amount)
		SELECT id_user, id_shipping, 3, amount * 0.15
		FROM shippings s
		WHERE CURRENT_DATE > expiration_date
		  AND status NOT IN (1, 3)
		  AND NOT EXISTS (
		      SELECT 1
				FROM transactions t
				WHERE t.id_shipping = s.id_shipping
				AND t.id_transaction_type = 3
		  );
	`
	//########//

	_, err := db.Exec(query)
	if err != nil {
		log.Printf("Error al ejecutar la tarea de Late Fees: %v", err)
	} else {
		log.Println("Tarea de Late Fees ejecutada correctamente.")
	}
}
