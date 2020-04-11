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

//var bookings []Booking

/*
Service is a interface to define the methods
*/
type Service interface {
	ServiceImpl(body []byte)
}

/*
ServiceImpl will retrieve 200 or 202 http status after successful operation.
In other case, will retrieve 400 http status.
*/
func ServiceImpl(body []byte) int {
	j, errUnmarshal := UnMarshalGroupByBytes(body)
	if errUnmarshal != nil {
		return http.StatusBadRequest
	}
	//check for availabe tables
	id := database.CheckAvailableTable(int(j.People))
	ok := false
	if id == 0 {
		ok = database.InsertBooking(j.ID, j.People, utils.WAITING, id)
	} else {
		ok = database.InsertBooking(j.ID, j.People, utils.EATING, id)
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
		database.UpdateStatusBookingByID(j.ID, utils.WAITING)
		return http.StatusBadRequest
	}
	return http.StatusOK

}

/*
UnMarshalGroupByBytes will retrieve a Booking entity and empty error after successful unmarshal by bytes.
In other case, will retrieve a filled error.
*/
func UnMarshalGroupByBytes(bi []byte) (Booking, error) {
	var bo Booking
	if err := json.Unmarshal(bi, &bo); err != nil {
		return bo, err
	}
	return bo, nil
}

/*
UnMarshalGroupByValues will retrieve a Booking entity and empty error after successful unmarshal by values.
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
