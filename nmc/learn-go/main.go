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

func checkNumber3(num int) bool {
	switch {
	case num == 18:
		return true
	case num > 18:
		return true
	default:
		return false
	}
}

func checkNumber4(num int) bool {
	switch newNum := num + 2; newNum {
	case 17:
		return false
	case 19:
		return true
	default:
		return false
	}
}

// pointer
func pointer() {
	a1 := 2             // a1 => value 2
	b1 := a1            // b1 => value 2
	fmt.Println(a1, b1) // 2, 2
	a1 = 10             // a1 => 10
	fmt.Println(a1, b1) // 10, 2

	//  & => get variable's "point"
	a2 := 2               // a2 => value 2
	b2 := 10              // b2 => value 10
	fmt.Println(&a2, &b2) // 0x140000180f8, 0x14000018110
	a3 := 2               // a3 => value 2
	b3 := &a3             // b3 => a3's point
	fmt.Println(&a3, b3)  // 0x14000018118, 0x14000018118

	// * => get pointer's "value"
	a4 := 2              // a4 => value 2
	b4 := &a4            // b4 => a4's point
	fmt.Println(a4, *b4) // 2, 2 => *b4 = b4's (a4's) value

	a5 := 2              // a5 => value 2
	b5 := &a5            // b5 => a5's value
	fmt.Println(a5, *b5) // 2, 2
	a5 = 10              // change a5's value => b5's value also changed
	fmt.Println(a5, *b5) // 5, 5 => b5 changed to value 5

	a6 := 2              // a6 => value 2
	b6 := &a6            // b6 => a6's point
	*b6 = 10             // change b6's value => a6's value also changed
	fmt.Println(a6, *b6) // 10, 10
}

func arraysAndSlices() {
	// static array (array)
	array1 := [5]string{"hello", "world"}
	fmt.Println(array1)

	array1[1] = ","
	array1[2] = "_"
	array1[3] = "world"
	array1[4] = "!"
	// array1[5] = "" => invalid argument: index 5 out of bounds [0:5]
	fmt.Println(array1)

	// dynamic array (slice)
	array2 := []string{"hello"}
	array2 = append(array2, ",", "_", "world", "!")
	fmt.Println(array2)
}

func main() {
	fmt.Println("Hello, World!")

	// constants
	fmt.Println("constants ============> start")
	const const1 = "const1"
	fmt.Println(const1)
	const const2 string = "const2"
	fmt.Println(const2)
	fmt.Println("constants ============> end")

	// variables
	fmt.Println("variables ============> start")
	var1 := "var1"
	fmt.Println(var1)
	var1 = "changedVar1"
	fmt.Println(var1)

	var var2 string = "var2"
	fmt.Println(var2)
	var2 = "changedVar2"
	fmt.Println(var2)
	fmt.Println("variables ============> end")

	// functions
	fmt.Println("functions ============> start")
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
	fmt.Println("functions ============> end")

	// defer
	fmt.Println("defer ============> start")
	totalLength4, upperName4 := lenAndUpper3("hello ian choi")
	fmt.Println(totalLength4, upperName4)
	fmt.Println("defer ============> end")

	// Loop - for
	fmt.Println("for ============> start")
	fmt.Println(adds1(0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	fmt.Println(adds2(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11))
	fmt.Println("for ============> end")

	// Conditional Statement - if
	fmt.Println("if ============> start")
	fmt.Println(checkNumber1(17))
	fmt.Println(checkNumber2(17))
	fmt.Println("if ============> end")

	// switch
	fmt.Println("switch ============> start")
	fmt.Println(checkNumber3(17))
	fmt.Println(checkNumber4(17))
	fmt.Println("switch ============> end")

	// pointer
	fmt.Println("pointer ============> start")
	pointer()
	fmt.Println("pointer ============> end")

	// Arrays and slices
	fmt.Println("arrays and slices ============> start")
	arraysAndSlices()
	fmt.Println("arrays and slices ============> end")
}
