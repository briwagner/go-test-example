package operator

import (
	"testing"
)

func TestPrints(t *testing.T) {
	got := Format("hello")
	want := "HELLO"
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestGreets(t *testing.T) {
	n := Normalizer{}

	got := n.Greet("Bitts")
	want := "Hello Bitts"
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}

	n.Formal = true
	got = n.Greet("Buttersworth")
	want = "Hello Mr Buttersworth"
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
