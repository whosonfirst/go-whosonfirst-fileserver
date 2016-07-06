package main

import (
	"flag"
	"fmt"
	"github.com/whosonfirst/go-httpony/cors"
	"github.com/whosonfirst/go-httpony/sso"
	"github.com/whosonfirst/go-httpony/tls"
	"net/http"
	"os"
	"path/filepath"
)

func main() {

	var host = flag.String("host", "localhost", "Hostname to listen on")
	var port = flag.Int("port", 8080, "Port to listen on")
	var path = flag.String("path", "./", "Path served as document root.")
	var cors_enable = flag.Bool("cors", false, "Enable CORS headers")
	var cors_allow = flag.String("allow", "*", "Enable CORS headers from these origins")
	var tls_enable = flag.Bool("tls", false, "Serve requests over TLS") // because CA warnings in browsers...
	var tls_cert = flag.String("tls-cert", "", "Path to an existing TLS certificate. If absent a self-signed certificate will be generated.")
	var tls_key = flag.String("tls-key", "", "Path to an existing TLS key. If absent a self-signed key will be generated.")
	var sso_enable = flag.Bool("sso", false, "Enable OAuth2 single-sign-on (SSO) provider hooks")
	var sso_config = flag.String("sso-config", "", "The path to a valid SSO provider config file")

	flag.Parse()

	docroot, err := filepath.Abs(*path)

	if err != nil {
		panic(err)
	}

	endpoint := fmt.Sprintf("%s:%d", *host, *port)

	root := http.Dir(docroot)
	fs := http.FileServer(root)

	handlers := make([]http.Handler, 0)
	handlers = append(handlers, fs)

	if *sso_enable {

		sso_provider, err := sso.NewSSOProvider(*sso_config, endpoint, docroot, *tls_enable)

		if err != nil {
			panic(err)
			return
		}

		last_handler := handlers[len(handlers)-1]
		sso_handler := sso_provider.SSOHandler(last_handler)

		handlers = append(handlers, sso_handler)
	}

	last_handler := handlers[len(handlers)-1]
	handler := cors.EnsureCORSHandler(last_handler, *cors_enable, *cors_allow)

	if *tls_enable {

		var cert string
		var key string

		if *tls_cert == "" && *tls_key == "" {

			root, err := tls.EnsureTLSRoot()

			if err != nil {
				panic(err)
			}

			cert, key, err = tls.GenerateTLSCert(*host, root)

			if err != nil {
				panic(err)
			}

		} else {
			cert = *tls_cert
			key = *tls_key
		}

		fmt.Printf("start and listen for requests at https://%s\n", endpoint)
		err = http.ListenAndServeTLS(endpoint, cert, key, handler)

	} else {

		fmt.Printf("start and listen for requests at http://%s\n", endpoint)
		err = http.ListenAndServe(endpoint, handler)
	}

	if err != nil {
		panic(err)
	}

	os.Exit(0)
}
