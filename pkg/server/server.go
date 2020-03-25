package server

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/daniOrtiz11/table-booking/pkg/bill"
	"github.com/daniOrtiz11/table-booking/pkg/tables"

	"github.com/daniOrtiz11/table-booking/pkg/booking"

	"github.com/daniOrtiz11/table-booking/internal/utils"
	"github.com/daniOrtiz11/table-booking/pkg/locate"
	"github.com/gorilla/mux"
)

type api struct {
	router http.Handler

	locate  locate.Service
	booking booking.Service
	bill    bill.Service
	tables  tables.Service
}

/*
Server is a
*/
type Server interface {
	Router() http.Handler
	Addr() string

	locateRequest(w http.ResponseWriter, r *http.Request)
	statusRequest(w http.ResponseWriter, r *http.Request)
	bookingRequest(w http.ResponseWriter, r *http.Request)
	billRequest(w http.ResponseWriter, r *http.Request)
	tablesRequest(w http.ResponseWriter, r *http.Request)
}

func (a *api) Router() http.Handler {
	return a.router
}

func (a *api) Addr() string {
	return fmt.Sprintf("%s:%s", utils.GetEnv("SERVER_HOST", "127.0.0.1"), utils.GetEnv("SERVER_PORT", "8080"))
}

/*
New is a
*/
func New() Server {
	a := &api{}
	r := mux.NewRouter()
	/*
		r := mux.NewRouter()
		api := r.PathPrefix("/api/v1").Subrouter()
		api.HandleFunc("/status", statusRequest).Methods(http.MethodGet)
		api.HandleFunc("/booking", bookingRequest).Methods(http.MethodPost)
		api.HandleFunc("/bill", billRequest).Methods(http.MethodPost)
		api.HandleFunc("/locate", locateRequest).Methods(http.MethodPost)
		api.HandleFunc("/tables", tablesRequest).Methods(http.MethodPut)
		log.Fatal(http.ListenAndServe(Cfg.Server.Port, r))
	*/
	r.HandleFunc("/status", a.statusRequest).Methods(http.MethodGet)
	r.HandleFunc("/booking", a.bookingRequest).Methods(http.MethodPost)
	r.HandleFunc("/bill", a.billRequest).Methods(http.MethodPost)
	r.HandleFunc("/locate", a.locateRequest).Methods(http.MethodPost)
	r.HandleFunc("/tables", a.tablesRequest).Methods(http.MethodPut)
	a.router = r

	return a
}

func (a *api) locateRequest(w http.ResponseWriter, r *http.Request) {
	/*
		Given a group ID such that ID=X, return the table the group is traveling
		with, or no table if they are still waiting to be served.
		Body required A url encoded form with the group ID such that ID=X
		Content Type application/x-www-form-urlencoded
		Accept application/json
		Responses:
		200 OK With the table as the payload when the group is assigned to a table.
		204 No Content When the group is waiting to be assigned to a table.
		404 Not Found When the group is not to be found.
		400 Bad Request When there is a failure in the request format or the
		payload can't be unmarshalled.
	*/

	contentType := utils.GetContentType(r)
	accept := utils.GetAccept(r)
	if (contentType != "application/x-www-form-urlencoded") || (accept != "application/json") {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		r.ParseForm()
		id, errArg := strconv.Atoi(r.FormValue("ID"))
		if errArg != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		status := locate.ServiceImpl(id)
		w.WriteHeader(status)
	}
}

func (a *api) statusRequest(w http.ResponseWriter, r *http.Request) {
	/*
		Indicate the service has started up correctly and is ready to accept requests.
		Responses:


		200 OK When the service is ready to receive requests.
	*/
	w.WriteHeader(http.StatusOK)
}

func (a *api) bookingRequest(w http.ResponseWriter, r *http.Request) {
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
	contentType := utils.GetContentType(r)
	if contentType != "application/json" {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		defer r.Body.Close()
		body, errBody := ioutil.ReadAll(r.Body)
		//_, errBody := ioutil.ReadAll(r.Body)
		if errBody != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		status := booking.ServiceImpl(body)
		w.WriteHeader(status)
	}

}

func (a *api) billRequest(w http.ResponseWriter, r *http.Request) {
	/*
		A group of people requests to be dropped off. Whether they traveled or not.
		Body required A form with the group ID, such that ID=X
		Content Type application/x-www-form-urlencoded
		Responses:


		200 OK or 204 No Content When the group is unregistered correctly.

		404 Not Found When the group is not to be found.

		400 Bad Request When there is a failure in the request format or the
		payload can't be unmarshalled.
	*/
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "post called"}`))
}

func (a *api) tablesRequest(w http.ResponseWriter, r *http.Request) {
	/*
		Load the list of available tables in the service and remove all previous data
		(existing bookings and tables). This method may be called more than once during
		the life cycle of the service.
		Body required The list of tables to load.
		Content Type application/json
		Sample:
		[
		  {
		    "id": 1,
		    "seats": 4
		  },
		  {
		    "id": 2,
		    "seats": 6
		  }
		]
		Responses:


		200 OK When the list is registered correctly.

		400 Bad Request When there is a failure in the request format, expected
		headers, or the payload can't be unmarshalled.
	*/
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(`{"message": "put called"}`))
}
