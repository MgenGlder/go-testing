package server

import "strconv"

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