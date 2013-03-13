# Scribe

Scribe generates PDF documents using a WebKit rendering engine.

To generate a PDF:

* Send a GET request with a `url` query string parameter, which points to webpage which will be rendered as a PDF.
* This URL is expected to be valid (i.e. domain resolves, returns a 2XX, 3XX status, etc).
* If this URL cannot be resolved or a successful GET is not possible, this request will return an HTTP 412 - Precondition Failed response.

## Requirements

* [PhantomJS 1.9.0](https://github.com/ariya/phantomjs) (Development) Version (compile from `master` branch).
* [Supervisord](http://supervisord.org/) - or some way to manage the `goscribed` process.
* `Rasterize.coffee` - included in the `libs` directory.

## How to run

Make sure that the whole goscribe repo is checked out appropriated via `go get github.com/nickpresta/goscribe`

1. Build scribe: `go build -o scribe github.com/nickpresta/scribe/`
2. `./scribe -binary /path/to/phantomjs/binary -script /path/to/rasterize.coffee`

See `./scribe --help` for more information.

## Goscribe

This is a library which exposes an HTTP Handler, which will accept GET requests to generate a PDF.

## Documentation

## License

Scribe and Goscribed are released under the MIT license. See `LICENSE.md` for details.

![Powered by Gophers](http://i.imgur.com/SwkPj.png "Powered by Gophers")
