package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kirega/banking/domain"
	"github.com/kirega/banking/service"
)

func Start() {

	router := mux.NewRouter()
	// wiring
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	// define the routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)

	// start the server
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
