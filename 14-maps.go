package main

import "fmt"

func main() {
	// Creating a map
	ages := make(map[string]int)

	// Adding key-value pairs
	ages["Alice"] = 30
	ages["Bob"] = 25
	ages["Charlie"] = 35

	// Accessing values
	fmt.Println("Alice's age:", ages["Alice"])

	// Looping through a map
	for name, age := range ages {
		fmt.Printf("%s is %d years old.\n", name, age)
	}

	// Deleting a key
	delete(ages, "Bob")
	fmt.Println("After deletion:", ages)

	// Checking if a key exists
	if age, exists := ages["Charlie"]; exists {
		fmt.Println("Charlie's age:", age)
	} else {
		fmt.Println("Charlie not found.")
	}

	// Length of the map
	fmt.Println("Number of entries in ages map:", len(ages))
	// Clearing the map
	ages = make(map[string]int)
	fmt.Println("Ages map cleared:", ages)
	fmt.Println("End of maps example.")
}
