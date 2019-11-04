package service

import (
	"log"
	"net/http"
	"os"

	"github.com/tetratelabs/go2sky"
	go2sky_http "github.com/tetratelabs/go2sky/plugins/http"
	"github.com/tetratelabs/go2sky/reporter"
)

// Tracer Global Variable
var Tracer *go2sky.Tracer

// StartWebServer starts the default http web server on the given port
func StartWebServer() {
	port := os.Getenv("SVCA_PORT")
	r, err := reporter.NewGRPCReporter(os.Getenv("SKYWALKING_BACKEND_SERVICE"))
	if err != nil {
		log.Fatalf("new reporter error %v \n", err)
	}
	defer r.Close()

	Tracer, err = go2sky.NewTracer(os.Getenv("SVCA_NAME"), go2sky.WithReporter(r))
	if err != nil {
		log.Fatalf("create tracer error %v \n", err)
	}
	Tracer.WaitUntilRegister()

	sm, err := go2sky_http.NewServerMiddleware(Tracer)
	if err != nil {
		log.Fatalf("create server middleware error %v \n", err)
	}

	router := NewRouter()
	// set default http default handler
	http.Handle("/", router)

	// start default http server
	log.Println("Starting HTTP service at " + port)
	err1 := http.ListenAndServe(":"+port, sm(router)) // Goroutine will block here

	if err1 != nil {
		log.Println("An error occurred starting HTTP listener at port " + port)
		log.Println("Error: " + err.Error())
	}
}
