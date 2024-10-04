package server

import (
	"testing"
	"fmt"
	"os"
)

func TestCleanup(t *testing.T) {
	f, _ := os.Create("tempfile")

	// Do some testing and asserting

	t.Cleanup(func() {
		os.Remove(f.Name())
	})
}

func TestTableTest(t *testing.T) {
	data := []struct {
		nameOfTest string
		firstName string
		expectedLastName string
		expectedFavColor string
		errorMsg string
	}{
		{"Test last name", "Kunle", "Oshiyoye", "green", "invalid last name"},
		{"Test color", "John", "Smith", "red", "invalid color"},
	}
	for _, d := range data {
		t.Run(d.nameOfTest, func(t *testing.T) {
			lastName, favColor := PersonDetails(d.firstName)
			if lastName != d.expectedLastName || favColor != d.expectedFavColor {
				t.Error(d.errorMsg)
			}
		})
	}
}


func TestFatalError(t *testing.T) {
	// If FooMightFail fails, the following code will panic.
	result, err := FooMightFail()

	if err != nil {
		t.Fatal("FooMightFail returned an unexpected error")
	}

	if result == nil {
		t.Fatal("FooMightFail didn't return a result")
	}

	result, err = FooWillPanic()

	if err != nil {
		t.Error("FooWillPanic returned an unexpected error")
	}

	if result == nil {
		t.Error("FooWillPanic didn't return a result")
	}
}

func TestRegularError(t *testing.T) {
	// Entire test is run, and all errors encountered are displayed.
	result, err := FooWillHaveErrors()

	if err != nil {
		t.Error("FooWillPanic returned an unexpected error")
	}

	if result == nil {
		t.Error("FooWillPanic didn't return a result")
	}
}

var packageVariable string

func TestMain(m *testing.M) {
	fmt.Println("Running examplePackageVariable scaffolding")
	packageVariable = "Some useful variable"
	exitCode := m.Run()

	os.Exit(exitCode)
}

func TestFooer(t *testing.T) {
	fmt.Println(packageVariable)
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