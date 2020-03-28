package database

import "log"

/*
CheckAvailableCar is a
*/
func CheckAvailableCar(requiredSeats int) int {
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
UpdateStatusCarByID is a
*/
func UpdateStatusCarByID(id int, newStatus int) bool {
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
