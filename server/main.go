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
	flag.Parse()

	addr := *host + *port

	server := &http.Server{
		Addr:    addr,
		Handler: newMux(),
	}

	log.Fatal(server.ListenAndServe())
}
