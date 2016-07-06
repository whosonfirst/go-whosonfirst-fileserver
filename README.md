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

Single sign-on functionality allows your static website to act as a delegated authentication (specifically OAuth2) consumer of a different service and to use that authorization as a kind of persistent login for your own application.

When enabled three things happen:

1. Your website will "grow" a `/signin/` endpoint. When visited a user will be sent to the SSO provider's OAuth2 authenticate endpoint to confirm that they want to allow your website to perform actions on their behalf.
2. Your website will "grow" a `/auth/` endpoint. If a user approves your request to perform actions on their behalf they will be sent back to this endpoint and your website will complete the process to retrieve a persistent authentication token binding your website to the current user. That user's access token will be stored, encrypted, in a browser cookie whose expiration date will match the expiration date of the token itself.
3. If a valid token cookie is found then, on all other HTML pages, it will be inserted in to page's `body` element in a `data-api-access-token` attribute. Additionally, a `data-api-endpoint` attribute (as defined in the SSO config) will be added.

There are some important caveats as of this writing:

* It is left up your web application to determine what to _do_ with these new endpoint and functionality. This includes embedding or rendering links to the `/signin` endpoint.
* There is currently no way to overwrite, or manually set, the expiration date for token cookies. This is _not_ a feature and will be addressed soon.
* There is currently no `/signout` endpoint through which a user may terminate their session and delete their token cookie. Again, this will be addressed soon.

The details of registering your web application, as an OAuth2 consumer, with any given third-party are outside the scope of this document. At a minimum if you are using `wof-fileserver` to run a web application locally you should make sure that the third-party service supports redirecting users to `http://localhost`

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
