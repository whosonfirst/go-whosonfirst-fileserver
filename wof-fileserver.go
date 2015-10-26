package main

import (
	"flag"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
)

func main() {

	var port = flag.Int("port", 8080, "Port to listen")
	var path = flag.String("path", "./", "Path served as document root.")
	// var cors = flag.Bool("cors", false, "Enable CORS headers")

	flag.Parse()

	docroot, err := filepath.Abs(*path)

	if err != nil {
		panic(err)
	}

	log.Printf("Static file server running at %s:%d. CTRL + C to shutdown\n", "http://localhost", *port)

	// TO DO: wrap up CORS stuff...
	handler := http.FileServer(http.Dir(docroot))

	str_port := ":" + strconv.Itoa(*port)

	err = http.ListenAndServe(str_port, handler)

	if err != nil {
		log.Fatal("Failed to start server, because %v\n", err)
	}
}
