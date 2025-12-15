package main

// import (
// 	"log"
// 	"os/exec"
// 	"runtime"
// )

// //Excerise 1
// //Func main
// func main() {

// 	// Command ls (list files) and -lah (long, all, KB, MB, not bytes)
// 	cmd := exec.Command("ls", "-lah")
// 	//runtime.GOOS is the var of system the program (windows, linux...)( programs to run on any system)
// 	if runtime.GOOS == "windows" {
// 		// if the system is windows, run the command talklist (A show all)(Q show owner)
// 		cmd = exec.Command("cmd","/C", "dir", "/Q", "/A" )
// 	}

// 	// run the command
// 	err := cmd.Run()
// 	// if there an error back err
// 	if err != nil {
// 		// if err is not nil, log the error and exit the program
// 		log.Fatalf("cmd.Run() failed with %s\n", err)
// 	}
// }

// // Exercise 2: Running a command and showing output
// import (
// 	"log"
// 	"os"
// 	"os/exec"
// 	"runtime"
// )

// func main() {
// 	// Command ls (list files) and -lah (long, all, KB, MB, not bytes)
// 	cmd := exec.Command("ls", "-lah")
// 	//runtime.GOOS is the var of system the program (windows, linux...)( programs to run on any system)
// 	if runtime.GOOS == "windows" {
// 		// if the system is windows, run the command talklist (A show all)(Q show owner)
// 		cmd = exec.Command("cmd", "/C", "dir", "/Q", "/A")
// 	}

// 	// Redirige the exit command amd print in the console
// 	cmd.Stdout = os.Stdout
// 	// Same as cmd.Stdout but for mensages error
// 	cmd.Stderr = os.Stderr
// 	//Run the command and wait for the end
// 	err := cmd.Run()
// 	if err != nil {
// 		// if err is not nil, log the error and exit the program
// 		log.Fatalf("cmd.Run() failed with %s\n", err)
// 	}
// }

// // Exercise 3: Running a command and capturing the output
// import (
// 	"fmt"
// 	"log"
// 	"os/exec"
// 	"runtime"
// )

// func main() {
// 	// Command ls (list files) and -lah (long, all, KB, MB, not bytes)
// 	cmd := exec.Command("ls", "-lah")
// 	//runtime.GOOS is the var of system the program (windows, linux...)( programs to run on any system)
// 	if runtime.GOOS == "windows" {
// 		// if the system is windows, run the command "talklist" change for dir (A show all)(Q show owner)
// 		cmd = exec.Command("cmd", "/C", "dir", "/Q", "/A")
// 	}

// 	// Capture in out the output anr error in the same time (1 byte)
// 	out, err := cmd.CombinedOutput()
// 	if err != nil {
// 		// if err is not nil, log the error and exit the program
// 		log.Fatalf("cmd.Run() failed with %s\n", err)
// 	}

// 	//Show the output
// 	fmt.Printf("combined out: \n%s\n", string(out))
// }

// //Exercise 4: Behind the scenes of CombinedOutput()

// func (c *Cmd) CombinedOutput() ([]byte, error) {
// 	if c.Stdout != nil {
// 		return nil, errors.New("exec: Stdout already set")
// 	}
// 	if c.Stderr != nil {
// 		return nil, errors.New("exec: Stderr already set")
// 	}
// 	var b bytes.Buffer
// 	c.Stdout = &b
// 	c.Stderr = &b
// 	err := c.Run()
// 	return b.Bytes(), err
// }

// Exercise 5
import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"runtime"
)

func main() {
	// Command ls (list files) and -lah (long, all, KB, MB, not bytes)
	cmd := exec.Command("ls", "-lah")
	//runtime.GOOS is the var of system the program (windows, linux...)( programs to run on any system)
	if runtime.GOOS == "windows" {
		// if the system is windows, run the command talklist (A show all)(Q show owner)
		cmd = exec.Command("cmd", "/C", "dir", "/Q", "/A")
	}

	// Vars for capture the output
	// bytes.Buffer: var for binary data
	var stdout, stderr bytes.Buffer
	// Assign the output normal
	cmd.Stdout = &stdout
	// Assign the output error
	cmd.Stderr = &stderr
	//Run the command and wait for the end
	err := cmd.Run()
	if err != nil {
		// if err is not nil, log the error and exit the program
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}

	// Convert the bytes.Buffer to string and show the output
	outStr, errStr := stdout.String(), stderr.String()
	fmt.Printf("out:\n%s\n err:\n%s\n", outStr, errStr)
}
