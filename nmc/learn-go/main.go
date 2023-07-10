package main

import (
	"fmt"
	"strings"
)

func plus(a int, b int) int {
	return a + b
}

func multiply(a, b int) int {
	return a * b
}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

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

	// functions
	fmt.Println(plus(2, 2))
	fmt.Println(multiply(2, 2))

	// function returns 2 value
	totalLength1, upperName1 := lenAndUpper("ian")
	fmt.Println(totalLength1, upperName1)

	// "_" will ignore returned value
	totalLength2, _ := lenAndUpper("choi")
	fmt.Println(totalLength2)

	// function with multiple arguments
	repeatMe("hello", "world", "ian", "choi")
}
