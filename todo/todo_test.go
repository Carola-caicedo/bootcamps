package todo

import (
	"testing"
)

func TestAdd(t *testing.T) {

	// Create a var tipe List
	var l List

	// var task is the text of the task
	task := "study for the exam"
	//Add the task to the list
	l.Add(task)

	// Verify the length of the list
	if len(l) != 1 {
		// If the length is different, log the error and exit the program and concatenate the variable the condition, and show in the message
		t.Errorf("Expected The list should have 1 item, got %d", len(l))
	}

	// Verify the text is the same of the task
	if l[0].Task != task {
		// If the text is different, log the error and exit the program and concatenate the variable the condition, and show in the message
		t.Errorf("Expected The task should be %s, got %s", task, l[0].Task)
	}

	// Verify the task is not done (false)
	if l[0].Done != false {
		// If the task is done, log the error and exit the program and concatenate the variable the condition, and show in the message
		t.Error("Expected The task should be not done")
	}

	// Verify the time the created the task completed is not zero
	if l[0].CreatedAt.IsZero() {
		// If the time is zero, log the error and exit the program
		t.Error("The metod CreateAt should not be zero")
	}

	// Verify the time the completed the task completed is zero
	if !l[0].CompletedAt.IsZero() {
		// If the time is zero, log the error and exit the program
		t.Error("The metod CompletedAt should be zero")
	}

	// var task2 is the text of the task
	// var task3 is the text of the task
	task2 := "study for the exam"
	task3 := "complete assignment"
	//Add the task to the list
	l.Add(task2)
	l.Add(task3)

	// Verify the length of the list
	if len(l) != 3 {
		// If the length is different, log the error and exit the program and concatenate the variable the condition, and show in the message
		t.Errorf("Expected The list should have 3 items, got %d", len(l))
	}

	// Verify the text is the same of the task
	if l[1].Task != task2 {
		// If the text is different, log the error and exit the program and concatenates the variables the condition, and show in the message
		t.Errorf("Expected The task 2 should be %s, got %s", task2, l[1].Task)
	}
	// Verify the text is the same of the task
	if l[2].Task != task3 {
		// If the text is different, log the error and exit the program and concatenates the variables the condition, and show in the message
		t.Errorf("Expected The task 3 should be %s, got %s", task3, l[2].Task)
	}

	// Verify the tasks with the same text are the same
	if l[0].Task != l[1].Task {
		// If the tasks with the same text are different, log the error and exit the program
		t.Error("The tasks with the same text should be same")
	}
}

func TestComplete(t *testing.T) {
	// Create a var tipe List
	var l List

	// add the tasks to the list (already created before)
	l.Add("task1")
	l.Add("task2")

	// Complete the task
	err := l.Complete(0)
	if err != nil {
		// If the error is not nil, log the error and exit the program
		t.Errorf("Error for complete: %v", err)
	}

	// Verify the task is complete (true)
	if !l[0].Done {
		// If the task is not complete, log the error and exit the program
		t.Error("Expected the task as complete")
	}

	// Verify the time the completed the task completed is not zero
	if l[0].CompletedAt.IsZero() {
		// If the time is zero, log the error and exit the program
		t.Error("CompletedAt should be zero after finished")
	}

	// Try complete the task again
	err = l.Complete(0)
	if err != nil {
		// If the error is not nil, log the error and exit the program
		t.Errorf("Error for complete, again? %v", err)
	}

}

func TestDelete(t *testing.T) {
	// Create a var tipe List
	var l List

	// Add tasks to the list
	l.Add("task1")
	l.Add("task2")
	l.Add("task3")

	// Delete the second task
	err := l.Delete(1)
	if err != nil {
		// If the error is not nil, log the error and exit the program
		t.Errorf("Error for deleting: %v", err)
	}

	// Verify the length of the list
	if len(l) != 2 {
		// If the length is different, log the error and exit the program and concatenate the variable the condition, and show in the message
		t.Errorf("Should have 2 tasks, got %d", len(l))
	}

	// Verify deleted the correct task
	if l[0].Task != "task1" {
		t.Errorf("First task should be 'task1', got '%s'", l[0].Task)
	}
	// Verify that task in the indice 1 is the correct task (befores task 3)
	if l[1].Task != "task3" {
		t.Errorf("Second task should be 'task3', got '%s'", l[1].Task)
	}

	// Try delete the invalid index
	err = l.Delete(5)
	// should be return an error
	if err == nil {
		// If the error is nil, log the error and exit the program
		t.Error("Should return an error with invalid index")
	}

}
