package main

import (
	"fmt"
	"strconv"
)

func main() {
	var passwords []string
	for i := 146810; i <= 612564; i++ {
		number := strconv.Itoa(i)
		if NumberIsPassword(number) {
			passwords = append(passwords, number)
		}
	}
	fmt.Println("Number of Passwords: " + strconv.Itoa(len(passwords)))
}

func NumberIsPassword(number string) bool {
	return TwoAdjacentAreTheSame(number) && DigitsDontDecrease(number)
}

func TwoAdjacentAreTheSame(number string) bool {
	for pos, _ := range number {
		if pos < len(number) - 1 && number[pos] == number[pos + 1]  {
			return true
		}
	}
	return false
}

func DigitsDontDecrease(number string) bool {
	for pos, _ := range number {
		if pos < len(number) - 1 && int(number[pos]) > int(number[pos + 1]) {
			return false
		}
	}
	return true
}