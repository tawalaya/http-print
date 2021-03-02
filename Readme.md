# http-print

This is a simple local http server that prints every request to console. Its similar to `python3 -m http.server` but also excepts POST, PUT, HEAD etc.

## Getting Started

If you want to use it for debugging, development. 
Just get it using go `go get github.com/tawalaya/http-print` and run it with `$GOPATH/bin/http-print`. You can also download binary's form the [release tab](https://github.com/tawalaya/http-print/releases).

### Usage 
You can modify the behavior using the following command line flags:
- `--port` : the port this service is listening on, default 80
- `--status` : the http status code that is returned on each request, default 200
- `--response` : filepath to use as the response payload, by default the response is empty
- `--raw` : raw response string, take care of escaping only used when set, `--response` has priority
- `--type` : the content type used to responsed to requests, default is `"application/json"`
- `--verbose` : activate verbose logging

### Prerequisites

- Go Lang 1.10+


## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/tawalaya/http-print/tags). 

## Authors

* **Sebastian Werner** - *Initial work* - [tawalaya](https://github.com/tawalaya)

See also the list of [contributors](https://github.com/tawalaya/http-print/contributors) who participated in this project.

## License

This project is licensed under the Apache License 2.0 - see the [LICENSE.md](LICENSE.md) file for details
