package main

import (
	"log"
	"net/http"
)

func newMux() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", rootHandler)

	return mux
}

func textReply(w http.ResponseWriter, _ *http.Request, status int, payload string) {
	w.WriteHeader(status)
	w.Write([]byte(payload))
}

func errorReply(w http.ResponseWriter, _ *http.Request, status int, payload string) {
	log.Printf("Error 500: open database.json: no such file or directory")

	http.Error(w, payload, status)
}
