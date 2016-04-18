# go-whosonfirst-fileserver

There are many static file servers written in Go (and friends). This one is ours.

## Usage

### Building

```
$> make build
```

### Running

```
$> bin/wof-fileserver -path /usr/local/mapzen/whosonfirst-data/data -port 9999
```

### CORS

Yes. Pass the `-cors` flag when starting up the server.

## Example

On the server side:

```
$> wof-fileserver -path /usr/local/mapzen/whosonfirst-data/data/ -port 9999 -cors
2015/10/26 11:35:10 Static file server running at http://localhost:9999. CTRL + C to shutdown
2015/10/26 11:35:16 [GET] /858/723/15/85872315.geojson
```

On the client side:

```
$> curl -v http://localhost:9999/858/723/15/85872315.geojson > /dev/null
* Connected to localhost (::1) port 9999 (#0)
> GET /858/723/15/85872315.geojson HTTP/1.1
> Host: localhost:9999
> User-Agent: curl/7.43.0
> Accept: */*
> 
< HTTP/1.1 200 OK
< Accept-Ranges: bytes
< Access-Control-Allow-Origin: *
< Content-Length: 8604
< Content-Type: text/plain; charset=utf-8
< Last-Modified: Mon, 05 Oct 2015 17:54:44 GMT
< Date: Mon, 26 Oct 2015 18:37:35 GMT
< 
{ [8604 bytes data]
```

## Does it do anything else? Tricks? Things you could talk about at a cocktail party?

No.

## See also

* https://github.com/cortesi/devd
* https://pauladamsmith.com/blog/2014/06/quickserver.html
