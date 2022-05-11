package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {

	router := mux.NewRouter()
	// define the routes
	router.HandleFunc("/greet", greet).Methods(http.MethodGet)
	router.HandleFunc("/customers", getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customer/{customer_id:[0-9]+}", getCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customer/", createCustomer).Methods(http.MethodPost)
	router.HandleFunc("/customer/", createCustomer).Methods(http.MethodPost)

	// start the server
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
