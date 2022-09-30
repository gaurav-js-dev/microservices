package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/gaurav-js-dev/microservices/banking/app"
)

func main() {
	// Set custom env variable
	os.Setenv("CUSTOM", "500")

	// fetcha all env variables
	for _, element := range os.Environ() {
		variable := strings.Split(element, "=")
		fmt.Println(variable[0], "=>", variable[1])
	}

	// fetch specific env variables
	fmt.Println(os.Getenv("CUSTOM"))
	fmt.Println("GOROOT=>", os.Getenv("GOROOT"))
	app.Start()
}
