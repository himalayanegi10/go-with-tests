package dependency_injection

import (
	"testing"
	"bytes"
	"os"
)

func TestDependency(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Chris")

	got := buffer.String()
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got , want)
	}
}

func TestStdoutDependency(t *testing.T) {
	Greet(os.Stdout, "Chris")
}