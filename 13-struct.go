package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	// Creating an instance of the struct
	alice := Person{Name: "Alice", Age: 30}
	fmt.Println("Name:", alice.Name, "Age:", alice.Age)

	// Modifying struct fields
	alice.Age = 31
	fmt.Println("Updated Age:", alice.Age)

	// Using a struct in a function
	printPersonInfo(alice)
}

func printPersonInfo(p Person) {
	fmt.Println("Person Information:")
	fmt.Println("Name:", p.Name)
	fmt.Println("Age:", p.Age)
}
