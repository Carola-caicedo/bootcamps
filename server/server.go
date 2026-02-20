package main

import (
	"encoding/json"
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

func jsonReply(w http.ResponseWriter, _ *http.Request, status int, payload *todoResponse) {
	w.WriteHeader(status)

	data, err := json.Marshal(payload)
	if err != nil {
		errorReply(w, nil, http.StatusInternalServerError, "Error generated Response JSON")
		return
		
	}

	w.Write(data)
}
