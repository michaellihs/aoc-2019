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
	program, err := ReadInput("/Users/mimi/workspace/learning/adventofcode/day05/puzzle-1/input.txt")
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

	program := make([]int, 2000)

	scanner := bufio.NewScanner(f)
	scanner.Scan()
	programStr := scanner.Text()
	progStrings := strings.Split(programStr, ",")
	for index, progString := range (progStrings) {
		progInt, err := strconv.Atoi(progString)
		if err != nil {
			return nil, err
		}
		program[index] = progInt
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
	par1Mode, par2Mode := 0, 0
	if opCode > 99 {
		opString := strconv.Itoa(opCode)
		B, C, D, E := 0, 0, 0, 0
		E,_ = strconv.Atoi(string(opString[len(opString)-1]))
		D,_ = strconv.Atoi(string(opString[len(opString)-2]))
		if len(opString) > 2 {
			C,_ = strconv.Atoi(string(opString[len(opString)-3]))
		}
		if len(opString) > 3 {
			B,_ = strconv.Atoi(string(opString[len(opString)-4]))
		}
		opCode = 10 * D + E
		par1Mode = C
		par2Mode = B
	}
	switch opCode {
	case 1:
		op1 := getParam(par1Mode, program, pos, 1)
		op2 := getParam(par2Mode, program, pos, 2)
		op3 := program[pos+3]
		program[op3] = op1 + op2
		pos += 4
	case 2:
		op1 := getParam(par1Mode, program, pos, 1)
		op2 := getParam(par2Mode, program, pos, 2)
		op3 := program[pos+3]
		program[op3] = op1 * op2
		pos += 4
	case 3:
		op1 := program[pos+1]
		program[op1] = 1
		pos += 2
	case 4:
		op1 := program[program[pos+1]]
		fmt.Println(op1)
		pos += 2
	case 99:
		return program, -1, nil
	default:
		return nil, -1, errors.New("Unkonw opcode " + strconv.Itoa(opCode))
	}
	return program, pos, nil
}

func getParam(mode int, program []int, pos int, offset int) int {
	if mode == 0 {
		return program[program[pos+offset]]
	} else {
		return program[pos+offset]
	}
}
