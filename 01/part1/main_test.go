package main

import "testing"

func TestCalcMass(t *testing.T) {
	masses := [5]int{12, 14, 1969, 100756, 4}
	want := [5]int{2, 2, 654, 33583, 20}
	for i := 0; i < 4; i++ {
		got := CalcFuel(masses[i])
		if got != want[i] {
			t.Errorf("CalcFuel(%d) = %d; want %d", masses[i], got, want[i])
		}
	}
}
