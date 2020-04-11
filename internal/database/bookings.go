package database

import (
	"log"
	"time"
)

/*
InsertBooking will retrieve true after successful insertion.
In other case, will retrieve false.
*/
func InsertBooking(id int, people int, state int, table int) bool {
	db, errCon := getConnection()
	if errCon != nil {
		return false
	}
	sqlStatement := `INSERT INTO "table-booking-sch"."BOOKINGS"
	("id", "people", "status", "timestamp_created", "timestamp_last_updated", "table") 
	VALUES ($1, $2, $3, NOW(), NOW(), $4)`
	_, err := db.Exec(sqlStatement, id, people, state, table)
	defer db.Close()
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

/*
UpdateStatusBookingByID will retrieve true after successful update.
In other case, will retrieve false.
*/
func UpdateStatusBookingByID(id int, newStatus int) bool {
	db, errCon := getConnection()
	if errCon != nil {
		return false
	}
	sqlStatement := `UPDATE "table-booking-sch"."BOOKINGS" SET "status" = $1 WHERE "id" = $2`
	_, err := db.Exec(sqlStatement, newStatus, id)
	defer db.Close()
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

/*
FindBookingByID will retrieve the values to build a booking after successful search.
In other case, will retrieve multiple zeros.
*/
func FindBookingByID(idToSearch int) (int, int, int, int) {
	db, errCon := getConnection()
	if errCon != nil {
		return 0, 0, 0, 0
	}
	sqlStatement := `SELECT * FROM "table-booking-sch"."BOOKINGS" WHERE "id" = $1`
	id, people, status, table := 0, 0, 0, 0
	timeCreated, timeUpdated := time.Now(), time.Now()
	rows, err := db.Query(sqlStatement, idToSearch)
	defer db.Close()
	if err != nil {
		log.Println(err)
		return 0, 0, 0, 0
	}
	exits := rows.Next()
	if exits {
		err = rows.Scan(&id, &people, &status, &timeCreated, &timeUpdated, &table)
		if err != nil {
			log.Println(err)
			return 0, 0, 0, 0
		}
	}
	return id, people, status, table
}

/*
TruncateBookings will retrieve empty error after successful truncate.
In other case, will retrieve a filled error.
*/
func TruncateBookings() error {
	db, errCon := getConnection()
	if errCon != nil {
		return errCon
	}
	sqlStatement := `TRUNCATE TABLE "table-booking-sch"."BOOKINGS"`
	_, err := db.Exec(sqlStatement)
	defer db.Close()
	return err
}
