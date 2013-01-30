package main

import (
	"github.com/NickPresta/scribe/goscribe"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
)

var (
	port   = flag.Int("port", 8080, "HTTP listen port")
	binary = flag.String("binary", "", "PhantomJS binary location")
	script = flag.String("script", "", "PhantomJS script location")
)

func stderr(s interface{}) {
	fmt.Fprintln(os.Stderr, s)
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

	goscribe.SetLog(os.Stdout)
	goscribe.SetPDFBinaryLocation(*binary)
	goscribe.SetPDFScriptLocation(*script)

	http.HandleFunc("/", goscribe.RequestHandler)
	err := http.ListenAndServe(fmt.Sprintf(":%d", *port), goscribe.LogHandler(http.DefaultServeMux))
	if err != nil {
		log.Fatal(err)
	}
}
