package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	program, err := ReadInput("/Users/mimi/workspace/learning/adventofcode/day02/puzzle-1/input.txt")
	if err != nil {
		return
	}
	program, err = RunCode(program, 0)
	if err != nil {
		return
	}
	fmt.Println(program)
	fmt.Println(program[0])

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