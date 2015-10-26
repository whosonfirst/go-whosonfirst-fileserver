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
	var cors = flag.Bool("cors", false, "Enable CORS headers")

	flag.Parse()

	docroot, err := filepath.Abs(*path)

	if err != nil {
		panic(err)
	}

	log.Printf("Static file server running at %s:%d. CTRL + C to shutdown\n", "http://localhost", *port)

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

	str_port := ":" + strconv.Itoa(*port)
	root := http.Dir(docroot)

	err = http.ListenAndServe(str_port, wof_handler(http.FileServer(root)))

	if err != nil {
		log.Fatal("Failed to start server, because %v\n", err)
	}
}
