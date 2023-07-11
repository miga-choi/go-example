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

func lenAndUpper1(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func lenAndUpper2(name string) (length int, uppercase string) {
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func lenAndUpper3(name string) (length int, uppercase string) {
	defer fmt.Println("I'm done")
	length, uppercase = len(name), strings.ToUpper(name)
	return
}

func adds1(numbers ...int) int {
	result := 0
	for i := 0; i < len(numbers); i++ {
		result += numbers[i]
	}
	return result
}

func adds2(numbers ...int) (result int) {
	// for index, value := range [] {}
	for _, number := range numbers {
		result += number
	}
	return
}

func checkNumber1(num int) bool {
	if num < 18 {
		return false
	} else {
		return true
	}
}

func checkNumber2(num int) bool {
	if newNum := num + 2; newNum < 18 {
		return false
	}
	return true
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
	totalLength1, upperName1 := lenAndUpper1("ian")
	fmt.Println(totalLength1, upperName1)

	// "_" will ignore returned value
	totalLength2, _ := lenAndUpper1("choi")
	fmt.Println(totalLength2)

	// function with multiple arguments
	repeatMe("hello", "world", "ian", "choi")

	// declare return value as variables
	totalLength3, upperName3 := lenAndUpper2("ian choi")
	fmt.Println(totalLength3, upperName3)

	// defer
	totalLength4, upperName4 := lenAndUpper3("hello ian choi")
	fmt.Println(totalLength4, upperName4)

	// Loop - for
	fmt.Println(adds1(0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	fmt.Println(adds2(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11))

	// Conditional Statement - if
	fmt.Println(checkNumber1(17))
	fmt.Println(checkNumber2(17))
}
