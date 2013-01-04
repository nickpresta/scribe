package goscribe

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// Generates a PDF document from a given URL.
// It is expected that the URL given is valid (i.e. domain resolves, returns 2XX, 3XX status)
// If the URL cannot be resolved or a successful GET is not possible, will return
// HTTP 412 - Precondition Failed
func generatePdfFromUrl(writer http.ResponseWriter, url string) string {
	response, err := http.Get(url)

	if err != nil {
		http.Error(writer, fmt.Sprintf("Could not GET url (%s): %v", url, err), http.StatusPreconditionFailed)
		return ""
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		http.Error(writer, fmt.Sprintf("Could read contents: %v", err), http.StatusInternalServerError)
		return ""
	}

	return generatePdfFromString(writer, string(body))
}

// Generates a PDF document from a given string.
func generatePdfFromString(writer http.ResponseWriter, body string) string {
	return body
}

// Accepts a GET request.
// It is required that the GET request has a URL query string parameter named `url` otherwise this method will return
// HTTP 412 - Precondition Failed
func getRequestHandler(writer http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	url := query.Get("url")

	if url != "" {
		fmt.Fprintf(writer, generatePdfFromUrl(writer, url))
		return
	}

	http.Error(writer, "Missing `url` parameter", http.StatusPreconditionFailed)
}

// Accepts a POST request.
// It is required that the POST request has a non-empty body otherwise this method will return
// HTTP 412 - Precondition Failed
func postRequestHandler(writer http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		http.Error(writer, fmt.Sprintf("Could not parse POST body: %v", err), http.StatusInternalServerError)
	}

	bodyContent := string(body)
	if bodyContent != "" {
		fmt.Fprintf(writer, generatePdfFromString(writer, bodyContent))
		return
	}

	http.Error(writer, "Cannot parse empty body", http.StatusPreconditionFailed)
}

// Request Handler which delegates GET and POST requests.
func RequestHandler(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "GET":
		getRequestHandler(writer, request)
	case "POST":
		postRequestHandler(writer, request)
	default:
		http.Error(writer, "501 Not Implemented", http.StatusNotImplemented)
	}
}
