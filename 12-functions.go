package main

import (
	"fmt"
)

func main() {
	// Function without parameters
	sayHello()

	// Function with parameters
	greet("Alice", 30)

	// Function with return value
	sum := add(5, 10)
	fmt.Println("Sum:", sum)

	// Function as a variable
	multiply := func(a, b int) int {
		return a * b
	}
	fmt.Println("Product:", multiply(5, 10))

	// Anonymous function
	anonymousFunc := func() {
		fmt.Println("This is an anonymous function.")
	}
	anonymousFunc()

	// Function with multiple return values
	result1, result2 := func(x, y int) (int, int) {
		return x + y, x * y
	}(3, 4)
	fmt.Println("Sum:", result1, "Product:", result2)
}

func sayHello() {
	fmt.Println("Hello!")
}

func greet(name string, age int) {
	fmt.Printf("Hello %s, you are %d years old.\n", name, age)
}

func add(a, b int) int {
	return a + b
}
