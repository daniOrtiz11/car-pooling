package database

import "log"

/*
InsertTable is a
*/
func InsertTable(id int, seats int, status int) bool {

	db := getConnection()
	sqlStatement := `INSERT INTO "table-booking-sch"."TABLES"
	("id", "seats", "status", "timestamp_created", "timestamp_last_updated") 
	VALUES ($1, $2, $3, NOW(), NOW())`
	_, err := db.Exec(sqlStatement, id, seats, status)
	defer db.Close()
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

/*
CheckAvailableTable is a
*/
func CheckAvailableTable(requiredSeats int) int {
	db := getConnection()
	sqlStatement := `SELECT "id" FROM "table-booking-sch"."TABLES" WHERE "status" = 1 AND "seats" >= $1`
	id := 0
	rows, err := db.Query(sqlStatement, requiredSeats)
	defer db.Close()
	if err != nil {
		log.Println(err)
		return 0
	}
	exits := rows.Next()
	if exits {
		err = rows.Scan(&id)
		if err != nil {
			log.Println(err)
			return 0
		}
	}
	return id
}

/*
UpdateStatusTableByID is a
*/
func UpdateStatusTableByID(id int, newStatus int) bool {
	db := getConnection()
	sqlStatement := `UPDATE "table-booking-sch"."TABLES" SET "status" = $1 WHERE id = $2`
	err := db.QueryRow(sqlStatement, newStatus, id)
	defer db.Close()
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

/*
TruncateTables is a
*/
func TruncateTables() error {
	db := getConnection()
	sqlStatement := `TRUNCATE TABLE "table-booking-sch"."TABLES"`
	_, err := db.Exec(sqlStatement)
	defer db.Close()
	return err
}
