package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	// port := flag.Int("p", 8080, "Port to server")
	port := flag.String("p", ":8080", "Port to server")
	host := flag.String("h", "localhost", "Host to server")
	datafile := flag.String("f", "datafile.json", "contains URL file as DB")
	flag.Parse()

	addr := *host + *port

	server := &http.Server{
		Addr:    addr,
		Handler: newMux(*datafile),
	}

	log.Fatal(server.ListenAndServe())
}
