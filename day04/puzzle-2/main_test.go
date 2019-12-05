package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNumberIsPassword(t *testing.T) {
	isPassword := NumberIsPassword("111122")
	assert.Equal(t, true, isPassword)

	isPassword = NumberIsPassword("112233")
	assert.Equal(t, true, isPassword)

	isPassword = NumberIsPassword("123444")
	assert.Equal(t, false, isPassword)

	isPassword = NumberIsPassword("111122")
	assert.Equal(t, true, isPassword)

	isPassword = NumberIsPassword("133333")
	assert.Equal(t, false, isPassword)

	isPassword = NumberIsPassword("588889")
	assert.Equal(t, false, isPassword)

	isPassword = NumberIsPassword("555888")
	assert.Equal(t, false, isPassword)
}

func TestTwoAdjacentAreTheSameAndThirdIsNotTheSame(t *testing.T) {
	same := TwoAdjacentAreTheSameAndThirdIsNotTheSame("111111")
	assert.Equal(t, false, same)

	same = TwoAdjacentAreTheSameAndThirdIsNotTheSame("112345")
	assert.Equal(t, true, same)

	same = TwoAdjacentAreTheSameAndThirdIsNotTheSame("112233")
	assert.Equal(t, true, same)

	same = TwoAdjacentAreTheSameAndThirdIsNotTheSame("123444")
	assert.Equal(t, false, same)

	same = TwoAdjacentAreTheSameAndThirdIsNotTheSame("111122")
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