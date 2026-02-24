package main

import (
	"net/http"
	"github.com/Carola-caicedo/todo"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	if r.URL.Path == "/" {
		errorReply(w, r, http.StatusNotFound, "404 page not found")
		return
	}

	textReply(w, r, http.StatusOK, "Hello World")
}

func getAllHandler (datafile string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		list := todo.List{}
		
		err := list.Get(datafile)
		if err != nil {
			errorReply(w, r, http.StatusInternalServerError, "Error reading error")
			return
		}

		response := &todoResponse{
			Results: list,
		}

		jsonReply(w, r, http.StatusOK, response)
	}
}
