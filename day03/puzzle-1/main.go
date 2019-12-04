package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Direction struct {
	Dir   string
	Steps int
}

func NewDirection(s string) Direction {
	steps, _ := strconv.Atoi(string(s[1:len(s)]))
	return Direction{
		Dir:   string(s[0]),
		Steps: steps,
	}
}

func (d Direction) Xy() (int, int) {
	switch d.Dir {
	case "R":
		return d.Steps, 0
	case "L":
		return -d.Steps, 0
	case "U":
		return 0, d.Steps
	case "D":
		return 0, -d.Steps
	default:
		return 0, 0
	}
}

type Point struct {
	x int
	y int
}

func (p Point) Move(dir Direction) Point {
	x, y := dir.Xy()
	return Point{
		p.x + x,
		p.y + y,
	}
}

type Line struct {
	p1 Point
	p2 Point
}

func (l1 Line) Intersect(l2 Line) Point {
	var vert, hor Line
	if (l1.p1.x == l1.p2.x) {
		if (l2.p1.x != l2.p2.x) {
			vert, hor = l1, l2
		}
	} else if l2.p1.x == l2.p2.x {
		vert, hor = l2, l1
	}
	if (hor.p1.x < vert.p1.x && vert.p1.x < hor.p2.x || hor.p2.x < vert.p1.x && vert.p1.x < hor.p1.x) &&
		(vert.p1.y < hor.p1.y && hor.p1.y < vert.p2.y || vert.p2.y < hor.p1.y && hor.p1.y < vert.p1.y) {
		// intersect exists
		return Point{vert.p1.x, hor.p1.y}
	}

	return inftyPoint()
}

func NewLine(x1, y1, x2, y2 int) Line {
	return Line{Point{x1,y1}, Point{x2,y2}}
}

func inftyPoint() Point {
	return Point{math.MaxInt32, math.MaxInt32}
}

func main() {
	w1, w2 := ParseWires("/Users/mimi/workspace/learning/adventofcode/day03/puzzle-1/input.txt")
	lines := Path2Lines(w1)
	shortestIntersect := ShortestIntersect(lines, w2)
	fmt.Println(shortestIntersect)
}

func Path2Lines(directions []Direction) []Line {
	var lines []Line
	currPos := Point{0,0}
	for _, dir := range directions {
		nextPos := currPos.Move(dir)
		lines = append(lines, Line{currPos, nextPos})
		currPos = nextPos
	}
	return lines
}

func ShortestIntersect(lines []Line, w2 []Direction) int {
	shortestIntersect := math.MaxInt64
	startPos := Point{0,0}
	for _, dir := range w2 {
		endPos := startPos.Move(dir)
		currLine := Line{startPos, endPos}
		for _, line := range lines {
			currIntersect := currLine.Intersect(line)
			currDist := TaxiCabDistance(0,0, currIntersect.x, currIntersect.y)
			if currDist < shortestIntersect {
				shortestIntersect = currDist
			}
		}
		startPos = endPos
	}
	return shortestIntersect
}

func ParseWires(file string) ([]Direction, []Direction) {
	f, err := os.Open(file)
	if err != nil {
		os.Exit(1)
	}
	defer f.Close()

	var directions1 []Direction
	var directions2 []Direction

	scanner := bufio.NewScanner(f)
	scanner.Scan()
	directionStrings1 := strings.Split(scanner.Text(), ",")
	for _, s := range directionStrings1 {
		directions1 = append(directions1, NewDirection(s))
	}
	scanner.Scan()
	directionStrings2 := strings.Split(scanner.Text(), ",")
	for _, s := range directionStrings2 {
		directions2 = append(directions2, NewDirection(s))
	}
	return directions1, directions2
}

func ParseWireDirections(wireString string) []string {
	return strings.Split(wireString, ",")
}

func TaxiCabDistance(x1 int, y1 int, x2 int, y2 int) int {
	return abs(x2-x1) + abs(y2-y1)
}

func abs(i int) int {
	if i >= 0 {
		return i
	} else {
		return -i
	}
}
