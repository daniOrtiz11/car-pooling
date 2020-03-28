package booking

import (
	"encoding/json"
	"net/http"

	"github.com/daniOrtiz11/table-booking/internal/database"
	"github.com/daniOrtiz11/table-booking/internal/utils"
)

/*
Booking Pojo
*/
type Booking struct {
	ID     int
	People int
	Status int
	Table  int
}

//booking created, booking lastupdate

var bookings []Booking

/*
Service is a
*/
type Service interface {
	ServiceImpl(body []byte)
}

/*
ServiceImpl is a
*/
func ServiceImpl(body []byte) int {
	j, errUnmarshal := UnMarshalGroupByBytes(body)
	if errUnmarshal != nil {
		return http.StatusBadRequest
	}
	//mirar si hay coche disponible
	id := database.CheckAvailableCar(int(j.People))
	ok := false
	if id == 0 {
		ok = database.InsertJourney(j.ID, j.People, utils.WATING, id)
	} else {
		ok = database.InsertJourney(j.ID, j.People, utils.EATING, id)
	}

	//check insert correcty
	if ok == false {
		return http.StatusBadRequest
	}

	if id == 0 {
		return http.StatusAccepted
	}

	//update status table
	ok = database.UpdateStatusTableByID(id, utils.EATING)
	if ok == false {
		//eliminate reference to table in booking
		database.UpdateStatusJourneyByID(j.ID, utils.WATING)
		return http.StatusBadRequest
	}
	return http.StatusOK

}

/*
UnMarshalGroupByBytes is a
*/
func UnMarshalGroupByBytes(bi []byte) (Booking, error) {
	var bo Booking
	if err := json.Unmarshal(bi, &bo); err != nil {
		return bo, err
	}
	return bo, nil
}

/*
UnMarshalGroupByValues is a
*/
func UnMarshalGroupByValues(v1 int, v2 int, v3 int, v4 int) (Booking, error) {
	bo := Booking{
		ID:     v1,
		People: v2,
		Status: v3,
		Table:  v4,
	}
	return bo, nil
}
