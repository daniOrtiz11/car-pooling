package database

import (
	"log"
	"time"
)

/*
InsertJourney is a
*/
func InsertJourney(id int, people int, state int, table int) bool {

	db := getConnection()
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
UpdateStatusJourneyByID is a
*/
func UpdateStatusJourneyByID(id int, newStatus int) bool {
	db := getConnection()
	sqlStatement := `UPDATE "table-booking-sch"."BOOKINGS" SET "status" = $1 WHERE id = $2`
	_, err := db.Exec(sqlStatement, newStatus, id)
	defer db.Close()
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

/*
FindBookingByID is a
*/
func FindBookingByID(idToSearch int) (int, int, int, int) {
	db := getConnection()
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
TruncateBookings is a
*/
func TruncateBookings() error {
	db := getConnection()
	sqlStatement := `TRUNCATE TABLE "table-booking-sch"."BOOKINGS"`
	_, err := db.Exec(sqlStatement)
	defer db.Close()
	return err
}

/*
	INSERT INTO users (age, email, first_name, last_name)
	VALUES ($1, $2, $3, $4)
	RETURNING id`
	id := 0
	err := db.QueryRow(sqlStatement, 30, "jon@calhoun.io", "Jonathan", "Calhoun").Scan(&id)
*/
