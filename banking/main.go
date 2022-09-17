package main

import (
	"github.com/gaurav-js-dev/microservices/banking/app"
	"github.com/gaurav-js-dev/microservices/banking/logger"
)

func main() {
	logger.Info("Starting the application...")
	app.Start()
}
