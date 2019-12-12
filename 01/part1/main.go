package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("mass.txt")
	CheckError(err)
	defer f.Close()

	total_fuel := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		mass, err := strconv.Atoi(scanner.Text())
		CheckError(err)
		total_fuel += CalcFuel(mass)
	}

	fmt.Printf("Total Fuel: %v\n", total_fuel)
}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func CalcFuel(mass int) (fuel int) {
	fuel = mass / 3
	fuel = fuel - 2
	return
}
