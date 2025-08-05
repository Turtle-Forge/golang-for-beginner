package main

/*
Go has three functions to output text:

Print()
Println()
Printf()
*/

import (
	"fmt"
)

func main() {
	var i, j string = "Hello", "World"

	fmt.Print(i)
	fmt.Print(j)
	fmt.Print(i, "\n", j)
	fmt.Print(i, "\n")
	fmt.Print(j, "\n")

	var ix string = "Hello"
	var jx int = 15

	fmt.Printf("i has value: %v and type: %T\n", ix, x)
	fmt.Printf("j has value: %v and type: %T", jx, jx)
}
