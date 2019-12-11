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

	total := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		mass, err := strconv.Atoi(scanner.Text())
		CheckError(err)
		total += CalcMass(mass)
	}

	fmt.Printf("Total mass: %v\n", total)
}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func CalcMass(total int) (mass int) {
	mass = total / 3
	mass = mass - 2
	return
}
