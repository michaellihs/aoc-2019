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
	return TwoAdjacentAreTheSameAndThirdIsNotTheSame(number) && DigitsDontDecrease(number)
}

func TwoAdjacentAreTheSameAndThirdIsNotTheSame(number string) bool {
	for i := 0; i < len(number); i++ {
		if i < len(number) - 1 && number[i] == number[i + 1] {
			if i < len(number) - 4 && number[i] == number[i + 4] {
				i += 4
			} else if i < len(number) - 3 && number[i] == number[i + 3] {
				i += 3
			} else if i < len(number) - 2 && number[i] == number[i + 2] {
				i += 2
			} else {
				return true
			}
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