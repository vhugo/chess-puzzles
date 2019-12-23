package main

import "testing"

func TestNotation(t *testing.T) {
	t.Run("wrap notation", func(t *testing.T) {
		got := wrapNotation("1.Rxc7 e4 2.Qxc7 Nd2+ 3.Kc1 Nec4 4.Rd1 Nb3+ 5.Kb1 Qc1+ 6.Rxc1 Ncd2#")
		wanted := "1.Rxc7 e4 \n2.Qxc7 Nd2+ \n3.Kc1 Nec4 \n4.Rd1 Nb3+ \n5.Kb1 Qc1+ \n6.Rxc1 Ncd2# "
		if got != wanted {
			t.Fatalf("got %q, want %q", got, wanted)
		}
	})
}
