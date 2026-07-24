package slices

import (
	"fmt"
)

/* The following are ARRAYS */
// - must be of the same type
// - FIXED Size upon delcaration

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
}

/* SLICES */
// - "dynamically-sized, flexible VIEW" of an array's contents
// - A function that has access to a slices can modify the array that slice points to
//		Note that the above is not true for arrasys (not mutated by functions)

func how2CreateSlice() {
	// Create slice from an array
	intArray := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	intSlice1 := intArray[:5]
	intSlice2 := intArray[5:]

	fmt.Println("intArray: ", intArray)
	fmt.Println("intSlice1: ", intSlice1)
	fmt.Println("intSlice2: ", intSlice2)

	// Create slice inline
	// - Don't specify the array size
	inlineSlice := []int{}
	inlineSlice = append(inlineSlice, 1_000)
	fmt.Println("inlineSlice: ", inlineSlice)

	// Create slice using make()
	makeSlice := make([]int, 7, 10)
	fmt.Println("makeSlice: ", makeSlice)
}

func sliceLenAndCap() {
	makeSlice := make([]int, 7, 10)

	for i := range 5 {
		makeSlice = append(makeSlice, i*10)
	}

	fmt.Println("makeSlice make([]int, 7, 10) contents: ", makeSlice)
	fmt.Println("makeSlice len: ", len(makeSlice), "cap: ", cap(makeSlice))
}

func variadicFnAndSpreadOperator(ints ...int) (combinedValue int) {
	combinedValue = 0

	fmt.Println("list of ints: ", ints)

	// Use range to return idx, element of a list (usually for a for-loop)
	for _, element := range ints {
		combinedValue += element
	}

	fmt.Println("list of ints combined value: ", combinedValue)

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
