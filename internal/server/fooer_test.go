package server

import (
	"testing"
)

func TestFooer(t *testing.T) {
	result := Fooer(2)
	if result != "2" {
		t.Error("Incorrect response, expected 2, got ", result)
	}
}

func Test_internalFooer(t *testing.T) {
	result := internalFooerMethod(2)
	if result != 2 {
		t.Error("Incorrect response, expected 2, got", result)
	}
}