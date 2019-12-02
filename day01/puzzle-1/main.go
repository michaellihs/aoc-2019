package part_1

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
		totalMass += (mass / 3) - 2
	}

	return totalMass, nil
}