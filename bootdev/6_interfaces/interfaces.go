package interfaces

import "fmt"

/* How I understand it */
// - Interfaces are like "shapes" or "forms" that a struct may or may not
//		adhere to.
// - They like templates that YOU have to build toward piece by piece
// - Once a struct is of this "shape" or "form", it can become a sort of
//		type-check of sorts

type languageWithPackageManager interface {
	showPackageManager() string
}

type programmingLanguage struct {
	name         string
	typeStrategy string
	creator      string
}

type python struct {
	programmingLanguage
}

func (p python) showPackageManager() string {
	return "uv"
}

// := can only be used inside of functions
var pythonLang = python{
	programmingLanguage: programmingLanguage{
		name:         "python",
		creator:      "Guido",
		typeStrategy: "dynamic typing",
	},
}

type javaScript struct {
	programmingLanguage
}

func (js javaScript) showPackageManager() string {
	return "npm"
}

var jsLang = javaScript{
	programmingLanguage: programmingLanguage{
		name:         "javascript",
		creator:      "Brendan",
		typeStrategy: "dynamic typing",
	},
}

type rust struct {
	programmingLanguage
}

func (r rust) showPackageManager() string {
	return "cargo"
}

var rustLang = rust{
	programmingLanguage: programmingLanguage{
		name:         "rust",
		creator:      "The Rust Team",
		typeStrategy: "static typing",
	},
}

func triggerInterfaceMethod(langName string, pl languageWithPackageManager) {
	// An interface can only call methods
	// Hence, langName is passed separately
	pkgManager := pl.showPackageManager()
	printTxt := fmt.Sprintf("%s's package manager is: %s", langName, pkgManager)
	fmt.Println(printTxt)
}

func triggerInterfaceMethodTypeAssertion(pl languageWithPackageManager) {
	pkgManager := pl.showPackageManager()

	var printTxt string

	if p, ok := pl.(python); ok {
		printTxt = fmt.Sprintf("%s's package manager is: %s", p.name, pkgManager)
	}

	if js, ok := pl.(javaScript); ok {
		printTxt = fmt.Sprintf("%s's package manager is: %s", js.name, pkgManager)
	}

	if rust, ok := pl.(rust); ok {
		printTxt = fmt.Sprintf("%s's package manager is: %s", rust.name, pkgManager)
	}

	fmt.Println(printTxt)
}

func triggerInterfaceMethodTypeSwitch(pl languageWithPackageManager) {
	pkgManager := pl.showPackageManager()

	var printTxt string

	switch v := pl.(type) {
	case python:
		printTxt = fmt.Sprintf("%s's package manager is: %s", v.name, pkgManager)
	case javaScript:
		printTxt = fmt.Sprintf("%s's package manager is: %s", v.name, pkgManager)
	case rust:
		printTxt = fmt.Sprintf("%s's package manager is: %s", v.name, pkgManager)
	default:
		printTxt = fmt.Sprintf("The language's package manager is: %s", pkgManager)

	}
	fmt.Println(printTxt)
}

func InterfacesMain() {
	fmt.Println("===== 6_interfaces =====")

	triggerInterfaceMethod(pythonLang.name, pythonLang)
	triggerInterfaceMethod(jsLang.name, jsLang)
	triggerInterfaceMethod(rustLang.name, rustLang)

	fmt.Println("")
	triggerInterfaceMethodTypeAssertion(pythonLang)
	triggerInterfaceMethodTypeAssertion(jsLang)
	triggerInterfaceMethodTypeAssertion(rustLang)

	fmt.Println("")
	triggerInterfaceMethodTypeSwitch(pythonLang)
	triggerInterfaceMethodTypeSwitch(jsLang)
	triggerInterfaceMethodTypeSwitch(rustLang)

}
