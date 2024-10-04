package server

import (
	"strconv"
	"errors"
)

func Fooer(input int) string {
	isfoo := (input % 3) == 0

	if isfoo {
		return "Foo"
	}

	return strconv.Itoa(input)
}

func internalFooerMethod(input int) int {
	return input
}

func FooMightFail() (interface{}, error) {
	return nil, errors.New("Unknown error")
}

func FooWillPanic() (interface{}, error) {
	return nil, errors.New("Unknown error")
}

func FooWillHaveErrors() (interface{}, error) {
	return nil, errors.New("Unknown error")
}