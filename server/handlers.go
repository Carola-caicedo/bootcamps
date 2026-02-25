package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

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

func getAllHandler(w http.ResponseWriter, r *http.Request, list *todo.List) {
	// err := list.Get(datafile)
	// if err != nil {
	// 	errorReply(w, r, http.StatusInternalServerError, "Error reading error")
	// 	return
	// }

	response := &todoResponse{
		Results: *list,
	}

	jsonReply(w, r, http.StatusOK, response)
}

type NewTask struct {
	Task string `json:"task"`
}

func addHandler(w http.ResponseWriter, r *http.Request, list *todo.List, datafile string) {
	var item NewTask

	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		errorReply(w, r, http.StatusBadRequest, "Invalid format JSON")
		return
	}

	list.Add(item.Task)

	err = list.Save(datafile)
	if err != nil {
		errorReply(w, r, http.StatusInternalServerError, "Error saving todo list")
		return
	}

	textReply(w, r, http.StatusCreated, "Success creating task")
}

func getOneHandler(w http.ResponseWriter, r *http.Request, list *todo.List, id int) {
	item := (*list)[id]
	oneTask := todo.List{item}

	response := &todoResponse{
		Results: oneTask,
	}
	jsonReply(w, r, http.StatusOK, response)
}

func validateID(idStr string, list *todo.List) (int, error) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, err
	}

	if id < 0 {
		return 0, strconv.ErrRange
	}

	if id >= len(*list) {
		return 0, strconv.ErrRange
	}
	return id, nil
}

func router(datafile string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		list := todo.List{}

		err := list.Get(datafile)
		if err != nil {
			errorReply(w, r, http.StatusInternalServerError, "Error reading todo list")
			return
		}

		path := r.URL.Path

		if strings.HasPrefix(path, "/todo/") && path != "/todo/" {
			idStr := strings.TrimPrefix(path, "/todo/")

			id, err := validateID(idStr, &list)
			if err != nil {
				errorReply(w, r, http.StatusNotFound, "ID not found")
				return
			}

			switch r.Method {
			case http.MethodGet:
				getOneHandler(w, r, &list, id)
			default:
				errorReply(w, r, http.StatusMethodNotAllowed, "Method not allowed")
			}
			return
		}

		if path == "/todo" || path == "/todo/" {

			switch r.Method {
			case http.MethodGet:
				getAllHandler(w, r, &list)
			case http.MethodPost:
				addHandler(w, r, &list, datafile)
			default:
				errorReply(w, r, http.StatusMethodNotAllowed, "Method not allowed")
			}
			return
		}

		errorReply(w, r, http.StatusNotFound, "404 page not found")

	}
}
