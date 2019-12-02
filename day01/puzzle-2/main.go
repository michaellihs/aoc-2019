package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var totalMass int

func main() {
	fuel, err :=calcFuel("masses.txt")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println(fuel)
}

func calcFuel(file string) (int, error) {
	f, err := os.Open(file)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		mass, err := strconv.Atoi(line)
		if err != nil {
			return 0, err
		}
		totalMass += calcFuelForMass(mass, 0)
	}

	return totalMass, nil
}

func calcFuelForMass(mass, fuel int) int {
	currFuel := (mass / 3) - 2
	if currFuel <= 0 {
		return fuel
	}
	return calcFuelForMass(currFuel, fuel + currFuel)
}