package todo

import (
	"errors"
	"time"
)

//item type private
type item struct {
	Task (string)
	Done (bool)
	CreatedAt (time.Time)
	CompletedAt (time.Time)
}

//List type private
type List []item


//Add function create a new task
//(*) modify the original list (not the copy)
//(Not error) always works 
func (l *List) Add(task string) {
	t := item{
		Task: task,
		Done: false,
		CreatedAt: time.Now(),//current time
		CompletedAt: time.Time{},//empty time
	}

	//append the new task to the list
	*l = append(*l, t)
}


//complete function mark a task as completed
//(*) modify the original list (not the copy)
//(error) if the index is invalid
func (l *List) Complete(i int) error {
	ls := *l
	//check the task range in the index
	if i < 0 || i >= len(ls) {
		return errors.New("invalid index")
	}

	//set the task as completed
	ls[i].Done = true
	ls[i].CompletedAt = time.Now()

	return nil
}


//Delete function remove a task from the list
//(*) modify the original list (not the copy)
//(error) if the index is invalid
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