package locate

import (
	"log"
	"net/http"

	"github.com/daniOrtiz11/table-booking/internal/database"
	"github.com/daniOrtiz11/table-booking/internal/utils"
	"github.com/daniOrtiz11/table-booking/pkg/booking"
)

/*
Service is a interface to define the methods
*/
type Service interface {
	ServiceImpl()
}

/*
ServiceImpl will retrieve 200 or 202 http status and table's id after successful operation.
In other case, will retrieve 400 http status and zero.
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

	if b.Status == utils.WAITING {
		return http.StatusNoContent, 0
	} else if b.Status == utils.EATING {
		return http.StatusOK, b.Table
	} else {
		//completed booking
		return http.StatusNotFound, 0
	}

}
