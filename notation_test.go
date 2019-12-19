package main

import "testing"

type NotationStrTest string

func (n NotationStrTest) String() string {
	return string(n)
}

func TestNotation(t *testing.T) {
	t.Run("print notation", func(t *testing.T) {
		p := printNotation(NotationStrTest("H"))
		if p != "   H" {
			t.Fatalf("got %q, want %q", p, "   H")
		}
	})
}
