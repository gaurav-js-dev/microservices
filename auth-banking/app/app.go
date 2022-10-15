package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gaurav-js-dev/microservices/auth-banking/domain"
	"github.com/gaurav-js-dev/microservices/auth-banking/service"
	"github.com/gaurav-js-dev/microservices/library/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
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

	authRepository := domain.NewAuthRepository(getDbClient())
	ah := AuthHandler{service.NewLoginService(authRepository, domain.GetRolePermissions())}

	// Define Routes

	router.HandleFunc("/auth/login", ah.Login).Methods(http.MethodPost)
	router.HandleFunc("/auth/register", ah.NotImplementedHandler).Methods(http.MethodPost)
	router.HandleFunc("/auth/refresh", ah.Refresh).Methods(http.MethodPost)
	router.HandleFunc("/auth/verify", ah.Verify).Methods(http.MethodGet)

	// Starting Server
	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT_AUTH")
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
