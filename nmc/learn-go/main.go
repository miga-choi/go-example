package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")

	// constants
	const const1 = "const1"
	fmt.Println(const1)
	const const2 string = "const2"
	fmt.Println(const2)

	// variables
	var1 := "var1"
	fmt.Println(var1)
	var1 = "changedVar1"
	fmt.Println(var1)

	var var2 string = "var2"
	fmt.Println(var2)
	var2 = "changedVar2"
	fmt.Println(var2)
}
