package loops

import "fmt"

/* FOR-LOOPS only */
// - Go only has for loops to cover all the loops variants

func classicForLoop(upper int) {
	// You can the parts of the loop as needed
	// for INITIAL; CONDITION; AFTER
	for i := 0; i < upper; i++ {
		fmt.Println(i)
	}

	// The `continue` and `break` keywords are also available
}

func LoopsMain() {
	fmt.Println("===== 8_loops =====")

	classicForLoop(5)
}
