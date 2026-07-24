package errors

import (
	"errors"
	"fmt"
)

/* Errors */
// Errors are essentially an the following interface
type errorCustom interface {
	Error() string
}

// In Go, errors are values that get passed around, and
// you're forced to handle them

func canThrowError(throw bool) (string, error) {
	if throw {
		nilString := ""
		throwError := errors.New("intentionally thrown error")
		return nilString, throwError
	} else {
		correctString := "correctString value"
		return correctString, nil
	}
}

func handlingFnsThatThrow(boolValue bool) {
	strValue, err := canThrowError(boolValue)
	if err != nil {
		fmt.Println("error was handled")
	} else {
		fmt.Println("strValue Check: ", strValue)
	}
}

/* Custom Error struct */
// - You can implement your own error for a struct
type customError struct {
	errorName string
}

func (cd customError) Error() string {
	return "You've just triggerd the " + cd.errorName + " error"
}

func triggerCustomError() error {
	return customError{
		errorName: "Super Dangerous Error",
	}
}

func handleCustomError() {
	err := triggerCustomError()
	if err != nil {
		fmt.Println("err check: ", err)
	}
}

/* Panic() & Recover() */
// You can trigger a Panic() that bubbles up until caught by a Recover()
// Generally, don't use this!
// Not this discussed in this repo, just search online

func ErrorsMain() {
	fmt.Println("===== 7_errors =====")

	handlingFnsThatThrow(true)
	handlingFnsThatThrow(false)
	handleCustomError()
}
