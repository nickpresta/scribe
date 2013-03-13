package goscribe

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os/exec"
)

var (
	pdfBinaryLocation = ""
	pdfScriptLocation = ""
)

func SetPDFBinaryLocation(location string) {
	pdfBinaryLocation = location
}

func SetPDFScriptLocation(location string) {
	pdfScriptLocation = location
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

	content, err := generatePdfFromString(escapedURL)
	if err != nil {
		http.Error(writer, "Could not generate PDF", http.StatusInternalServerError)
	}

	writer.Header().Set("Content-Type", "application/pdf; charset=utf-8")
	writer.Header().Set("Content-Disposition", "attachment; filename='download.pdf'")
	fmt.Fprintf(writer, "%s", content)
}

// Generates a PDF document from a given URL.
func generatePdfFromString(url string) (string, error) {
	out, err := exec.Command(pdfBinaryLocation, pdfScriptLocation, url, "/dev/stdout", "Letter").Output()
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	return string(out), nil
}

// Accepts a GET request.
// It is required that the GET request has a URL query string parameter named `url` otherwise this method will return
// HTTP 412 - Precondition Failed
func RequestHandler(writer http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	url := query.Get("url")

	if url != "" {
		generatePdfFromUrl(writer, url)
		return
	}

	http.Error(writer, "`url` query parameter required", http.StatusPreconditionFailed)
}
