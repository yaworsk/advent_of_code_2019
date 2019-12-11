package main

import "testing"

func TestCalcMass(t *testing.T) {
	want := [5]int{2, 2, 654, 33583, 20}
	masses := [5]int{12, 14, 1969, 100756, 4}
	for i := 0; i < 4; i++ {
		got := CalcMass(masses[i])
		if got != want[i] {
			t.Errorf("CalcMass(%d) = %d; want %d", masses[i], got, want[i])
		}
	}
}
