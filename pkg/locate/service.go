package locate

import (
	"log"
	"net/http"

	"github.com/daniOrtiz11/table-booking/internal/database"
	"github.com/daniOrtiz11/table-booking/internal/utils"
	"github.com/daniOrtiz11/table-booking/pkg/booking"
)

/*
Service is a
*/
type Service interface {
	ServiceImpl()
}

/*
ServiceImpl is a
*/
func ServiceImpl(idToSearch int) (int, int) {
	v1, v2, v3, v4 := database.FindBookingByID(idToSearch)
	if v1 == 0 {
		return http.StatusNotFound, 0
	}
	b, err := booking.UnMarshalGroupByValues(v1, v2, v3, v4)
	if err != nil {
		log.Fatal(err)
		return http.StatusNotFound, 0
	}

	if b.Status == utils.WATING {
		return http.StatusNoContent, 0
	} else if b.Status == utils.EATING {
		return http.StatusAccepted, b.Table
	} else {
		//COMPLETED booking
		return http.StatusNotFound, 0
	}

}
