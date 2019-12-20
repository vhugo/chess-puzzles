package main

import "testing"

func TestNotation(t *testing.T) {
	t.Run("wrap notation", func(t *testing.T) {
		got := wrapNotation(`
	1.aaaa bbbb 2.cccc ddddd 3.eeee fffff 4.gggg hhhhh`)

		wanted := `1.aaaa bbbbb 
2.cccc ddddd 
3.eeee fffff 
4.gggg hhhhh `

		if got != wanted {
			t.Fatalf("got %q, want %q", got, wanted)
		}
	})
}
