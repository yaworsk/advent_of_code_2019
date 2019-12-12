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

	fmt.Printf("Total mass: %v\n", total_fuel)
}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func CalcFuel(mass int) (total_fuel int) {
	remaining_mass := mass

	for remaining_mass > 0 {
		fuel := remaining_mass / 3
		fuel = fuel - 2
		if fuel > 0 {
			total_fuel += fuel
		}
		remaining_mass = fuel
	}
	return
}
