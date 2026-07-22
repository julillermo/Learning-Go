package variables

import (
	"fmt"
)

/* BASIC VARIABLES
- int
- bool
- string
- float64
*/

/* TYPE SIZES
int int8 int16 int32 int64
uint uint8 uint16 uint32 uint64 uintptr
float32 float64
complex64 complex128
byte // alias for uint8
rune // alias for int32 (represents a Unicode code point)
*/

func variables_1() {
	// the sad way
	var value int
	value = 10
	fmt.Println("value is: ", value)
}

func variables_2() {
	var value int = 20
	fmt.Println("value is: ", value)
}

func variables_3() {
	// Short variable declaration
	// the type is inferred
	value := 30
	fmt.Println("value is: ", value)
}

func convertingBetweenTypes() {
	someFloat := 3.14
	someInt := int(someFloat) // Convert float64 to int
	fmt.Println("Converted value:", someInt)
}

func concatStrings() {
	str1 := "string #1"
	str2 := "string #2"
	concatString := str1 + " " + str2
	fmt.Println("Concatenated string:", concatString)
}

func sameLineDeclaration() {
	someNum, someString := 123, "ABC"
	fmt.Println("someNum:", someNum, "someString:", someString)
}

func Variables_main() {
	fmt.Println("===== 1_Variables =====")
	variables_1()
	variables_2()
	variables_3()
	convertingBetweenTypes()
	concatStrings()
}
