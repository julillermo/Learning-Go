package conditionals

import "fmt"

/* NOTES
- Parenthesis for are optional
*/

func ifElse() {
	someBool := false

	if someBool {
		fmt.Println("someBool is true")
	} else {
		fmt.Println("someBool is false")
	}
}

func initialIfStatement() {
	if someBool := true; someBool {
		fmt.Println("someBool is true")
	} else {
		fmt.Println("someBool is false")
	}
}

func switchStatement() {
	// In Go, the break statement is always implied per case

	someSwitchValue := "ABC"
	switch someSwitchValue {
	case "A":
		fmt.Println("the value is A")
	case "B":
		fmt.Println("the value is B")
	case "C":
		fmt.Println("the value is C")
	default:
		fmt.Println("the value is not A, B, or C")
	}
}

func ConditionalsMain() {
	fmt.Println("===== 3_Conditionals =====")

	ifElse()
	initialIfStatement()
}
