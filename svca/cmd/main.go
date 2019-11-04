package main

import (
	"fmt"

	"github.com/18981119465/go2sky-example/svca/data"
	"github.com/18981119465/go2sky-example/svca/service"
)

var appName = "account-service"

// env variable: SKYWALKING_BACKEND_SERVICE, SVCA_NAME, SVCA_PORT, SVCB_HOST, SVCB_PORT

func main() {
	fmt.Printf("Starting %v\n", appName)
	initDB()
	service.StartWebServer()
}

// Creates instance and calls the OpenBoltDb and Seed funcs
func initDB() {
	service.DBClient = &data.BoltClient{}
	service.DBClient.OpenDb()
	service.DBClient.Seed()
}
