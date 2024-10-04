package server

import (
	"testing"
)

func TestFooer(t *testing.T) {
	result := fooer(2)
	if result != "2" {
		t.Error("Incorrect response, expected 2, got ", result)
	}

}