package main_test

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

var (
	binName  = "todo"
	fileName = ".todo.json"
)

func TestMain(m *testing.M) {
	fmt.Println("Building tool...")

	if runtime.GOOS == "windows" {
		binName += ".exe"
	}

	build := exec.Command("go", "build", "-o", binName)
	err := build.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot build tool %s: %s", binName, err)
		os.Exit(1)
	}

	err = os.WriteFile(fileName, []byte{}, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot create file %s", fileName)
		os.Exit(1)
	}

	fmt.Println("Running tests....")
	result := m.Run()

	fmt.Println("Cleaning up....")
	os.Remove(binName)
	os.Remove(fileName)

	os.Exit(result)
}

func TestTodoCLI(t *testing.T) {
	task := "New Task"

	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	cmdPath := filepath.Join(dir, binName)

	t.Run("AddNewTask", func(t *testing.T) {
		// use the flag -task for add a new task
		cmd := exec.Command(cmdPath, "-task", task)
		fmt.Println(cmd)
		err := cmd.Run()
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("ListTasks", func(t *testing.T) {
		// use the flag -list for list all tasks
		cmd := exec.Command(cmdPath, "-list")
		fmt.Println(cmd)
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatal(err)
		}

		// Check for review the format (separately) converts to text
		output := string(out)
		// search "tittle" with in the message
		if !strings.Contains(output, "Title: "+task) {
			t.Errorf("Output should contain 'Title: %s', got: %s", task, output)
		}
		// search "done: false" with in the message
		if !strings.Contains(output, "Done: false") {
			t.Errorf("Output should contain 'Done: false', got: %s", output)
		}
		// search when create the task (date)
		if !strings.Contains(output, "CreatedAt: ") {
			t.Errorf("Output should contain 'CreatedAt: ', got: %s", output)
		}
		// search when complete the task (date)
		if !strings.Contains(output, "CompletedAt: ") {
			t.Errorf("Output should contain 'CompletedAt: ', got: %s", output)
		}
	})

	t.Run("CompleteTask", func(t *testing.T) {
		// use the flag -complete for complete a task
		newTask := "Task to complete?"
		cmd := exec.Command(cmdPath, "-task", newTask)
		err := cmd.Run()
		if err != nil {
			t.Fatal(err)
		}

		// Complete the task (should be task number 1 since we added it)
		cmd = exec.Command(cmdPath, "-complete", "1")
		fmt.Println(cmd)
		err = cmd.Run()
		if err != nil {
			t.Fatal(err)
		}

		// use the flag -list for list all tasks (should show the task as complete)
		cmd = exec.Command(cmdPath, "-list")
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatal(err)
		}

		// Should not show incomplete task
		notExpected := fmt.Sprintf("task: %s it's done: False ", task)
		if strings.Contains(string(out), notExpected) {
			t.Errorf("The task could be complete and Shouldn't show incomplete task (list), but got %s", string(out))
		}

	})

	t.Run("DeleteTask", func(t *testing.T) {
		task2 := "Task to delete?"
		cmd := exec.Command(cmdPath, "-task", task2)
		err := cmd.Run()
		if err != nil {
			t.Fatal(err)
		}

		// Delete the task2 because it is incomplete
		cmd = exec.Command(cmdPath, "-delete", "1")
		fmt.Println(cmd)
		err = cmd.Run()
		if err != nil {
			t.Fatal(err)
		}

		// use the flag -list for list all tasks (shouldn't show the task2)
		cmd = exec.Command(cmdPath, "-list")
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatal(err)
		}

		// Should not show incomplete task
		if strings.Contains(string(out), "task:") {
			t.Errorf("Shouldn't show incomplete task, bbut got %s", string(out))
		}
	})

	t.Run("NoFlagsError", func(t *testing.T) {
		// Running without flags should produce an error
		cmd := exec.Command(cmdPath)
		fmt.Println(cmd)
		err := cmd.Run()
		if err == nil {
			t.Error("Expected error when running without flags, but got none")
		}
	})
}
