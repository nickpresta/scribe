package goscribe

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os/exec"
)

var (
	logger            = log.New(ioutil.Discard, "", 0)
	pdfBinaryLocation = ""
	pdfScriptLocation = ""
)

func SetLog(w io.Writer) {
	logger = log.New(w, "package", log.LstdFlags|log.Lmicroseconds)
}

func SetPDFBinaryLocation(location string) {
	pdfBinaryLocation = location
}

func SetPDFScriptLocation(location string) {
	pdfScriptLocation = location
}

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

func LogHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		loggedWriter := &loggedResponseWriter{200, writer}
		handler.ServeHTTP(loggedWriter, request)
		log.Printf("%s %s %d %s %s", request.RemoteAddr, request.Method, loggedWriter.StatusCode, request.URL, request.UserAgent())
	})
}

// Generates a PDF document from a given URL.
// It is expected that the URL given is valid (i.e. domain resolves, returns 2XX, 3XX status)
// If the URL cannot be resolved or a successful GET is not possible, will return
// HTTP 412 - Precondition Failed
func generatePdfFromUrl(writer http.ResponseWriter, urlQuery string) {
	escapedURL, err := url.QueryUnescape(urlQuery)
	if err != nil {
		http.Error(writer, "Invalid encoded URL parameter", http.StatusPreconditionFailed)
	}

	_, err = http.Get(escapedURL)
	if err != nil {
		http.Error(writer, fmt.Sprintf("Could not GET url (%s): %v", escapedURL, err), http.StatusPreconditionFailed)
	}

	content, err := generatePdfFromString(urlQuery)
	if err != nil {
		http.Error(writer, "Could not generate PDF", http.StatusInternalServerError)
	}

	writer.Header().Set("Content-Type", "application/pdf; charset=utf-8")
	writer.Header().Set("Content-Disposition", "attachment; filename='download.pdf'")
	fmt.Fprintf(writer, "%s", content)
}

// Generates a PDF document from a given string.
func generatePdfFromString(urlQuery string) (string, error) {
	out, err := exec.Command(pdfBinaryLocation, pdfScriptLocation, urlQuery, "/dev/stdout", "Letter").Output()
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	return string(out), nil
}

// Accepts a GET request.
// It is required that the GET request has a URL query string parameter named `url` otherwise this method will return
// HTTP 412 - Precondition Failed
func getRequestHandler(writer http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	url := query.Get("url")

	if url != "" {
		generatePdfFromUrl(writer, url)
		return
	}

	http.Error(writer, "", http.StatusPreconditionFailed)
}

// Request Handler which delegates GET requests.
func RequestHandler(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "GET":
		getRequestHandler(writer, request)
	default:
		http.Error(writer, "501 Not Implemented", http.StatusNotImplemented)
	}
}
