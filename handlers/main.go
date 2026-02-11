package main

import (
	"log"
	"net/http"
)

type worldHandler struct {
}

func (wh worldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello about"))
}

func main() {
	mux := http.NewServeMux()

	wh := worldHandler{}

	mux.Handle("/world", wh)

	mux.Handle("/about", http.HandlerFunc(aboutHandler))

	mux.HandleFunc("/help", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello help"))
	})

	log.Fatal(http.ListenAndServe(":8080", mux))
}
