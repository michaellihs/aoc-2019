package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Orbit struct {
	Name          string
	Parent        string
	IndirectCount int
}

func main() {
	orbitMap, err := ReadInput("/Users/mimi/workspace/learning/adventofcode/day06/puzzle-1/input.txt")
	if err != nil {
		return
	}

	totalOrbitCount := TotalOrbitCount(orbitMap)
	fmt.Println("Total Orbit Count: " + strconv.Itoa(totalOrbitCount))
}

func ReadInput(file string) (map[string]Orbit, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	orbitMap := make(map[string]Orbit)

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		orbit := ParseOrbit(scanner.Text())
		orbitMap[orbit.Name] = orbit
	}

	return orbitMap, nil
}

func ParseOrbit(orbitString string) Orbit {
	splits := strings.Split(orbitString, ")")
	return Orbit{Name: splits[1], Parent: splits[0], IndirectCount: -1}
}

func TotalOrbitCount(orbitMap map[string]Orbit) int {
	totalOrbitCount := 0
	for _, orbit := range orbitMap {
		totalOrbitCount += CalcOrbitCount(orbitMap, orbit.Name)
	}
	return totalOrbitCount
}

func CalcOrbitCount(orbitMap map[string]Orbit, name string) int {
	orbitCount := 1

	orbit, ok := orbitMap[name]
	for ok {
		orbit, ok = orbitMap[orbit.Parent]
		if ok { orbitCount++ }
	}

	return orbitCount
}


