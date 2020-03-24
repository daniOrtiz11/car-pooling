package database

import "log"

/*
CheckAvailableCar is a
*/
func CheckAvailableCar(requiredSeats int) int {
	db := getConnection()
	sqlStatement := `SELECT "id" FROM "car-pooling-sch"."CARS" WHERE "status" = 1 AND "seats" >= $1`
	id := 0
	rows, err := db.Query(sqlStatement, requiredSeats)
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
UpdateStatusCarByID is a
*/
func UpdateStatusCarByID(id int, newStatus int) bool {
	db := getConnection()
	sqlStatement := `UPDATE "car-pooling-sch"."CARS" SET "status" = $1 WHERE id = $2`
	err := db.QueryRow(sqlStatement, newStatus, id)
	defer db.Close()
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

/*
	INSERT INTO users (age, email, first_name, last_name)
	VALUES ($1, $2, $3, $4)
	RETURNING id`
	id := 0
	err := db.QueryRow(sqlStatement, 30, "jon@calhoun.io", "Jonathan", "Calhoun").Scan(&id)
*/
