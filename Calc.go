package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func romanToArabic(romanNumeral string) string {
	romanNumerals := map[rune]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}

	var result int
	var lastChar rune
	for i := len(romanNumeral) - 1; i >= 0; i-- {
		char := rune(romanNumeral[i])
		if i == len(romanNumeral)-1 || romanNumerals[char] >= romanNumerals[lastChar] {
			result += romanNumerals[char]
		} else {
			result -= romanNumerals[char]
		}
		lastChar = char
	}

	resultstring := strconv.Itoa(result)

	return resultstring
}

func calculate(a, b, sign string) (string, error) {
	numA, numB, err := checkStrings(a, b)

	if err != nil {
		return "", err
	}

	numAint, _ := strconv.Atoi(numA)
	numBint, _ := strconv.Atoi(numB)

	switch sign {
	case "+":
		return strconv.Itoa(numAint + numBint), nil
	case "-":
		return strconv.Itoa(numAint - numBint), nil
	case "*":
		return strconv.Itoa(numAint * numBint), nil
	case "/":
		if numBint == 0 {
			return "", errors.New("Ошибка: деление на ноль")
		}
		return strconv.Itoa(numAint / numBint), nil
	default:
		return "", fmt.Errorf("Ошибка: неизвестная операция %s", sign)
	}
}

func checkStrings(a, b string) (res1, res2 string, err error) {
	aTypeIsRoman, aerr := isRomanNumeral(a)
	bTypeIsRoman, berr := isRomanNumeral(b)

	if aerr != nil {
		return "0", "0", aerr
	}

	if berr != nil {
		return "0", "0", berr
	}

	if aTypeIsRoman != bTypeIsRoman {
		return "0", "0", fmt.Errorf("нельзя скрестить ужа с ежом %s и %s", a, b)
	}

	if aTypeIsRoman == true && bTypeIsRoman == true {
		res1 = romanToArabic(a)
		res2 = romanToArabic(b)
	} else {
		res1 = a
		res2 = b
	}

	return res1, res2, nil
}

func toRoman(p string) (string, error) {
	n, _ := strconv.Atoi(p)
	if n < 0 {
		return "", fmt.Errorf("В римской системе нет отрицательных чисел")
	}
	romanNumerals := []struct {
		Value  int
		Symbol string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	var result strings.Builder
	for _, pair := range romanNumerals {
		for n >= pair.Value {
			result.WriteString(pair.Symbol)
			n -= pair.Value
		}
	}
	return result.String(), nil
}

func isRomanNumeral(s string) (bool, error) {
	romanNumerals := map[rune]bool{'I': true, 'V': true, 'X': true, 'L': true, 'C': true, 'D': true, 'M': true}

	var isRoman, isNumeric bool
	for _, char := range s {
		if _, ok := romanNumerals[char]; ok {
			isRoman = true
		} else if _, err := strconv.Atoi(string(char)); err == nil {
			isNumeric = true
		} else {
			return false, fmt.Errorf("Ошибка: %s не является римской цифрой или числом", s)
		}
	}

	if isRoman && !isNumeric {
		return true, nil
	}
	if isNumeric && !isRoman {
		return false, nil
	}
	return false, fmt.Errorf("Ошибка: %s не является римской цифрой или числом", s)
}
