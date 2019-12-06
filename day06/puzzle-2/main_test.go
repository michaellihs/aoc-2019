package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinHops(t *testing.T) {
	orbitMap, _ := ReadInput("/Users/mimi/workspace/learning/adventofcode/day06/puzzle-2/input-01.txt")
	CountToOrbits(&orbitMap)
	minHops := MinHops(&orbitMap, "SAN", "YOU")
	assert.Equal(t, 4, minHops)
}