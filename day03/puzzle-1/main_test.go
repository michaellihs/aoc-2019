package main

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestPoint_Move_right(t *testing.T) {
	move := Point{0,0}.Move(Direction{"R", 2})
	assert.Equal(t, Point{2, 0}, move)
}

func TestPoint_Move_left(t *testing.T) {
	move := Point{0,0}.Move(Direction{"L", 2})
	assert.Equal(t, Point{-2, 0}, move)
}

func TestPoint_Move_up(t *testing.T) {
	move := Point{0,0}.Move(Direction{"U", 2})
	assert.Equal(t, Point{0, 2}, move)
}

func TestPoint_Move_down(t *testing.T) {
	move := Point{0,0}.Move(Direction{"D", 2})
	assert.Equal(t, Point{0, -2}, move)
}

func TestParseWires(t *testing.T) {
	// R8,U5,L5,D3 [{{0 0} {8 0}} {{8 0} {8 5}} {{8 5} {3 5}} {{3 5} {3 2}}]
	// U7,R6,D4,L4

	d1, d2 := ParseWires("/Users/mimi/workspace/learning/adventofcode/day03/puzzle-1/test-input-0.txt")
	assert.Equal(t, []Direction{NewDirection("R8"), NewDirection("U5"), NewDirection("L5"), NewDirection("D3")}, d1)
	assert.Equal(t, []Direction{NewDirection("U7"), NewDirection("R6"), NewDirection("D4"), NewDirection("L4")}, d2)
}

func TestParseWireDirections(t *testing.T){
	directions := ParseWireDirections("R8,U5,L5,D3")
	assert.Equal(t, []string{"R8", "U5", "L5", "D3"}, directions)
}

func TestLine_Intersect_intersects(t *testing.T) {
	l1 := Line{Point{1,1}, Point{1,4}}
	l2 := Line{Point{0,2}, Point{4,2}}
	expIntersect := Point{1,2}

	intersect := l1.Intersect(l2)

	assert.Equal(t, expIntersect, intersect)
}

func TestLine_Intersect_vert_parallel(t *testing.T) {
	l1 := Line{Point{1,1}, Point{1,4}}
	l2 := Line{Point{2,1}, Point{2,4}}
	expIntersect := Point{math.MaxInt32, math.MaxInt32}

	intersect := l1.Intersect(l2)

	assert.Equal(t, expIntersect, intersect)
}

func TestLine_Intersect_hor_parallel(t *testing.T) {
	l1 := Line{Point{1,1}, Point{4,1}}
	l2 := Line{Point{1,2}, Point{4,2}}
	expIntersect := Point{math.MaxInt32, math.MaxInt32}

	intersect := l1.Intersect(l2)

	assert.Equal(t, expIntersect, intersect)
}

func TestLine_Intersect_non_intersect(t *testing.T) {
	l1 := Line{Point{1,3}, Point{1,4}}
	l2 := Line{Point{0,2}, Point{4,2}}
	expIntersect := Point{math.MaxInt32, math.MaxInt32}

	intersect := l1.Intersect(l2)

	assert.Equal(t, expIntersect, intersect)
}

func TestLine_Intersect_non_intersect_touch(t *testing.T) {
	l1 := Line{Point{1,3}, Point{1,4}}
	l2 := Line{Point{1,2}, Point{4,2}}
	expIntersect := Point{math.MaxInt32, math.MaxInt32}

	intersect := l1.Intersect(l2)

	assert.Equal(t, expIntersect, intersect)
}

func TestTaxiCabDistance_of_0_0_and_1_1_is_2(t *testing.T) {
	d := TaxiCabDistance(0,0,1,1)
	assert.Equal(t, 2, d)
}

func TestPointMoveDirR(t *testing.T) {
	p := Point{0,0}
	dir := NewDirection("R10")
	p2 := p.Move(dir)
	assert.Equal(t, 10, p2.x)
	assert.Equal(t, 0, p2.y)
}

func TestPath2Lines(t *testing.T) {
	wire1 := []Direction{NewDirection("R8"), NewDirection("U5"), NewDirection("L5"), NewDirection("D3")}
	lines := Path2Lines(wire1)

	assert.Equal(t, 4, len(lines))
	assert.Equal(t, []Line{NewLine(0,0,8,0), NewLine(8,0,8,5), NewLine(8,5,3,5), NewLine(3,5,3,2)}, lines)
}

func TestShortestIntersect_input_0(t *testing.T) {
	// R8,U5,L5,D3 [{{0 0} {8 0}} {{8 0} {8 5}} {{8 5} {3 5}} {{3 5} {3 2}}]
	// U7,R6,D4,L4 [{{0 0} {0 7}} {{0 7} {6 7}} {{6 7} {6 3}} {{6 3} {2 3}}]
	w1, w2 := ParseWires("/Users/mimi/workspace/learning/adventofcode/day03/puzzle-1/test-input-0.txt")
	lines := Path2Lines(w1)
	shortestIntersect := ShortestIntersect(lines, w2)

	assert.Equal(t, 6, shortestIntersect)
}

func TestShortestIntersect_input_1(t *testing.T) {
	w1, w2 := ParseWires("/Users/mimi/workspace/learning/adventofcode/day03/puzzle-1/test-input-1.txt")
	lines := Path2Lines(w1)
	shortestIntersect := ShortestIntersect(lines, w2)

	assert.Equal(t, 159, shortestIntersect)
}

func TestShortestIntersect_input_2(t *testing.T) {
	w1, w2 := ParseWires("/Users/mimi/workspace/learning/adventofcode/day03/puzzle-1/test-input-2.txt")
	lines := Path2Lines(w1)
	shortestIntersect := ShortestIntersect(lines, w2)

	assert.Equal(t, 135, shortestIntersect)
}

func TestLine_Intersect(t *testing.T) {
	l1 := Line{Point{3, 5}, Point{3, 2}}
	l2 := Line{Point{6, 3}, Point{2, 3}}
	expIntersect := Point{3, 3}

	intersect := l1.Intersect(l2)

	assert.Equal(t, expIntersect, intersect)
}
