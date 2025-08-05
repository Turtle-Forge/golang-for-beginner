package main

import (
	"fmt"
)

func main() {
	var a = 15 + 25
	fmt.Println(a)

	var (
		sum1 = 100 + 50    // 150 (100 + 50)
		sum2 = sum1 + 250  // 400 (150 + 250)
		sum3 = sum2 + sum2 // 800 (400 + 400)
	)
	fmt.Println(sum3)

	var x = 10
	fmt.Println(x)

	x += 5
	fmt.Println(x)

	var y = 5
	var z = 3
	fmt.Println(y > z) // returns 1 (true) because 5 is greater than 3
}
