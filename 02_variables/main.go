package main

import "fmt"

func main() {

	// this variable's type cannot be infered so it must be defined
	var myAge int32 = 21

	// this one can, so bool can be omitted
	const isCool = true

	name, email := "Luca", "lucaisamazing@gmail.com"

	fmt.Println(name, myAge, email, isCool)

	// this prints the variable type
	fmt.Printf("%T\n", name)
}
