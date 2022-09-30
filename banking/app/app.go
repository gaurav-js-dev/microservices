package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gaurav-js-dev/microservices/banking/domain"
	"github.com/gaurav-js-dev/microservices/banking/logger"
	"github.com/gaurav-js-dev/microservices/banking/service"
	"github.com/joho/godotenv"

	"github.com/gorilla/mux"
)

func envCheck() {
	err := godotenv.Load(".env")
	if err != nil {
		logger.Error("Cannot load env file")
	}
	envProps := []string{
		"MYSQL_HOST",
		"MYSQL_PORT",
		"MYSQL_ROOT_USERNAME",
		"MYSQL_ROOT_PASSWORD",
		"MYSQL_DB_NAME",
		"SERVER_ADDRESS",
	}
	for _, k := range envProps {
		if os.Getenv(k) == "" {
			logger.Fatal(fmt.Sprintf("Environment variable %s not defined. Terminating application...", k))
		}
	}
}

func Start() {
	envCheck()
	// create own multiplexer added gorilla mux
	router := mux.NewRouter()

	//wiring
	// ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	// define routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)
	// starting server
	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	logger.Info(fmt.Sprintf("Starting server on %s:%s ...", address, port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))

}
