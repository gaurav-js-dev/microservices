package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gaurav-js-dev/go-banking-lib/logger"
	"github.com/gaurav-js-dev/microservices/banking/domain"
	"github.com/gaurav-js-dev/microservices/banking/service"
	"github.com/jmoiron/sqlx"
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

	router := mux.NewRouter()

	dbClient := getDbClient()
	customerRepositoryDb := domain.NewCustomerRepositoryDb(dbClient)
	accountRepositoryDb := domain.NewAccountRepositoryDb(dbClient)

	ch := CustomerHandlers{service.NewCustomerService(customerRepositoryDb)}
	ah := AccountHandler{service.NewAccountService(accountRepositoryDb)}

	// Define Routes

	router.
		HandleFunc("/customers", ch.getAllCustomers).
		Methods(http.MethodGet).
		Name("GetAllCustomers")
	router.
		HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).
		Methods(http.MethodGet).
		Name("GetCustomer")
	router.
		HandleFunc("/customers/{customer_id:[0-9]+}/account", ah.NewAccount).
		Methods(http.MethodPost).
		Name("NewAccount")
	router.
		HandleFunc("/customers/{customer_id:[0-9]+}/account/{account_id:[0-9]+}", ah.MakeTransaction).
		Methods(http.MethodPost).
		Name("NewTransaction")

	am := AuthMiddleware{domain.NewAuthRepository()}
	router.Use(am.authorizationHandler())

	// Starting Server
	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	logger.Info(fmt.Sprintf("Starting server on %s:%s ...", address, port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))

}
func getDbClient() *sqlx.DB {
	dbUser := os.Getenv("MYSQL_ROOT_USERNAME")
	dbPasswd := os.Getenv("MYSQL_ROOT_PASSWORD")
	dbAddr := os.Getenv("MYSQL_HOST")
	dbName := os.Getenv("MYSQL_DB_NAME")
	dbPort := os.Getenv("MYSQL_PORT")

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPasswd, dbAddr, dbPort, dbName)
	client, err := sqlx.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return client
}
