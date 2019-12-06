package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Orbit struct {
	Name          string
	Parent        string
	IndirectCount int
	CountTo       map[string]int
}

func main() {
	orbitMap, err := ReadInput("/Users/mimi/workspace/learning/adventofcode/day06/puzzle-1/input.txt")
	if err != nil {
		return
	}
	CountToOrbits(&orbitMap)
	minHops := MinHops(&orbitMap, "SAN", "YOU")
	fmt.Println(minHops)
}

func MinHops(orbits *map[string]Orbit, start string, destination string) int {
	count := 0
	destOrbit := (*orbits)[destination]
	parent, ok := (*orbits)[(*orbits)[start].Parent]
	for ok {
		otherCount, ok := destOrbit.CountTo[parent.Name]
		if ok {
			return count + otherCount - 1
		}
		parent, ok = (*orbits)[parent.Parent]
		count++
	}
	return 0
}

func CountToOrbits(orbits *map[string]Orbit) {
	for _, orbit := range *orbits {
		CountToOrbit(orbits, orbit)
	}
}

func CountToOrbit(orbits *map[string]Orbit, orbit Orbit) {
	count := 1
	orbit.CountTo = make(map[string]int)
	parent, ok := (*orbits)[orbit.Parent]
	for ok {
		orbit.CountTo[parent.Name] = count
		count++
		parent, ok = (*orbits)[parent.Parent]
		(*orbits)[orbit.Name] = orbit
	}
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


