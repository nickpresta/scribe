package main

import (
	"flag"
	"fmt"
	"github.com/NickPresta/handlers" // My fork of gorilla/handlers with Apache Combined Log Format
	"github.com/NickPresta/scribe/goscribe"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"os"
	"runtime"
)

var (
	port            = flag.Int("port", 8080, "HTTP listen port")
	binary          = flag.String("binary", "", "PhantomJS binary location")
	script          = flag.String("script", "", "PhantomJS script location")
	httpLogLocation = flag.String("log", "/tmp/scribe.log", "HTTP log file")
	httpLogger      io.Writer
)

func stderr(s interface{}) {
	fmt.Fprintln(os.Stderr, s)
}

func setupLogging(location string) {
	var err error
	httpLogger, err = os.OpenFile(location, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0660)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	flag.Parse()
	if *binary == "" || *script == "" {
		stderr("Missing binary or script argument")
		stderr("goscribed -binary BINARY -script SCRIPT [-port PORT]\nFlags:")
		flag.PrintDefaults()
		os.Exit(1)
	}

	log.Printf("Now serving on http://localhost:%d\n", *port)

	numCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPU)

	log.Printf("Setting GOMAXPROCS to %d", numCPU)

	goscribe.SetPDFBinaryLocation(*binary)
	goscribe.SetPDFScriptLocation(*script)

	setupLogging(*httpLogLocation)

	router := mux.NewRouter()
	router.HandleFunc("/", goscribe.RequestHandler).Methods("GET")

	http.Handle("/", router)
	err := http.ListenAndServe(fmt.Sprintf(":%d", *port), handlers.CombinedLoggingHandler(httpLogger, http.DefaultServeMux))

	if err != nil {
		log.Fatal(err)
	}
}
