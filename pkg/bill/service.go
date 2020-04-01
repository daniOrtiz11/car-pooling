package bill

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
func ServiceImpl(idToSearch int) int {
	completed := database.UpdateStatusBookingByID(idToSearch, utils.COMPLETED)
	if completed == false {
		return http.StatusNotFound
	}
	v1, v2, v3, v4 := database.FindBookingByID(idToSearch)
	if v1 == 0 {
		return http.StatusNotFound
	}
	b, err := booking.UnMarshalGroupByValues(v1, v2, v3, v4)
	if err != nil {
		log.Fatal(err)
		return http.StatusNotFound
	}

	completed = database.UpdateStatusTableByID(b.Table, utils.WAITING)
	if completed == false {
		//undo db changes
		database.UpdateStatusBookingByID(idToSearch, utils.EATING)
		return http.StatusNotFound
	}
	return http.StatusOK
}
