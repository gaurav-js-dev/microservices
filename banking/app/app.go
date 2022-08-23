package app

import (
	"log"
	"net/http"
)

func Start() {
	//create own multiplexer
	mux := http.NewServeMux()

	// define routes
	mux.HandleFunc("/greet", greet)
	mux.HandleFunc("/customers", getAllCustomers)

	// starting server
	log.Fatal(http.ListenAndServe("localhost:8000", mux))

}
