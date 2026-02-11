package main

import (
	"log"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is home page."))
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is the about page."))
}

func blogHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is the blog page."))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is the hello page."))
}

func worldHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is the world page."))
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/about", aboutHandler)
	mux.HandleFunc("/blog", blogHandler)
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/world", worldHandler)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
