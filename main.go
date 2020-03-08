package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type car struct {
	id         int
	capacity   int
	seats      int
	seatsTaken int
	journeying bool
}

var cars []car

type group struct {
	id         int
	people     int
	journeying bool
}

var groups []group

func statusRequest(w http.ResponseWriter, r *http.Request) {
	/*
		Indicate the service has started up correctly and is ready to accept requests.
		Responses:


		200 OK When the service is ready to receive requests.
	*/
	w.WriteHeader(http.StatusOK)
}

func journeyRequest(w http.ResponseWriter, r *http.Request) {
	/*
		A group of people requests to perform a journey.
		Body required The group of people that wants to perform the journey
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
	contentType := getContentType(r)

	if contentType != "application/json" {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		defer r.Body.Close()
		body, errBody := ioutil.ReadAll(r.Body)
		if errBody != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		newGroup, errUnMarshal := unMarshalGroup(body)
		if errUnMarshal != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		groups = append(groups, newGroup)
	}

}

func dropoffRequest(w http.ResponseWriter, r *http.Request) {
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

func locateRequest(w http.ResponseWriter, r *http.Request) {
	/*
		Given a group ID such that ID=X, return the car the group is traveling
		with, or no car if they are still waiting to be served.
		Body required A url encoded form with the group ID such that ID=X
		Content Type application/x-www-form-urlencoded
		Accept application/json
		Responses:


		200 OK With the car as the payload when the group is assigned to a car.

		204 No Content When the group is waiting to be assigned to a car.

		404 Not Found When the group is not to be found.

		400 Bad Request When there is a failure in the request format or the
		payload can't be unmarshalled.
	*/
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "post called"}`))
}

func carsRequest(w http.ResponseWriter, r *http.Request) {
	/*
		Load the list of available cars in the service and remove all previous data
		(existing journeys and cars). This method may be called more than once during
		the life cycle of the service.
		Body required The list of cars to load.
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

func main() {
	log.Println("Starting server...")
	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/status", statusRequest).Methods(http.MethodGet)
	api.HandleFunc("/journey", journeyRequest).Methods(http.MethodPost)
	api.HandleFunc("/dropoff", dropoffRequest).Methods(http.MethodPost)
	api.HandleFunc("/locate", locateRequest).Methods(http.MethodPost)
	api.HandleFunc("/cars", carsRequest).Methods(http.MethodPut)
	log.Fatal(http.ListenAndServe(":8080", r))
}
