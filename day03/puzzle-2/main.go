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

func (p Point) Equals(point Point) bool {
	if p.x == point.x && p.y == point.y {
		return true
	} else {
		return false
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
	return Line{Point{x1, y1}, Point{x2, y2}}
}

func inftyPoint() Point {
	return Point{math.MaxInt32, math.MaxInt32}
}

type Intersection struct {
	point  Point
	count1 int
	count2 int
}

func (i Intersection) StepCount(w1, w2 []Direction) (int, int) {
	count1 := countSteps(i, w1)
	count2 := countSteps(i, w2)

	return count1, count2
}

func countSteps(intersection Intersection, w []Direction) int {
	stepCount := 1
	currPos := Point{0,0}
	for _, dir := range w {
		for i := 1; i <= dir.Steps; i++ {
			currPos = currPos.Move(Direction{dir.Dir, 1})
			if !intersection.point.Equals(currPos) {
				stepCount++
			} else {
				return stepCount
			}
		}
	}
	return math.MaxInt32
}

func main() {
	w1, w2 := ParseWires("/Users/mimi/workspace/learning/adventofcode/day03/puzzle-1/input.txt")
	lines := Path2Lines(w1)
	fewestSteps := FewestSteps(lines, w1, w2)
	fmt.Println(fewestSteps)
}

func Path2Lines(directions []Direction) []Line {
	var lines []Line
	currPos := Point{0, 0}
	for _, dir := range directions {
		nextPos := currPos.Move(dir)
		lines = append(lines, Line{currPos, nextPos})
		currPos = nextPos
	}
	return lines
}

func FewestSteps(lines []Line, w1, w2 []Direction) int {
	fewestSteps := math.MaxInt64
	startPos := Point{0, 0}
	for _, dir := range w2 {
		endPos := startPos.Move(dir)
		currLine := Line{startPos, endPos}
		for _, line := range lines {
			currIntersect := currLine.Intersect(line)
			if currIntersect.x < math.MaxInt32 {
				intersect := Intersection{currIntersect, 0, 0}
				count1, count2 := intersect.StepCount(w1, w2)
				if (count1 + count2) < fewestSteps {
					fewestSteps = count1 + count2
				}
			}
		}
		startPos = endPos
	}
	return fewestSteps
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
