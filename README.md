# Scribe

![Powered by Gophers](http://i.imgur.com/SwkPj.png "Powered by Gophers")

## Requirements

* [PhantomJS 1.9.0](https://github.com/ariya/phantomjs) (Development) Version (compile from `master` branch).
* [Supervisord](http://supervisord.org/) - or some way to manage the `goscribed` process.
* `Rasterize.coffee` - included in the `libs` directory.

## Goscribe

This is a library which exposes an HTTP Handler, which will accept GET requests to generate a PDF.
You do not need this unless you want to use it.

## Goscribed

This is an application which uses Goscribe to accept HTTP requests.
You can run this to start converting PDFs.

## Documentation

See `goscribe/README.md` and `goscribed/README.md` for more information specific to each package.

## License

Goscribe and Goscribed are released under the MIT license. See `LICENSE.md` for details.
