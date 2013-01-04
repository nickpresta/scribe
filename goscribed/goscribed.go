package main

import (
	"../goscribe"
	"flag"
	"fmt"
	"log"
	"net/http"
)

type loggedResponseWriter struct {
	StatusCode     int
	ResponseWriter http.ResponseWriter
}

func (writer *loggedResponseWriter) WriteHeader(code int) {
	writer.StatusCode = code
	writer.ResponseWriter.WriteHeader(code)
}

func (writer *loggedResponseWriter) Header() http.Header {
	return writer.ResponseWriter.Header()
}

func (writer *loggedResponseWriter) Write(d []byte) (int, error) {
	return writer.ResponseWriter.Write(d)
}

func logHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		loggedWriter := &loggedResponseWriter{200, writer}
		handler.ServeHTTP(loggedWriter, request)
		log.Printf("%s %s %d %s %s", request.RemoteAddr, request.Method, loggedWriter.StatusCode, request.URL, request.UserAgent())
	})
}

var (
	port = flag.Int("port", 8080, "HTTP listen port")
)

func main() {
	flag.Parse()

	fmt.Printf("Now serving on http://localhost:%d\n", *port)

	http.HandleFunc("/", goscribe.RequestHandler)
	err := http.ListenAndServe(fmt.Sprintf(":%d", *port), logHandler(http.DefaultServeMux))
	if err != nil {
		log.Fatal(err)
	}
}
