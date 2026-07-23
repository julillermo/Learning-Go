package constantsandformatting

import "fmt"

/* NOTES
- Constants must be declared the explicit way instead of using ;= operator
*/

func declareConstant() {
	const someConstValue = "This is a constant value"
	fmt.Println("Constant value:", someConstValue)
}

func computedConstant() {
	const constA = "text A"
	const constB = "text B"
	const computedConst = constA + " " + constB
	fmt.Println("Computed constant value:", computedConst)
}

func stringTemplateFormatting() {
	// Use `Srpintf()` to generate a formated string
	// Follows the C/C++ approach of string formats

	// You can also just append values using "+"

	intValue := fmt.Sprintf("Integer value: %d", 42)
	fmt.Println(intValue)

	limitFloat := fmt.Sprintf("Limit float decimals to 2: %.2f", 3.14159)
	fmt.Println(limitFloat)
}

// Go stores string as a sequence of bytes.
// Go also has the `rune` type (alias for int32) which represents a Unicode code point.
// Go also uses UTF-8 encoding for strings, allowing use of emojis and JKC characters
// When trying to break up a string in Go, reach for the `rune` type
func runesAndString() {

}

func ConstantsFormatMain() {
	fmt.Println("===== 2_Constants_and_Formatting =====")

	declareConstant()
	computedConstant()
	stringTemplateFormatting()
}
