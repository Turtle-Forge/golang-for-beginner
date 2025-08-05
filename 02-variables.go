package main

import "fmt"

func main() {
	// deklarasi variable
	var username string

	// inisiasi variabel dengan nilai
	username = "gbennnn"

	// cetak
	fmt.Println("Username: ", username)

	var a, b, c, d int = 1, 3, 5, 7

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	var aa, bb = 6, "Hello"
	cc, dd := 7, "World!"

	fmt.Println(aa)
	fmt.Println(bb)
	fmt.Println(cc)
	fmt.Println(dd)

	var (
		ab int
		bc int    = 1
		cd string = "hello"
	)

	fmt.Println(ab)
	fmt.Println(bc)
	fmt.Println(cd)
}
