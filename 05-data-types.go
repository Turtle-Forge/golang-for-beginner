package main

import (
	"fmt"
)

func main() {

	// Basic
	var a bool = true    // Boolean
	var b int = 5        // Integer
	var c float32 = 3.14 // Floating point number
	var d string = "Hi!" // String

	fmt.Println("Boolean: ", a)
	fmt.Println("Integer: ", b)
	fmt.Println("Float:   ", c)
	fmt.Println("String:  ", d)

	// Boolean
	var b1 bool = true // typed declaration with initial value
	var b2 = true      // untyped declaration with initial value
	var b3 bool        // typed declaration without initial value
	b4 := true         // untyped declaration with initial value

	fmt.Println(b1) // Returns true
	fmt.Println(b2) // Returns true
	fmt.Println(b3) // Returns false
	fmt.Println(b4) // Returns true

	// Integer
	var i1 int = 500
	var i2 int = -4500
	fmt.Printf("Type: %T, value: %v", i1, i1)
	fmt.Printf("Type: %T, value: %v", i2, i2)

	// Float
	var f1 float32 = 123.78
	var f2 float32 = 3.4e+38
	fmt.Printf("Type: %T, value: %v\n", f1, f1)
	fmt.Printf("Type: %T, value: %v", f2, f2)

	var x float64 = 1.7e+308
	fmt.Printf("Type: %T, value: %v", x, x)

	// String
	var txt1 string = "Hello!"
	var txt2 string
	txt3 := "World 1"

	fmt.Printf("Type: %T, value: %v\n", txt1, txt1)
	fmt.Printf("Type: %T, value: %v\n", txt2, txt2)
	fmt.Printf("Type: %T, value: %v\n", txt3, txt3)
}
