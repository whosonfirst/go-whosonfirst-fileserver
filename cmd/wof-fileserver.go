package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
)

func main() {

	var host = flag.String("host", "localhost", "Hostname to listen on")
	var port = flag.Int("port", 8080, "Port to listen on")
	var path = flag.String("path", "./", "Path served as document root.")
	var cors = flag.Bool("cors", false, "Enable CORS headers")

	flag.Parse()

	docroot, err := filepath.Abs(*path)

	if err != nil {
		panic(err)
	}

	log.Printf("Static file server (%s) running at %s:%d. CTRL + C to shutdown\n", docroot, *host, *port)

	wof_handler := func(next http.Handler) http.Handler {

		fn := func(w http.ResponseWriter, r *http.Request) {

			log.Printf("[%s] %s\n", r.Method, r.URL)

			if *cors {
				w.Header().Set("Access-Control-Allow-Origin", "*")
			}

			next.ServeHTTP(w, r)
		}

		return http.HandlerFunc(fn)
	}

	endpoint := fmt.Sprintf("%s:%d", *host, *port)
	root := http.Dir(docroot)

	err = http.ListenAndServe(endpoint, wof_handler(http.FileServer(root)))

	if err != nil {
		log.Fatal("Failed to start server, because %v\n", err)
	}
}
