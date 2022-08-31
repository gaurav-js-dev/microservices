package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gaurav-js-dev/microservices/banking/domain"
	"github.com/gaurav-js-dev/microservices/banking/service"

	"github.com/gorilla/mux"
)

func Start() {
	//create own multiplexer added gorilla mux
	router := mux.NewRouter()

	//wiring
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}

	// define routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	// starting server
	log.Fatal(http.ListenAndServe("localhost:8000", router))

}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(w, vars["customer_id"])
}

func createCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Post request received")
}
