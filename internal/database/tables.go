package database

import "log"

/*
InsertTable will retrieve true after successful insertion.
In other case, will retrieve false.
*/
func InsertTable(id int, seats int, status int) bool {
	db, errCon := getConnection()
	if errCon != nil {
		return false
	}
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
CheckAvailableTable will retrieve the table's id after successful search.
In other case, will retrieve a zero.
*/
func CheckAvailableTable(requiredSeats int) int {
	db, errCon := getConnection()
	if errCon != nil {
		return 0
	}
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
UpdateStatusTableByID will retrieve true after successful update.
In other case, will retrieve a zero.
*/
func UpdateStatusTableByID(id int, newStatus int) bool {
	db, errCon := getConnection()
	if errCon != nil {
		return false
	}
	sqlStatement := `UPDATE "table-booking-sch"."TABLES" SET "status" = $1 WHERE "id" = $2`
	_, err := db.Exec(sqlStatement, newStatus, id)
	defer db.Close()
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

/*
TruncateTables will retrieve empty error after successful truncate.
In other case, will retrieve a filled error.
*/
func TruncateTables() error {
	db, errCon := getConnection()
	if errCon != nil {
		return errCon
	}
	sqlStatement := `TRUNCATE TABLE "table-booking-sch"."TABLES"`
	_, err := db.Exec(sqlStatement)
	defer db.Close()
	return err
}
