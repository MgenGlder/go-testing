package server

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"fmt"
	"os"
)

var examplePackageVariable int

func TestHandler(t *testing.T) {
	s := &Server{}
	server := httptest.NewServer(http.HandlerFunc(s.HelloWorldHandler))
	defer server.Close()
	resp, err := http.Get(server.URL)

	if err != nil {
		t.Fatalf("error making request to server. Err: %v", err)
	}

	defer resp.Body.Close()

	// Assertions
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", resp.Status)
	}

	expected := "{\"message\":\"Hello World\"}"
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		t.Fatalf("error reading response body. Err: %v", err)
	}

	if expected != string(body) {
		t.Errorf("expected response body to be %v; got %v", expected, string(body))
	}
}

func TestMain(m *testing.M) {
	fmt.Println("Running examplePackageVariable scaffolding")
	examplePackageVariable = 2
	exitCode := m.Run()

	os.Exit(exitCode)
}

func FuzzStringCompare2(f *testing.F) {
	f.Add("hello", "world")
	f.Add("go", "golang")

	f.Fuzz(func(t *testing.T, a, b string) {
		result := strings.Compare(a, b)
		if result == 0 && a != b {
			t.Errorf("strings.Compare(%q, %q) returned 0, but strings are not equal", a, b)
		}
	})
}

func FuzzStringCompare(f *testing.F) {
	f.Add("hello", "world")
	f.Add("go", "golang")

	f.Fuzz(func(t *testing.T, a, b string) {
		result := strings.Compare(a, b)
		if result == 0 && a != b {
			t.Errorf("strings.Compare(%q, %q) returned 0, but strings are not equal", a, b)
		}
	})
}
