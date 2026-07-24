package main

import (
	"fmt"

	variables "github.com/julillermo/Learning-Go/bootdev/1_variables"
	constantsandformatting "github.com/julillermo/Learning-Go/bootdev/2_constants_and_formatting"
	conditionals "github.com/julillermo/Learning-Go/bootdev/3_conditionals"
	functions "github.com/julillermo/Learning-Go/bootdev/4_functions"
	structs "github.com/julillermo/Learning-Go/bootdev/5_structs"
	interfaces "github.com/julillermo/Learning-Go/bootdev/6_interfaces"
	errors "github.com/julillermo/Learning-Go/bootdev/7_errors"
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

	fmt.Println("")
	interfaces.InterfacesMain()

	fmt.Println("")
	errors.ErrorsMain()
}
