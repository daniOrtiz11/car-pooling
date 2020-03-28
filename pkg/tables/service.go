package tables

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/daniOrtiz11/table-booking/internal/database"
)

/*
Table Pojo
*/
type Table struct {
	ID     int
	Seats  int
	Status bool
}

//var tables []table

/*
Service is a
*/
type Service interface {
	ServiceImpl()
}

/*
ServiceImpl is a
*/
func ServiceImpl(body []byte) int {
	tables, errUnmarshal := unMarshalTablesByBytes(body)
	if errUnmarshal != nil {
		return http.StatusBadRequest
	}

	errTrunT := database.TruncateTables()
	errTrunB := database.TruncateBookings()

	if (errTrunT != nil) || (errTrunB != nil) {
		log.Println(errTrunT)
		log.Println(errTrunT)
		return http.StatusBadRequest
	}
	//errInsert := database.insertManyTables(tables)

	log.Println(tables)
	return 0
}

/*
UnMarshalTablesByBytes is a
*/
func unMarshalTablesByBytes(bi []byte) ([]Table, error) {
	var tables []Table
	if err := json.Unmarshal(bi, &tables); err != nil {
		return tables, err
	}
	return tables, nil
}
