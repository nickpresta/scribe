# goscribed

Goscribe generates PDF documents using a WebKit rendering engine.

To generate a PDF:

* Send a GET request with a `url` query string parameter, which points to webpage which will be rendered as a PDF.
* This URL is expected to be valid (i.e. domain resolves, returns a 2XX, 3XX status, etc).
* If this URL cannot be resolved or a successful GET is not possible, this request will return an HTTP 412 - Precondition Failed response.

## To run goscribed

Make sure that the whole goscribe repo is checked out appropriated via `go get github.com/nickpresta/goscribe`

1. Build goscribed: `go build github.com/nickpresta/gopdf/goscribed`
2. `mv goscribed /some/destination`
3. `./goscribed`

See `goscribed --help` for more information.
