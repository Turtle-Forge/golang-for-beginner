package main

import (
	"fmt"
)

func main() {
	age := 20
	if age >= 18 {
		fmt.Println("You are eligible to vote.")
	} else {
		fmt.Println("You are not eligible to vote.")
	}

	marks := 75
	if marks >= 60 {
		fmt.Println("You passed the exam.")
	} else {
		fmt.Println("You failed the exam.")
	}

	temperature := 30
	if temperature > 25 {
		fmt.Println("It's a warm day.")
	} else {
		fmt.Println("It's a cool day.")
	}

	// Nested if statement
	if age >= 18 {
		if marks >= 60 {
			fmt.Println("You are an adult and you passed the exam.")
		} else {
			fmt.Println("You are an adult but you failed the exam.")
		}
	} else {
		fmt.Println("You are not an adult.")
	}

	// else if statement
	if age < 13 {
		fmt.Println("You are a child.")
	} else if age < 20 {
		fmt.Println("You are a teenager.")
	} else {
		fmt.Println("You are an adult.")
	}
	fmt.Println("End of conditions example.")
}
