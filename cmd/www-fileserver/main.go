package main

import (
	"flag"
	"fmt"
	"github.com/NYTimes/gziphandler"
	"github.com/rs/cors"
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
	
	// var tls = flag.Bool("tls", false, "...")

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

	address := fmt.Sprintf("%s:%d", *host, *port)
	log.Printf("listening on %s\n", address)

	/*
	if *tls {

		// see also: https://github.com/Shyp/generate-tls-cert

		cert := "fixme"
		key := "fixme"

		err = http.ListenAndServeTLS(address, cert, key, mux)

	} else {
		err = http.ListenAndServe(address, mux)
	}
	*/
	
	err = http.ListenAndServe(address, mux)
	
	if err != nil {
		log.Fatal(err)
	}

	os.Exit(0)
}
