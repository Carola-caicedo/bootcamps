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

// Exercise 2: Running a command and showing output
import (
	"log"
	"os"
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

	// Redirige the exit command amd print in the console
	cmd.Stdout = os.Stdout
	// Same as cmd.Stdout but for mensages error
	cmd.Stderr = os.Stderr
	//Run the command and wait for the end
	err := cmd.Run()
	if err != nil {
		// if err is not nil, log the error and exit the program
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
}
