package pointers

import "fmt"

/* NOTES */
// - A pointer is a variable storing an address in memory
// - Prepending a variable with `&` retrieves its address in memory
// - Prepending a pointer with `*` dereferences (points back to the value )

func ampersandAsteriskMemoryAddress() {
	someText := "poopy butt head"
	someTextPtr := &someText
	someTextDereference := *someTextPtr

	fmt.Println("variable value: ", someText)
	fmt.Println("variable address: ", someTextPtr)
	fmt.Println("variable dereference: ", someTextDereference)
}

func userPointersToPassByReference(
	strPtr *string, // expects an addressed to be passed in
) {
	// perform string manipulation direct to the location in memory
	*strPtr += "... text has been modified"
	fmt.Println("Value for the modified text inside pass by ref function: ", *strPtr)
}

func twoVarsPointSameLocInMemory() {
	// Here, both x and y poin to the same thing in memory
	// - x is a regular variable
	// - y is a "pointer" variable
	var x int = 50
	var y *int = &x
	*y = 100
}

func checkNilPointerB4Dereference(someStringPtr *string) {
	// Dereferencing a nil pointer resutls in a panic / crash
	if someStringPtr == nil {
		return
	}
	// function logic here ...
}

func nilPointerCrashApplication() {
	var someStringPtr *string = nil
	fmt.Println("dereference the nil ptr: ", *someStringPtr)
}

/* Practice for modifying pass by in functions */
// - It's generally better to use pointers and pass by reference if you
//		want a function to modify the values (though not functional approach).
//		This is best practice even if the value can be mutated (structs, maps)

/* Performance gains from using Pointers in GO */
// - For simple values, the compute cost of copying a value is negligible
// - The performance gain of using Pointers is mainly apparent in large datagst

func PointersMain() {
	fmt.Println("===== 11_pointers =====")

	ampersandAsteriskMemoryAddress()

	someStringValue := "mew mew mew"
	userPointersToPassByReference(&someStringValue) // note you can't simple pass the address string
	fmt.Println("Value for the modified text inside pass by ref function: ", someStringValue)

	twoVarsPointSameLocInMemory()
	checkNilPointerB4Dereference(&someStringValue)

	// nilPointerCrashApplication() // This will crash the application
}
