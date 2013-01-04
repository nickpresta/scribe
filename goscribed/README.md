# goscribed

Goscribe generates PDF documents using a WebKit rendering engine.

It has two ways of generating a PDF.

1. Send a GET request with a `url` query string parameter, which points to webpage which will be rendered as a PDF.
    This URL is expected to be valid (i.e. domain resolves, returns a 2XX, 3XX status, etc).
    If this URL cannot be resolved or a successful GET is not possible, this request will return an HTTP 412 - Precondition Failed response.

2. Send a POST request with a body containing the webpage string that should be rendered as a PDF.
    If the body is empty, this request will return an HTTP 412 - Precondition Failed response.

## To run goscribed

1. Build goscribed: `go build`
2. `mv goscribed /some/destination`
3. `./goscribed`

See `goscribed --help` for more information.
