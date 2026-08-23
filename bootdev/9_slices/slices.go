package slices

import (
	"fmt"
)

// Like structs, I think slices are big topic point in go.

/* The following are ARRAYS */
// - must be of the same type
// - FIXED Size upon delcaration [5]int, [10]int, etc.
func stringArrayInitialization() {
	// long way
	var someInts1 [5]int

	// short way
	someInts2 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	// auto count the size (still an array)
	autoCountSize := [...]int{100, 200, 300}

	fmt.Println("verbose initialization: ", someInts1)
	fmt.Println("inline initialization: ", someInts2)
	fmt.Println("autoCountSize initialization: ", autoCountSize)
	fmt.Println()
}

/* SLICES */
// - "dynamically-sized, flexible VIEW" of an array's contents
// 		Kind of like a "window" / "scope" into an array
// - A function that has access to a slice can modify the array that slice points to
//		Note that the above is not true for arrasys (not mutated by functions)

func how2CreateSlice() {
	// Create slice from an array
	intArray := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0} // array since specified size
	intSlice1 := intArray[:5]                         // A "view" of an array
	intSlice2 := intArray[5:]                         // A "view" of an array

	fmt.Println("intArray: ", intArray)
	fmt.Println("intSlice1: ", intSlice1)
	fmt.Println("intSlice2: ", intSlice2)
	fmt.Println()

	// Create slice inline
	// - Don't specify the array size
	inlineSlice := []int{} // not an array since not size was specified
	inlineSlice = append(inlineSlice, 1_000)
	fmt.Println("inlineSlice: ", inlineSlice)
	fmt.Println()

	// Create slice using make()
	// 	- 1st digit is length of elements
	// 	- 2nd digit is max capacity of elements
	makeSlice := make([]int, 7, 10)
	fmt.Println("makeSlice: ", makeSlice)
	fmt.Println()
}

func sliceLenAndCap() {
	makeSlice := make([]int, 7, 10)

	for i := range 5 {
		makeSlice = append(makeSlice, i*10)
	}

	// Go automatically adjusted the capacity to assign a larger block

	fmt.Println("makeSlice make([]int, 7, 10) contents: ", makeSlice)
	fmt.Println("makeSlice len: ", len(makeSlice), "cap: ", cap(makeSlice))
	fmt.Println()
}

func variadicFnAndSpreadOperator(ints ...int) (combinedValue int) {
	// I think variadic is just like spread in JS for unknown number of values
	combinedValue = 0

	fmt.Println("list of ints: ", ints)

	// Use range to return idx, element of a list (usually for a for-loop)
	for _, element := range ints {
		combinedValue += element
	}

	fmt.Println("list of ints combined value: ", combinedValue)
	fmt.Println()

	return combinedValue
}

func create2dArrays() {
	a := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	a[0][1] = 42
	fmt.Println(a)
}

// Again arrays are when size is specified.
func create2dSlices() {
	s := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	s[0][1] = 42
	fmt.Println(s)
}

/* Note on APPEND */
// Always make sure the first argument is the same slice you're modifying

func appendSliceIntoAnotherSlice() {
	slice1 := []int{1, 3, 5, 7, 9}
	slice2 := []int{2, 4, 6, 8, 0}

	sliceAppendedIntoSlice := []int{}

	sliceAppendedIntoSlice = append(sliceAppendedIntoSlice, slice1...)
	sliceAppendedIntoSlice = append(sliceAppendedIntoSlice, slice2...)

	fmt.Println("sliceAppendedIntoSlice: ", sliceAppendedIntoSlice)
	fmt.Println()
}

func SlicesMain() {
	fmt.Println("===== 9_slices =====")

	stringArrayInitialization()
	how2CreateSlice()
	sliceLenAndCap()
	variadicFnAndSpreadOperator(1, 2, 3, 4, 5, 6)
	appendSliceIntoAnotherSlice()

	fmt.Println("")
	create2dArrays()
	create2dSlices()
}
