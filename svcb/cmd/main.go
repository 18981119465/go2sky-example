package main

import (
	"fmt"

	"github.com/18981119465/go2sky-example/svcb/service"
)

var appName = "quote-service"

// env variable: SKYWALKING_BACKEND_SERVICE, SVCB_HOST, SVCB_PORT

func main() {
	fmt.Printf("Starting %v\n", appName)
	service.StartWebServer()
}
