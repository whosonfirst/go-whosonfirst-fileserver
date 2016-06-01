package main

import (
	"flag"
	"fmt"
	"github.com/whosonfirst/go-httpony/tls"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {

	var host = flag.String("host", "localhost", "Hostname to listen on")
	var port = flag.Int("port", 8080, "Port to listen on")
	var path = flag.String("path", "./", "Path served as document root.")
	var cors = flag.Bool("cors", false, "Enable CORS headers")
	var tls = flag.Bool("tls", false, "Serve requests over TLS") // because CA warnings in browsers...
	var tls_cert = flag.String("tls-cert", "", "Path to an existing TLS certificate. If absent a self-signed certificate will be generated.")
	var tls_key = flag.String("tls-key", "", "Path to an existing TLS key. If absent a self-signed key will be generated.")

	flag.Parse()

	docroot, err := filepath.Abs(*path)

	if err != nil {
		panic(err)
	}

	log.Printf("Static file server (%s) running at %s:%d. CTRL + C to shutdown\n", docroot, *host, *port)

	wof_handler := func(next http.Handler) http.Handler {

		fn := func(rsp http.ResponseWriter, req *http.Request) {

			log.Printf("[%s] %s\n", req.Method, req.URL)

			if *cors {
				rsp.Header().Set("Access-Control-Allow-Origin", "*")
			}

			/*
			if req.URL.Path == "/foo" {

				rsp.Write(body)
			   	return 		      
			}
			*/
			
			next.ServeHTTP(rsp, req)
		}

		return http.HandlerFunc(fn)
	}


	endpoint := fmt.Sprintf("%s:%d", *host, *port)

	root := http.Dir(docroot)
	handler := wof_handler(http.FileServer(root))
	
	if *tls {

		var cert string
		var key string

		if *tls_cert == "" && *tls_key == "" {

		   	root, err := httpony.EnsureTLSRoot()

			if err != nil {
				panic(err)
			}
			
			cert, key, err = httpony.GenerateTLSCert(*host, root)
			
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
