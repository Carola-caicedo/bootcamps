package main

import (
	"flag"
	"log"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path == "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello World"))
}

func main() {
	port := flag.String("p", ":8080", "Port to server")
	flag.Parse()

	mux := http.NewServeMux()

	mux.HandleFunc("/", rootHandler)

	log.Fatal(http.ListenAndServe(*port, mux))
}
