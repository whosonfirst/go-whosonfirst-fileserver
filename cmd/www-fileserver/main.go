package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"github.com/NYTimes/gziphandler"
	"github.com/rs/cors"
	"golang.org/x/crypto/acme/autocert"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	var host = flag.String("host", "localhost", "Hostname to listen on")
	var port = flag.Int("port", 8080, "Port to listen on")
	var path = flag.String("path", "./", "Path served as document root.")
	var enable_gzip = flag.Bool("gzip", false, "gzip response bodies")
	var enable_cors = flag.Bool("cors", false, "Enable CORS header on responses")
	var cors_origins = flag.String("cors-allowed-origins", "*", "Comma-separated list of CORS allowed origins")

	var enable_tls = flag.Bool("tls", false, "...")

	flag.Parse()

	docroot, err := filepath.Abs(*path)

	if err != nil {
		log.Fatal(err)
	}

	root := http.Dir(docroot)
	fs_handler := http.FileServer(root)

	if *enable_gzip {
		fs_handler = gziphandler.GzipHandler(fs_handler)
	}

	if *enable_cors {

		c := cors.New(cors.Options{
			AllowedOrigins: strings.Split(*cors_origins, ","),
		})

		fs_handler = c.Handler(fs_handler)
	}

	mux := http.NewServeMux()
	mux.Handle("/", fs_handler)

	if *enable_tls {

		// https://godoc.org/golang.org/x/crypto/acme/autocert
		// https://letsencrypt.org/docs/certificates-for-localhost/
		// https://github.com/FiloSottile/mkcert
		// https://blog.filippo.io/mkcert-valid-https-certificates-for-localhost/
		// https://github.com/golang/go/issues/20640

		/*

		> make tools ; sudo ./bin/www-fileserver -tls -host aa.local -path ./
		go build -mod vendor -o bin/www-fileserver cmd/www-fileserver/main.go
		2019/08/02 18:02:42 listening on aa.local:443
		2019/08/02 18:02:58 NAME aa.local
		2019/08/02 18:03:03 http: TLS handshake error from 127.0.0.1:57254: 400 urn:acme:error:malformed: Error creating new authz :: Name does not end in a public suffix

		*/
		
		*port = 443
		
		address := fmt.Sprintf("%s:%d", *host, *port)
		log.Printf("listening on %s\n", address)
		
		m := autocert.Manager{
			Prompt:     autocert.AcceptTOS,
			HostPolicy: autocert.HostWhitelist(*host),
			Cache:      autocert.DirCache("./certs"),
		}

		s := &http.Server{
			Addr:      ":https",
			Handler: mux, 
			TLSConfig: &tls.Config{
				GetCertificate: m.GetCertificate,
				MinVersion:               tls.VersionTLS12,
				CipherSuites: []uint16{
					tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
					tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
				},
				PreferServerCipherSuites: true,
			},
		}

		err = s.ListenAndServeTLS("", "")

	} else {

		address := fmt.Sprintf("%s:%d", *host, *port)
		log.Printf("listening on %s\n", address)
		
		err = http.ListenAndServe(address, mux)
	}

	// err = http.ListenAndServe(address, mux)

	if err != nil {
		log.Fatal(err)
	}

	os.Exit(0)
}
