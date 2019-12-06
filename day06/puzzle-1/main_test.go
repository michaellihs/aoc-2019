package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseOrbit(t *testing.T) {
	orbit := ParseOrbit("N7R)Z77")

	assert.Equal(t, "Z77", orbit.Name)
	assert.Equal(t, "N7R", orbit.Parent)
	assert.Equal(t, -1, orbit.IndirectCount)
}

func TestReadInput(t *testing.T) {
	orbitMap, err := ReadInput("/Users/mimi/workspace/learning/adventofcode/day06/puzzle-1/input-01.txt")

	assert.Equal(t, nil, err)
	assert.Equal(t, 11, len(orbitMap))
	assert.Equal(t, Orbit{Name: "B", Parent: "COM", IndirectCount: -1}, orbitMap["B"])
}

func TestCalcOrbitCount(t *testing.T) {
	orbitMap, _ := ReadInput("/Users/mimi/workspace/learning/adventofcode/day06/puzzle-1/input-01.txt")

	count := CalcOrbitCount(orbitMap, "D")

	assert.Equal(t, 3, count)
}

func TestTotalOrbitCount(t *testing.T) {
	orbitMap, _ := ReadInput("/Users/mimi/workspace/learning/adventofcode/day06/puzzle-1/input-01.txt")
	totalOrbitCount := TotalOrbitCount(orbitMap)

	assert.Equal(t, 42, totalOrbitCount)
}