package Utils

import (
	"strconv"
	"strings"
)

var Romans = map[string]int{
	"I":    1,
	"II":   2,
	"III":  3,
	"IV":   4,
	"V":    5,
	"VI":   6,
	"VII":  7,
	"VIII": 8,
	"IX":   9,
	"X":    10,
}

func IntToRoman(num int) string {
	var roman string
	var numbers = []int{1, 4, 5, 9, 10, 40, 50, 90, 100}
	var romans = []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C"}
	var index = len(romans) - 1

	for num > 0 {
		for numbers[index] <= num {
			roman += romans[index]
			num -= numbers[index]
		}
		index -= 1
	}
	return roman
}

func DigitsCheck(firstDigit string, secondDigit string, b bool, r map[string]int) {
	var f, s int
	if b {
		f = r[firstDigit]
		s = r[secondDigit]
	} else {
		f, _ = strconv.Atoi(firstDigit)
		s, _ = strconv.Atoi(secondDigit)
	}

	if f <= 0 || f > 10 || s <= 0 || s > 10 ||
		strings.HasPrefix(secondDigit, "0") ||
		strings.HasPrefix(firstDigit, "0") {
		panic("Только целые числа от 1 до 10")
	}
}

func Result(op string, f string, s string, b bool, m map[string]int) string {
	var r, fd, sd int
	if !b {
		fd, _ = strconv.Atoi(f)
		sd, _ = strconv.Atoi(s)
	} else {
		fd, _ = m[f]
		sd, _ = m[s]
	}
	switch op {
	case "+":
		r = fd + sd
	case "-":
		r = fd - sd
	case "*":
		r = fd * sd
	case "/":
		r = fd / sd
	default:
		panic("Междy цифрами можно ставить только: + - * /")
	}

	if r < 1 && b {
		panic("Римские не могут быть меньше единицы")
	} else if b {
		return IntToRoman(r)
	}
	return strconv.Itoa(r)
}
