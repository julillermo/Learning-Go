package functions

import "fmt"

func acceptingParams(arg1 int, arg2 string) {
	fmt.Println("arg1:", arg1, " ", "arg2:", arg2)
}

func multipleParameterSameType(arg1, arg2 int, arg3 string) {
	fmt.Println("arg1:", arg1, " ", "arg2:", arg2, " ", "arg3:", arg3)
}

/* Passing by Value
- Unless the following datatypes, Go always passes by value into a function:
	- slices
	- pointers
	- maps
	- channels
	- intefaces
- This also means that Go functoins can't mutate if not any of the above datatypes
*/

func ignoreReturnValue() (string, string) {
	return "Hello", "World"
}

func namedReturnValues() (firstValue, secondValue string) {
	firstValue = "Hello"
	secondValue = "World"
	return
}

func explicitNamedReturnValues() (firstValue, secondValue string) {
	firstValue = "Hello"
	secondValue = "World"
	return firstValue, secondValue
}

/* Early returns
- Like other languages, Go can have an early return in a function
*/

func passingFunctionsAsParameters(
	str1,
	str2 string,
	combineFunc func(string, string) string,
) string {
	return combineFunc(str1, str2)
}

func combine2Strings(str1, str2 string) string {
	return str1 + str2
}

func usingAnonymousFunctions() {
	varUsingAnonymousFunc := passingFunctionsAsParameters(
		"Hello",
		"World",
		func(str1, str2 string) string {
			return str1 + str2
		})

	fmt.Println("varUsingAnonymousFunc:", varUsingAnonymousFunc)
}

func saveAnonymousFunctionToVariable() {
	varSaveAnonymousFunc := func(str1, str2 string) string {
		return str1 + str2
	}

	fmt.Println("varSaveAnonymousFunc:", varSaveAnonymousFunc("Hello", "World"))
}

/* Using Defers
- Like in HTML, you use use DEFER to delay the execution of a function
	until the surrounding function returns
- This is useful for closing files and database connections.
*/

/* SCOPES
- All variables are contained within a scope {}.
	They aren't accessible from outside
*/

func closuresConcept() func(int) int {
	// Returning a function the refers to a variable in the original scope
	// 	keeps that variable accessible from that function
	var sum int = 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func FunctionsMain() {
	fmt.Println("===== 4_Functions =====")

	acceptingParams(100, "Goodbye")
	multipleParameterSameType(100, 200, "nyello")
	_, secondValue := ignoreReturnValue()
	fmt.Println("ignored first value and only used secondValue:", secondValue)

	namedFirstValue, namedSecondValue := namedReturnValues()
	fmt.Println("namedFirstValue:", namedFirstValue, "namedSecondValue:", namedSecondValue)

	explicitNamedFirstValue, explicitNamedSecondValue := explicitNamedReturnValues()
	fmt.Println("explicitNamedFirstValue:", explicitNamedFirstValue, "explicitNamedSecondValue:", explicitNamedSecondValue)

	funcAsArg := passingFunctionsAsParameters("FristString", "SecondString", combine2Strings)
	fmt.Println("funcAsArg:", funcAsArg)

	usingAnonymousFunctions()
	saveAnonymousFunctionToVariable()

	closureAccessFn := closuresConcept()
	closureValue := closureAccessFn(1)
	closureValue = closureAccessFn(2)
	closureValue = closureAccessFn(3)
	fmt.Println("closureValue:", closureValue)
}
