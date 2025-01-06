package grammar

import (
	"testing"
)

func TestHelloWorld(t *testing.T) {
	got := HelloWorld()
	want := "Hello, World"

	if got != want {
		t.Errorf("HelloWorld() = %v, want %v", got, want)
	}
}
