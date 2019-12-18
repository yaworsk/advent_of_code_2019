package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")
	CheckError(err)
	defer f.Close()

	var instructions []int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		instructions = append(instructions, i)
		CheckError(err)
	}

	fmt.Println(configureComputer(instructions))
}

func configureComputer(input []int) []int {
	result := make([]int, len(input))
	copy(result, input)
	x := 0

	for x < len(result) {
		if result[x] == 99 {
			break
		}

		pos1 := result[x+1]
		pos2 := result[x+2]
		pos3 := result[x+3]
		updateVal := 0

		if result[x] == 1 {
			updateVal = result[pos1] + result[pos2]
			result[pos3] = updateVal
		} else if result[x] == 2 {
			updateVal = result[pos1] * result[pos2]
			result[pos3] = updateVal
		}

		x += 4
	}

	return result
}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
