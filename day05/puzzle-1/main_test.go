package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRunCode_example1(t *testing.T) {
	testProgram := []int{1,0,0,0,99}
	expectedResult := []int{2,0,0,0,99}
	actual, err := RunCode(testProgram, 0)
	assert.Equal(t, nil, err)
	assert.Equal(t, expectedResult, actual)
}

func TestRunCode_example2(t *testing.T) {
	testProgram := []int{2,3,0,3,99}
	expectedResult := []int{2,3,0,6,99}
	actual, err := RunCode(testProgram, 0)
	assert.Equal(t, nil, err)
	assert.Equal(t, expectedResult, actual)
}

func TestRunCode_example3(t *testing.T) {
	testProgram := []int{2,4,4,5,99,0}
	expectedResult := []int{2,4,4,5,99,9801}
	actual, err := RunCode(testProgram, 0)
	assert.Equal(t, nil, err)
	assert.Equal(t, expectedResult, actual)
}

func TestRunCode_example4(t *testing.T) {
	testProgram := []int{1,1,1,4,99,5,6,0,99}
	expectedResult := []int{30,1,1,4,2,5,6,0,99}
	actual, err := RunCode(testProgram, 0)
	assert.Equal(t, nil, err)
	assert.Equal(t, expectedResult, actual)
}

func TestStepCode_returns_error_for_unknown_opcode(t *testing.T) {
	faultyProgram := []int {5,9,10,3,2,3,11,0,99,30,40,50}
	_, _, err := StepCode(faultyProgram, 0)
	assert.NotEqual(t, nil, err)
}

func TestStepCode_return_expected_program_for_proper_input(t *testing.T) {
	program := []int {1,9,10,3,2,3,11,0,99,30,40,50}
	expect  := []int{1,9,10,70,2,3,11,0,99,30,40,50}
	actual, pos, err := StepCode(program, 0)
	assert.Equal(t, nil, err)
	assert.Equal(t, expect, actual)
	assert.Equal(t, 4, pos)
}

func TestStepCode_return_final_pos_for_end_sequence(t *testing.T) {
	program := []int {1,9,10,3,2,3,11,0,99,30,40,50}
	actual, pos, err := StepCode(program, 8)
	assert.Equal(t, program, actual)
	assert.Equal(t, -1, pos)
	assert.Equal(t, nil, err)
}

//func TestStepCode_use_immediate(t *testing.T) {
//	program := []int {1101,100,-1,4,0}
//	_, pos, err := StepCode(program, 0)
//	_, pos, err = StepCode(program, 4)
//	//assert.Equal(t, program, actual)
//	assert.Equal(t, -1, pos)
//	assert.Equal(t, nil, err)
//}
