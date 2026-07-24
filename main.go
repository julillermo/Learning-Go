package main

import (
	"fmt"

	variables "github.com/julillermo/Learning-Go/bootdev/1_variables"
	constantsandformatting "github.com/julillermo/Learning-Go/bootdev/2_constants_and_formatting"
	conditionals "github.com/julillermo/Learning-Go/bootdev/3_Conditionals"
	functions "github.com/julillermo/Learning-Go/bootdev/4_functions"
	structs "github.com/julillermo/Learning-Go/bootdev/5_Structs"
)

func main() {
	fmt.Println("Hello, World!")

	fmt.Println("")
	variables.VariablesMain()

	fmt.Println("")
	constantsandformatting.ConstantsFormatMain()

	fmt.Println("")
	conditionals.ConditionalsMain()

	fmt.Println("")
	functions.FunctionsMain()

	fmt.Println("")
	structs.StructsMain()
}
