package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNumberIsPassword(t *testing.T) {
	isPassword := NumberIsPassword("111111")
	assert.Equal(t, true, isPassword)
}

func TestTwoAdjacentAreTheSame(t *testing.T) {
	same := TwoAdjacentAreTheSame("111111")
	assert.Equal(t, true, same)

	same = TwoAdjacentAreTheSame("112345")
	assert.Equal(t, true, same)

	same = TwoAdjacentAreTheSame("123455")
	assert.Equal(t, true, same)
}

func TestDigitsDontDecrease(t *testing.T) {
	dontDecrease := DigitsDontDecrease("111111")
	assert.Equal(t, true, dontDecrease)

	dontDecrease = DigitsDontDecrease("123456")
	assert.Equal(t, true, dontDecrease)

	dontDecrease = DigitsDontDecrease("212222")
	assert.Equal(t, false, dontDecrease)

	dontDecrease = DigitsDontDecrease("123454")
	assert.Equal(t, false, dontDecrease)
}