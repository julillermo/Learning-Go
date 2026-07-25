package main

import (
	"fmt"

	maps "github.com/julillermo/Learning-Go/bootdev/10_maps"
	pointers "github.com/julillermo/Learning-Go/bootdev/11_pointers"
	channels "github.com/julillermo/Learning-Go/bootdev/13_channels"
	variables "github.com/julillermo/Learning-Go/bootdev/1_variables"
	constantsAndFormatting "github.com/julillermo/Learning-Go/bootdev/2_constants_and_formatting"
	conditionals "github.com/julillermo/Learning-Go/bootdev/3_conditionals"
	functions "github.com/julillermo/Learning-Go/bootdev/4_functions"
	structs "github.com/julillermo/Learning-Go/bootdev/5_structs"
	interfaces "github.com/julillermo/Learning-Go/bootdev/6_interfaces"
	errors "github.com/julillermo/Learning-Go/bootdev/7_errors"
	loops "github.com/julillermo/Learning-Go/bootdev/8_loops"
	slices "github.com/julillermo/Learning-Go/bootdev/9_slices"
)

func main() {
	fmt.Println("Hello, World!")

	fmt.Println("")
	variables.VariablesMain()

	fmt.Println("")
	constantsAndFormatting.ConstantsFormatMain()

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

	fmt.Println("")
	loops.LoopsMain()

	fmt.Println("")
	slices.SlicesMain()

	fmt.Println("")
	maps.MapsMain()

	fmt.Println("")
	pointers.PointersMain()

	fmt.Println("")
	channels.ChannelsMain()
}
