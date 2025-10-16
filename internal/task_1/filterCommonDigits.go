package task1

import (
	"errors"
)

func DigistArray(a int, sp *[10]int) {
	for a > 0 {
		digit := a % 10
		sp[digit] = 1
		a /= 10
	}

}

func FinallyDigit(digit int, sp *[10]int) int {
	fDigit := 0
	c := 1
	for digit > 0 {
		a := digit % 10
		digit /= 10
		if sp[a] == 0 {
			fDigit += a * 10 * c
			c++
		}
	}

	return fDigit
}

func FilterCommonDigits(first_digit, second_digit int) (int, int, error) {
	if first_digit == 0 || second_digit == 0 {
		return 0, 0, errors.New("ErrNegNums")
	}

	digits := [10]int{}

	DigistArray(first_digit, &digits)
	DigistArray(second_digit, &digits)

	final_digit_first := FinallyDigit(first_digit, &digits)
	final_digit_second := FinallyDigit(second_digit, &digits)

	if final_digit_first == 0 || final_digit_second == 0 {
		return 0, 0, errors.New("ErrEmptyNum")
	}

	return final_digit_first, final_digit_second, nil
}
