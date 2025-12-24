package todo

import (
	"testing"
)

func TestAdd(t *testing.T) {
	var l List

	// var task is the text of the task
	task := "study for the exam"
	l.Add(task)

	if len(l) != 1 {
		t.Errorf("Expected The list should have 1 item, got %d", len(l))
	}

	// Verify the text is the same of the task
	if l[0].Task != task {
		t.Errorf("Expected The task should be %s, got %s", task, l[0].Task)
	}

	// Verify the task is not done (false)
	if l[0].Done != false {
		// If the task is done, log the error and exit the program and concatenate the variable the condition, and show in the message
		t.Error("Expected The task should be not done")
	}

	// Verify the time the created the task completed is not zero
	if l[0].CreatedAt.IsZero() {
		t.Error("The metod CreateAt should not be zero")
	}

	// Verify the time the completed the task completed is zero
	if !l[0].CompletedAt.IsZero() {
		t.Error("The metod CompletedAt should be zero")
	}

	task2 := "study for the exam"
	task3 := "complete assignment"
	l.Add(task2)
	l.Add(task3)

	if len(l) != 3 {
		t.Errorf("Expected The list should have 3 items, got %d", len(l))
	}

	// Verify the text is the same of the task
	if l[1].Task != task2 {
		t.Errorf("Expected The task 2 should be %s, got %s", task2, l[1].Task)
	}
	if l[2].Task != task3 {
		// If the text is different, log the error and exit the program and concatenates the variables the condition, and show in the message
		t.Errorf("Expected The task 3 should be %s, got %s", task3, l[2].Task)
	}

	if l[0].Task != l[1].Task {
		t.Error("The tasks with the same text should be same")
	}
}

func TestComplete(t *testing.T) {
	var l List

	// add the tasks to the list (already created before)
	l.Add("task1")
	l.Add("task2")

	err := l.Complete(0)
	if err != nil {
		t.Errorf("Error for complete: %v", err)
	}

	// Verify the task is complete (true)
	if !l[0].Done {
		t.Error("Expected the task as complete")
	}

	// Verify the time the completed the task completed is not zero
	if l[0].CompletedAt.IsZero() {
		t.Error("CompletedAt should be zero after finished")
	}

	err = l.Complete(0)
	if err != nil {
		t.Errorf("Error for complete, again? %v", err)
	}

}

func TestDelete(t *testing.T) {
	var l List

	l.Add("task1")
	l.Add("task2")
	l.Add("task3")

	err := l.Delete(1)
	if err != nil {
		t.Errorf("Error for deleting: %v", err)
	}

	if len(l) != 2 {
		// If the length is different, log the error and exit the program and concatenate the variable the condition
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
		t.Error("Should return an error with invalid index")
	}

}
