package todo

import (
	"testing"

	"os"
)

func TestSaveAndGet(t *testing.T) {

	// Create a temporary file empty
	// "" is the directory temporary
	// "todo_test*.json" is the name of the file
	tempFile, err := os.CreateTemp("", "todo_test*.json")
	// If the file cannot be created, returned an error
	if err != nil {
		// If the error is not nil, log the error and exit the program
		t.Errorf("Cannot create file temp: %v", err)
	}
	// Defer postpone the deletion of the file
	// os.Remove delete the temporary file
	defer os.Remove(tempFile.Name())

	// Create a new list and add tasks
	list1 := List{}
	list1.Add("Task 1")
	list1.Add("Task 2")

	// Save the list in the temporary file
	err = list1.Save(tempFile.Name())
	//
	if err != nil {
		// If the error is not nil, with fatal stop the test immediately
		t.Fatalf("Error saving file: %v", err)
	}

	// Create a new list and load de datafrom the temporary file
	list2 := List{}
	err = list2.Get(tempFile.Name())

	if err != nil {
		// If the error is not nil, with fatal stop the test immediately
		t.Fatalf("Error getting file: %v", err)

	}

	// Verify the length of the lists
	if len(list1) != len(list2) {
		// If the length is different, log the error and exit the program and concatenates the variables the condition
		t.Errorf("The lists have different lengths: %d and %d", len(list1), len(list2))
	}

	// Verify the tasks of the lists one by one
	for i := 0; i < len(list1); i++ {
		// If the task is different, log the error and exit the program and concatenates the variables the condition
		if list1[i].Task != list2[i].Task {
			t.Errorf("The task is: %d, saved: %s and loaded: %s", i, list1[i].Task, list2[i].Task)
			continue
		}

		// If the done is different, log the error and exit the program and concatenates the variables the condition
		if list1[i].Done != list2[i].Done {
			// If there an error, log the error and exit the program
			t.Errorf("The tasks are different, saved done: %v and loaded done: %v", list1[i].Done, list2[i].Done)
		}

	}

}

func TestGetEmptyFile(t *testing.T) {

	// Create a temporary file empty
	// "" is the directory temporary
	// "todo_empty_test_*.json" is the name of the file
	tempFile, err := os.CreateTemp("", "todo_empty_test_*.json")
	// If the file cannot be created, returned an error
	if err != nil {
		// If the error is not nil, log the error and exit the program
		t.Errorf("Cannot create file temp: %v", err)
	}

	// Defer postpone the deletion of the file
	// os.Remove delete the temporary file
	defer os.Remove(tempFile.Name())

	// Create a new list and load de datafrom the temporary file
	// Try to upload an empty file
	list := List{}
	// It should return an error for the empty file.
	err = list.Get(tempFile.Name())

	// If the error is nil, log the error and exit the program
	if err == nil {
		t.Error("The file is empty, but the error is nil")
	}

}
