package utils

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

func Result(op string, f int, s int, b bool) string {
	var r int

	switch op {
	case "+":
		r = f + s
	case "-":
		r = f - s
	case "*":
		r = f * s
	case "/":
		r = f / s
	default:
		panic("Междy цифрами можно ставить только: + - * /")
	}

	if r < 1 && b {
		panic("Римские не могут быть меньше единицы")
	}
	if b {
		return IntToRoman(r)
	}
	return strconv.Itoa(r)
}

// GetDigits Проверяет, что числа соответствуют условию и возвращает их в int. Оператор оставляет в string
func GetDigits(splitInput []string, b bool) (int, string, int) {
	if len(splitInput) != 3 {
		panic("неверный ввод")
	}
	firstDigit, operator, secondDigit := splitInput[0], splitInput[1], splitInput[2]

	var f, s int
	var err error
	if b {
		if _, ok := Romans[firstDigit]; !ok {
			panic("Это не римская цифра")
		} else {
			f = Romans[firstDigit]
		}
		if _, ok := Romans[secondDigit]; !ok {
			panic("Это не римская цифра")
		} else {
			s = Romans[secondDigit]
		}
	} else {
		f, err = strconv.Atoi(firstDigit)
		if err != nil {
			panic("только числа от 1 до 10")
		}
		s, err = strconv.Atoi(secondDigit)
		if err != nil {
			panic("только числа от 1 до 10")
		}
	}

	if f <= 0 || f > 10 || s <= 0 || s > 10 ||
		strings.HasPrefix(secondDigit, "0") ||
		strings.HasPrefix(firstDigit, "0") {
		panic("Только целые числа от 1 до 10")
	}

	return f, operator, s
}
