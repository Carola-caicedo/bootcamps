package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	//Ontain user data
	minSrt, maxSrt, valuesStr := getInput()


	//convert to minimum
	min, err := strconv.ParseFloat(minSrt, 64)
	if err != nil {
		fmt.Println("Error: Invalid minimum value")
		return
	}


	// convert to maximum
	max, err := strconv.ParseFloat(maxSrt, 64)
	if err != nil {
		fmt.Println("Error: Invalid maximum value")
		return
	}

	// Verify that minimum value is less than maximum value
	if min >= max {
		fmt.Println("Error: Minimum value must be less than maximum value")
		return
	}

	// Convert values to float64
	var values []float64
	for _, str := range strings.Fields(valuesStr) {
		if num, err := strconv.ParseFloat(str, 64); err == nil {
			values = append(values, num)
		}
	}

	// Filter values within the range
	result := minmax(min, max, values...)

	// Show the result
	fmt.Printf("List of values within the range (%v, %v): %v\n", min, max, result)
}

func minmax(min, max float64, values ...float64) []float64 {
	var insideRange []float64
	for _, value := range values {
		if value >= min && value <= max {
			insideRange = append(insideRange, value)
		}
	}
	return insideRange
}

func getInput() (string, string, string) {
	scanner := bufio.NewScanner(os.Stdin)

	// Ask for minimum value
	fmt.Print("Enter minimum values: ")
	scanner.Scan()
	minSrt := strings.TrimSpace(scanner.Text())

	// Ask for maximum value
	fmt.Print("Enter maximum values: ")
	scanner.Scan()
	maxSrt := strings.TrimSpace(scanner.Text())

		// Obtain list of values
	fmt.Print("Enter values (separated by spaces): ")
	scanner.Scan()
	valuesStr := strings.TrimSpace(scanner.Text())

	return minSrt, maxSrt, valuesStr
}


