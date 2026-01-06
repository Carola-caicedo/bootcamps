package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

// item type private
type item struct {
	Task        string    `json:"task"`
	Done        bool      `json:"done"`
	CreatedAt   time.Time `json:"created_at"`
	CompletedAt time.Time `json:"completed_at"`
}

// List type private
type List []item

func (l *List) String() string {

	if len(*l) == 0 {
		return ""
	}

	var result strings.Builder

	for i, t := range *l {

		status := "[ ]"
		if t.Done {
			status = "[X]"
		}

		line := fmt.Sprintf("%s %d: %s", status, i, t.Task)
		result.WriteString(line)

		if i < len(*l)-1 {
			result.WriteString("\n")
		}

	}
	return result.String()
}

// Add function create a new task
// (*) modify the original list (not the copy)
// (Not error) always works
func (l *List) Add(task string) {
	t := item{
		Task: task,
		Done: false,
		//current time
		CreatedAt: time.Now(),
		//empty time
		CompletedAt: time.Time{},
	}

	//add the new task to the list
	*l = append(*l, t)
}

// complete function mark a task as completed
// (*) modify the original list (not the copy)
// (error) if the index is invalid
func (l *List) Complete(i int) error {
	ls := *l
	//check the task range in the index
	if i < 0 || i >= len(ls) {
		return errors.New("invalid index")
	}

	//mark the task as completed
	ls[i].Done = true
	ls[i].CompletedAt = time.Now()

	return nil
}

// Delete function remove a task from the list
// (*) modify the original list (not the copy)
// (error) if the index is invalid
func (l *List) Delete(i int) error {
	ls := *l
	//check the task range in the index
	if i < 0 || i >= len(ls) {
		return errors.New("invalid index")
	}

	//remove the task from the list
	*l = append(ls[:i], ls[i+1:]...)

	return nil
}

// Save method saves the list to a JSON file
// filename: name of the file to save
// Returns error if the save process fails
func (l *List) Save(filename string) error {
	// Marshal the list to JSON
	jsonData, err := json.Marshal(l)
	if err != nil {
		return err
	}

	// Write the data to file (JSON)
	err = os.WriteFile(filename, jsonData, 0644)
	if err != nil {
		return err
	}
	return nil
}

// Get method reads the list from a JSON file
// filename: name of the file to read
// Returns error if the read process fails
func (l *List) Get(filename string) error {
	// Read the file
	file, err := os.ReadFile(filename)
	if err != nil {
		// if file not found error
		if errors.Is(err, os.ErrNotExist) {
			return errors.New("file not exists/found")
		}

		return err
	}

	//Verify the file is not empty
	if len(file) == 0 {
		return errors.New("file is empty")
	}

	// Unmarshal (to read and receive) the JSON data
	err = json.Unmarshal(file, l)
	if err != nil {
		return err
	}
	return nil
}
