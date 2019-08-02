# go-whosonfirst-fileserver

There are many static file servers written in Go (and friends). This one is ours.

## Install

You will need to have both `Go` (specifically [version 1.12](https://golang.org/dl/) or higher because we're using [Go modules](https://github.com/golang/go/wiki/Modules)) and the `make` programs installed on your computer. Assuming you do just type:

```
make tools
```

All of this package's dependencies are bundled with the code in the `vendor` directory.

## Tools

### www-fileserver

```
> ./bin/www-fileserver -h
Usage of ./bin/www-fileserver:
  -cors
    	Enable CORS header on responses
  -cors-allowed-origins string
    	Comma-separated list of CORS allowed origins (default "*")
  -gzip
    	gzip response bodies
  -host string
    	Hostname to listen on (default "localhost")
  -path string
    	Path served as document root. (default "./")
  -port int
    	Port to listen on (default 8080)
```

## See also

* https://github.com/NYTimes/gziphandler
* https://github.com/rs/cors