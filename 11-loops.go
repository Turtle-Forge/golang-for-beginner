package main

import (
	"fmt"
)

func main() {
	// For loop example
	for i := 1; i <= 5; i++ {
		fmt.Println("Iteration:", i)
	}

	// While loop example
	j := 1
	for j <= 5 {
		fmt.Println("Iteration:", j)
		j++
	}

	// Infinite loop example
	/*
		for {
			fmt.Println("This will run forever")
		}
	*/

	// Loop with break statement
	for k := 1; k <= 10; k++ {
		if k == 6 {
			fmt.Println("Breaking the loop at:", k)
			break
		}
		fmt.Println("Iteration:", k)
	}

	// Loop with continue statement
	for l := 1; l <= 10; l++ {
		if l%2 == 0 {
			fmt.Println("Skipping even number:", l)
			continue
		}
		fmt.Println("Odd number:", l)
	}

	// Nested loops example
	for m := 1; m <= 3; m++ {
		for n := 1; n <= 2; n++ {
			fmt.Println("Outer loop:", m, "Inner loop:", n)
		}
	}

	// Using range in a loop
	numbers := []int{1, 2, 3, 4, 5}
	for i, v := range numbers {
		fmt.Println("Index:", i, "Value:", v)
	}
}
