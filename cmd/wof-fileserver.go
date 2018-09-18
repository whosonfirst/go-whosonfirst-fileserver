package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {

	var host = flag.String("host", "localhost", "Hostname to listen on")
	var port = flag.Int("port", 8080, "Port to listen on")
	var path = flag.String("path", "./", "Path served as document root.")
	var tls = flag.Bool("tls", false, "...")

	flag.Parse()

	docroot, err := filepath.Abs(*path)

	if err != nil {
		log.Fatal(err)
	}

	root := http.Dir(docroot)
	fs_handler := http.FileServer(root)

	mux := http.NewServeMux()
	mux.Handle("/", fs_handler)

	address := fmt.Sprintf("%s:%d", *host, *port)
	log.Printf("listening on %s\n", address)

	if *tls {

		// see also: https://github.com/Shyp/generate-tls-cert
		
		cert := "fixme"
		key := "fixme"

		err = http.ListenAndServeTLS(address, cert, key, mux)
		
	} else {
		err = http.ListenAndServe(address, mux)
	}

	if err != nil {
		log.Fatal(err)
	}

	os.Exit(0)
}
