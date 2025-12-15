package main

import (
	"log"
	"os/exec"
	"runtime"
)

//Excerise 1
//Func main
func main() {

	// Command ls (list files) and -lah (long, all, KB, MB, not bytes)
	cmd := exec.Command("ls", "-lah")
	//runtime.GOOS is the var of system the program (windows, linux...)( programs to run on any system)
	if runtime.GOOS == "windows" {
		// if the system is windows, run the command talklist (A show all)(Q show owner)
		cmd = exec.Command("cmd","/C", "dir", "/Q", "/A" )
	}

	// run the command
	err := cmd.Run()
	// if there an error back err
	if err != nil {
		// if err is not nil, log the error and exit the program
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
}