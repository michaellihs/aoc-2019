package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const expected = 19690720

func main() {
	naiveSolution()
	advancedSolution()
}

func advancedSolution() {
	upper, lower := 99, 0
	diff := 99
	steps := 0
	prev, noun, verb := 45, 45, 99
	for noun <= 99 {
		steps++
		program, err := ReadInput("/Users/mimi/workspace/learning/adventofcode/day02/puzzle-2/input.txt")
		if err != nil {
			return
		}
		program[1] = noun
		program[2] = verb
		program, err = RunCode(program, 0)
		if err != nil {
			return
		}
		result := program[0]

		fmt.Printf("[%d] noun: %d verb: %d result: %d diff: %d", steps, noun, verb, result, diff)
		fmt.Println()


		if result < expected && diff > 0 {
			lower = noun
			noun += (upper - noun) / 2
			diff = absolute(prev - noun)
			prev = noun
		} else if result > expected && diff > 0 {
			upper = noun
			noun -= (noun - lower) / 2
			diff = absolute(prev - noun)
			prev = noun
		} else {
			fmt.Println("found - noun must be " + strconv.Itoa(noun))
			break
		}
	}

	verb = 45
	upper, lower = 99, 0
	prev = 45

	for verb <= 99 {
		steps++
		program, err := ReadInput("/Users/mimi/workspace/learning/adventofcode/day02/puzzle-2/input.txt")
		if err != nil {
			return
		}
		program[1] = noun
		program[2] = verb
		program, err = RunCode(program, 0)
		if err != nil {
			return
		}
		result := program[0]

		fmt.Printf("[%d] noun: %d verb: %d result: %d", steps, noun, verb, result)
		fmt.Println()

		if result < expected {
			lower = verb
			verb += (upper - verb) / 2
			prev = verb
		} else if result > expected {
			upper = verb
			verb -= (verb - lower) / 2
			prev = verb
		} else if result == expected {
			fmt.Println("found - verb must be " + strconv.Itoa(verb))
			return
		}
	}
}

func absolute(i int) int {
	if i < 0 {
		return -i
	} else {
		return i
	}
}

func naiveSolution() {
	steps := 0
	noun, verb := 0, 0
	for noun <= 99 {
		for verb <= 99 {
			steps++
			program, err := ReadInput("/Users/mimi/workspace/learning/adventofcode/day02/puzzle-2/input.txt")
			if err != nil {
				return
			}
			program[1] = noun
			program[2] = verb
			program, err = RunCode(program, 0)
			if err != nil {
				return
			}
			fmt.Printf("noun: %d verb: %d result: %d", noun, verb, program[0])
			fmt.Println()

			if program[0] == expected {
				fmt.Printf("noun: %d verb: %d result: %d", noun, verb, program[0])
				fmt.Println()
				fmt.Println((100 * noun) + verb)
				fmt.Println("naive required steps: " + strconv.Itoa(steps))
				return
			}

			verb++
		}
		verb = 0
		noun++
	}
}

func ReadInput(file string) ([]int, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var program []int

	scanner := bufio.NewScanner(f)
	scanner.Scan()
	programStr := scanner.Text()
	progStrings := strings.Split(programStr, ",")
	for _, progString := range(progStrings) {
		progInt, err := strconv.Atoi(progString)
		if err != nil {
			return nil, err
		}
		program = append(program, progInt)
	}
	return program, nil
}

func RunCode(program []int, pos int) ([]int, error) {
	if pos != -1 {
		program, pos, err := StepCode(program, pos)
		if err != nil {
			return nil, err
		}
		RunCode(program, pos)
	}
	return program, nil
}

func StepCode(program []int, pos int) ([]int, int, error) {
	opCode := program[pos]
	switch opCode {
	case 1:
		op1 := program[program[pos + 1]]
		op2 := program[program[pos + 2]]
		program[program[pos + 3]] = op1 + op2
	case 2:
		op1 := program[program[pos + 1]]
		op2 := program[program[pos + 2]]
		program[program[pos + 3]] = op1 * op2
	case 99:
		return program, -1, nil
	default:
		return nil, 0, errors.New("Unknow op code: " + strconv.Itoa(opCode))
	}
	return program, pos + 4, nil
}