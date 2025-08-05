package main

import (
	"fmt"
)

func main() {
	day := 3
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid day")
	}

	// Switch with multiple cases
	grade := 'B'
	switch grade {
	case 'A', 'B':
		fmt.Println("Excellent")
	case 'C':
		fmt.Println("Good")
	case 'D':
		fmt.Println("Fair")
	case 'F':
		fmt.Println("Poor")
	default:
		fmt.Println("Invalid grade")
	}

}

// Switch statement example
// This program uses a switch statement to print the name of the day based on the number provided
