package main

import (
	"encoding/json"
	"time"
	"github.com/Carola-caicedo/todo"
)

type todoResponse struct {
	Results todo.List `json:"results"`
}

func (r *todoResponse) MarshalJSON() ([]byte, error) {
	response := struct {
		Results      todo.List `json:"results"`
		Date         time.Time `json:"date"`
		TotalResults int       `json:"total_results"`
	}{
		Results:      r.Results,
		Date:         time.Now(),
		TotalResults: len(r.Results),
	}
	return json.Marshal(response)
}
