package main

import "testing"

func TestMain(t *testing.T) {
	//1,0,0,0,99 becomes 2,0,0,0,99 (1 + 1 = 2).
	//2,3,0,3,99 becomes 2,3,0,6,99 (3 * 2 = 6).
	//2,4,4,5,99,0 becomes 2,4,4,5,99,9801 (99 * 99 = 9801).
	//1,1,1,4,99,5,6,0,99 becomes 30,1,1,4,2,5,6,0,99.
	wanted := make([][]int, 4)
	wanted[0] = []int{2, 0, 0, 0, 99}
	wanted[1] = []int{2, 3, 0, 6, 99}
	wanted[2] = []int{2, 4, 4, 5, 99, 9801}
	wanted[3] = []int{30, 1, 1, 4, 2, 5, 6, 0, 99}

	input := make([][]int, 4)
	input[0] = []int{1, 0, 0, 0, 99}
	input[1] = []int{2, 3, 0, 3, 99}
	input[2] = []int{2, 4, 4, 5, 99, 0}
	input[3] = []int{1, 1, 1, 4, 99, 5, 6, 0, 99}

	for i, _ := range input {
		got := configureComputer(input[i])
		if len(wanted[i]) != len(got) {
			t.Errorf("configureComputer(%d) = %d; wanted %d", input[i], got, wanted[i])
		} else {
			for x, _ := range wanted[i] {
				if wanted[i][x] != got[x] {
					t.Errorf("configureComputer(%d) = %d; wanted %d", input[i], got, wanted[i])
				}
			}
		}
	}
}
