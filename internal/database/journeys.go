package database

import (
	"log"
)

/*
InsertJourney is a
*/
func InsertJourney(id int, people int, state int, car int) bool {

	db := getConnection()
	sqlStatement := `INSERT INTO "car-pooling-sch"."JOURNEYS"
	("id", "people", "status", "timestamp_created", "timestamp_last_updated", "car") 
	VALUES ($1, $2, $3, NOW(), NOW(), $4)`
	_, err := db.Exec(sqlStatement, id, people, state, car)
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
	sqlStatement := `UPDATE "car-pooling-sch"."JOURNEYS" SET "status" = $1 WHERE id = $2`
	_, err := db.Exec(sqlStatement, newStatus, id)
	defer db.Close()
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

/*
FindJourneyByID is a
*/
func FindJourneyByID(idToSearch int) int {
	db := getConnection()
	sqlStatement := `SELECT "*" FROM "car-pooling-sch"."JOURNEYS" WHERE "id" = $1`
	id := 0
	rows, err := db.Query(sqlStatement, idToSearch)
	defer db.Close()
	if err != nil {
		log.Println(err)
		return 0
	}
	exits := rows.Next()
	if exits {
		err = rows.Scan(id)
		if err != nil {
			log.Println(err)
			return 0
		}
	}
	return id
}

/*
	INSERT INTO users (age, email, first_name, last_name)
	VALUES ($1, $2, $3, $4)
	RETURNING id`
	id := 0
	err := db.QueryRow(sqlStatement, 30, "jon@calhoun.io", "Jonathan", "Calhoun").Scan(&id)
*/
