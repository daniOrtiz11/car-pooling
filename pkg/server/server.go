package server

import (
	"encoding/json"
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
		status, response := locate.ServiceImpl(id)
		w.WriteHeader(status)
		if response != 0 {
			json.NewEncoder(w).Encode(response)
		}
	}
}

func (a *api) statusRequest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (a *api) bookingRequest(w http.ResponseWriter, r *http.Request) {
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
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "post called"}`))
}

func (a *api) tablesRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(`{"message": "put called"}`))
}
