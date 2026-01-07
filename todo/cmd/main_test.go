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

func cleanFile() error {
	os.Remove(fileName)
	return os.WriteFile(fileName, []byte{}, 0644)
}

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
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	cmdPath := filepath.Join(dir, binName)

	t.Run("AddNewTask", func(t *testing.T) {

		cleanFile()

		// use the flag -task for add a new task
		task := "New task"
		cmd := exec.Command(cmdPath, "-task", task)
		fmt.Println(cmd)
		err := cmd.Run()
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("ListTasks", func(t *testing.T) {
		cleanFile()

		// add a new task
		task := "New task"
		cmd := exec.Command(cmdPath, "-task", task)
		err := cmd.Run()
		if err != nil {
			t.Fatal(err)
		}

		// use the flag -list for list all tasks
		cmd = exec.Command(cmdPath, "-list")
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatal(err)
		}

		// Check for review the format (separately) converts to text
		output := string(out)

		expected := fmt.Sprintf("[ ] 0: %s", task)
		if !strings.Contains(output, expected) {
			t.Errorf("Output should contain '%s', got: %s", expected, output)
		}
	})

	t.Run("CompleteTask", func(t *testing.T) {
		cleanFile()

		// use the flag -complete for complete a task
		newTask := "Task to complete?"
		cmd := exec.Command(cmdPath, "-task", newTask)
		err := cmd.Run()
		if err != nil {
			t.Fatal(err)
		}

		// Complete the task (should be task number 1 since we added it)
		cmd = exec.Command(cmdPath, "-complete", "1")
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

		output := string(out)

		expected := fmt.Sprintf("[X] 0: %s", newTask)
		if !strings.Contains(output, expected) {
			t.Errorf("The task could be complete,expected %s, but got %s", expected, output)
		}

	})

	t.Run("DeleteTask", func(t *testing.T) {
		cleanFile()

		task2 := "Task to delete?"
		cmd := exec.Command(cmdPath, "-task", task2)
		err := cmd.Run()
		if err != nil {
			t.Fatal(err)
		}

		// Delete the task2 because it is incomplete
		cmd = exec.Command(cmdPath, "-delete", "1")
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

		output := string(out)

		// Should not show incomplete task
		if strings.Contains(output, task2) {
			t.Errorf("Should be deleted task, but got %s", output)
		}
	})

	t.Run("NoFlagsError", func(t *testing.T) {

		cleanFile()

		// Running without flags should produce an error
		cmd := exec.Command(cmdPath)
		err := cmd.Run()
		if err == nil {
			t.Error("Expected error when running without flags, but got none")
		}
	})
}
