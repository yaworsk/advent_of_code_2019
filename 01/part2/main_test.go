package main

import "testing"

func TestCalcFuel(t *testing.T) {
	masses := [5]int{12, 14, 1969, 100756}
	want := [5]int{2, 2, 966, 50346}
	for i := 0; i < 4; i++ {
		got := CalcFuel(masses[i])
		if got != want[i] {
			t.Errorf("CalcMass(%d) = %d; want %d", masses[i], got, want[i])
		}
	}
}
