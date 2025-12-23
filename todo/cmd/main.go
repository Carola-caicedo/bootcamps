package main

import (
	"flag"
	"fmt"
	"os"

	// Import the todo package for made the functions
	"todo"
)

func main() {

	// Define flags
	listFlag := flag.Bool("list", false, "List all tasks")
	taskFlag := flag.String("task", "", "Add a new task")
	completefLag := flag.Int("complete", 0, "Mark task as complete by its number")
	deleteFalg := flag.Int("delete", 0, "Delete task by its number")

	flag.Parse()

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
			fmt.Fprintf(os.Stderr, "Error loanding tasks: %v\n:", err)
			os.Exit(1)
		}
	}

	// priority for flags
	if *listFlag {
		// List incomplete tasks
		incompleteCount := 0
		for _, item := range *l {
			if !item.Done {
				// Format the time
				completedAt := item.CompletedAt.Format("2006-01-02 15:04:05")
				createdAt := item.CreatedAt.Format("2006-01-02 15:04:05")
				// Message according to the assignment
				fmt.Printf("Title: %s, Done: %t, CreatedAt: %s, CompletedAt: %s\n", item.Task, item.Done, createdAt, completedAt)
				// Count the number of incomplete tasks
				incompleteCount++
			}
		}
		// If there are no incomplete tasks, show a message
		if incompleteCount == 0 {
			fmt.Println("No incomplete tasks.")
		}
		return
	}

	if *completefLag > 0 {
		// Complete the task
		index := *completefLag - 1
		if err := l.Complete(index); err != nil {
			fmt.Fprintf(os.Stderr, "Error completing task: %v\n:", err)
			os.Exit(1)
		}

		// save changes
		if err := l.Save(filename); err != nil {
			fmt.Fprintf(os.Stderr, "Error saving tasks: %v\n:", err)
			os.Exit(1)
		}
		return
	}

	if *deleteFalg > 0 {
		// Delete the task
		index := *deleteFalg - 1
		if err := l.Delete(index); err != nil {
			fmt.Fprintf(os.Stderr, "Error deleting task: %v\n:", err)
			os.Exit(1)
		}
		// save changes
		if err := l.Save(filename); err != nil {
			fmt.Fprintf(os.Stderr, "Error saving tasks: %v\n:", err)
			os.Exit(1)
		}
		return
	}

	if *taskFlag != "" {
		// Add the task
		l.Add(*taskFlag)

		// save changes
		if err := l.Save(filename); err != nil {
			fmt.Fprintf(os.Stderr, "Error saving tasks: %v\n:", err)
			os.Exit(1)
		}
		fmt.Println("Task added")
		return
	}

	// if pass a flag not valid
	fmt.Fprintf(os.Stderr, "Invalid flag\n, please use -list, -task, -complete or -delete\n")
	flag.Usage()
	os.Exit(1)
}
