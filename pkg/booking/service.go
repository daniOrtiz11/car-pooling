package booking

import (
	"encoding/json"
	"net/http"

	"github.com/daniOrtiz11/table-booking/internal/database"
)

/*
Journey Pojo
*/
type booking struct {
	ID     int
	People int
}

//booking created, booking lastupdate

var bookings []booking

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
	j, errUnmarshal := unMarshalGroup(body)
	if errUnmarshal != nil {
		return http.StatusBadRequest
	}
	//mirar si hay coche disponible
	id := database.CheckAvailableCar(int(j.People))
	ok := false
	if id == 0 {
		ok = database.InsertJourney(j.ID, j.People, 1, id)
	} else {
		ok = database.InsertJourney(j.ID, j.People, 2, id)
	}

	//check insert correcty
	if ok == false {
		return http.StatusBadRequest
	}

	if id == 0 {
		return http.StatusAccepted
	}

	//update status table
	ok = database.UpdateStatusCarByID(id, 2)
	if ok == false {
		//eliminate reference to table in booking
		database.UpdateStatusJourneyByID(j.ID, 1)
		return http.StatusBadRequest
	}
	return http.StatusOK

}

func unMarshalGroup(b []byte) (booking, error) {
	var g booking
	if err := json.Unmarshal(b, &g); err != nil {
		return g, err
	}
	return g, nil
}

/*
	A group of people requests to perform a booking.
	Body required The group of people that wants to perform the booking
	Content Type application/json
	Sample:
	{
	  "id": 1,
	  "people": 4
	}
	Responses:


	200 OK or 202 Accepted When the group is registered correctly

	400 Bad Request When there is a failure in the request format or the
	payload can't be unmarshalled.
*/
