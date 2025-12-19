package main

import (
	"fmt"
	"os"
	"strings"

	// Import the todo package for manage the functions
	"todo"
)

func main() {
	// Create a new list empty
	// & work only the original list
	l := &todo.List{}

	// File for save the tasks
	filename := ".todo.json"

	// Load existing tasks
	// If the file not exists or empty, start with empty list
	err := l.Get(filename)
	if err != nil {
		// If error is different to "file not exists" or "file is empty" unexpected error
		errMsg := err.Error()
		if errMsg != "file not exists/found" && errMsg != "file is empty" {
			// Unexpected error
			os.Exit(1)
		}
	}

	// If passed arguments
	if len(os.Args) > 1 {
		// Add all the words in a only snetence except the first (todo)
		task := strings.Join(os.Args[1:], " ")

		// Add the task
		l.Add(task)

		// Save the list in the file
		l.Save(filename)
		return
	}

	// If no arguments: List tasks
	for _, item := range *l {
		// Print task one by one
		fmt.Println(item.Task)
	}
}
