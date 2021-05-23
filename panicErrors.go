package main

import (
	"errors"
	"fmt"
	"log"
	"runtime/debug"
)

type ErrorWithTrace struct {
	text  string
	trace string
}

func New(text string) error {
	return &ErrorWithTrace{
		text:  text,
		trace: string(debug.Stack()),
	}
}

func (e *ErrorWithTrace) Error() string {
	return fmt.Sprintf("error: %s\ntrace:\n%s", e.text, e.trace)
}

func recoveryFunction() {
	fmt.Println("recovery")
	if err := recover(); err != nil {
		var newError error
		newError = errors.New("This is my new error with trace")
		log.Println("panic occurred:", err)
	}
}

func executePanic() {
	defer recoveryFunction()
	panic("Panic error")
	fmt.Println("Function finished to execute")
}

func main() {
	executePanic()
	fmt.Println("Main function block finished to execute")
}
