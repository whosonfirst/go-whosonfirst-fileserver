# go-whosonfirst-fileserver

There are many static file servers written in Go (and friends). This one is ours.

## Usage

### Building

```
$> make build
```

_See note below about installing [dependencies](#dependencies)._

### Running

```
$> bin/wof-fileserver -path /usr/local/mapzen/whosonfirst-data/data -port 9999
```

### CORS

Yes. Pass the `-cors` flag when starting up the server.

### Single sign-on (SSO)

Yes. Pass the `-sso` and `-sso-config PATH_TO_CONFIG_FILE` flags when starting up the server.

_Please write the long detailed version here._

#### SSO config files

_Example:_

```
[oauth]
client_id=OAUTH2_CLIENT_ID
client_secret=OAUTH2_CLIENT_SECRET
auth_url=https://example.com/oauth2/request/
token_url=https://example.com/oauth2/token/
api_url=https://example.com/api/
scopes=write

[www]
cookie_name=sso
cookie_secret=SSO_COOKIE_SECRET
```

SSO config files are standard `ini` style config files.

### TLS

Yes. Pass the `-tls` flag when startup up the server. If you have your own TLS key and certificate then you would specify them using the `-tls-key` and `-tls-cert` arguments respectively. If not then the server will generate a self-signed TLS key and certificate pair (which will make your browser complain so use this feature with the appropriate amount of caution and diligence).

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

## Dependencies

### Vendoring

Vendoring has been disabled for the time being because when trying to fetch some vendored dependencies goes pear-shape with errors like this:

```
make deps
# cd /Users/local/mapzen/mapzen-slippy-map/www-server/vendor/src/github.com/whosonfirst/go-httpony; git submodule update --init --recursive
fatal: no submodule mapping found in .gitmodules for path 'vendor/src/golang.org/x/net'
package github.com/whosonfirst/go-httpony: exit status 128
make: *** [deps] Error 1
```

I have no idea and would welcome suggestions. Also something something something about the way that the dependencies for `go-httpony` aren't pulled in automatically and need to be explicitly fetched in this package's [Makefile](Makefile).

## See also

* https://github.com/whosonfirst/go-httpony
* https://github.com/cortesi/devd
* https://pauladamsmith.com/blog/2014/06/quickserver.html
