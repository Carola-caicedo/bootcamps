package fl

import "os"

type Flag struct {
	name        string // name of the flag
	value       *bool  // pointer to the boolean value
	description string // description of the flag
}

var flags []Flag // define value to control here all program flags

func Parse() {
	args := os.Args[1:]
	// Evaluate if the defined flags should change their default value

	for i := range flags {

		found := false

		// Evaluate if the flag is a boolean flag
		for _, arg := range args {
			if arg == flags[i].name {
				found = true
				break
			}
		}

		// if founf, set the flag to true
		if found {
			// Evaluate if the flag is a boolean flag
			*flags[i].value = true
		}
	}
}

func Bool(cmd string, value bool, description string) *bool {
	// Add logic to create boolean flags

	// Create a new flag
	flagValue := value

	// Create struct for the flag
	newFlag := Flag{
		name:        cmd,
		value:       &flagValue, // pointer to obtain the memory address
		description: description,
	}

	//append the flag to the list of flags
	flags = append(flags, newFlag)

	// return the pointer to the flag
	return &flagValue
}
