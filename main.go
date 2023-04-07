package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input := readInput()
	a, sign, b, err1 := prepareArgs(input)
	if err1 != nil {
		fmt.Println("неверное выражение")
		return
	}
	a1 := strings.ToUpper(a)
	b1 := strings.ToUpper(b)
	result, err2 := calculate(a1, b1, sign)
	if err2 != nil {
		fmt.Println(err2.Error())
	} else {
		isRomanNumeral, _ := isRomanNumeral(a1)
		if isRomanNumeral == true {
			res, err := toRoman(result)
			if err != nil {
				fmt.Println(fmt.Errorf(err.Error()))
			} else {
				fmt.Println(res)
			}

		} else {
			fmt.Println(result)
		}

	}

}

func readInput() []string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	words := strings.Split(input, " ")
	return words
}

func prepareArgs(in []string) (a, sign, b string, err error) {
	if len(in) != 3 {
		return "", "", "", fmt.Errorf("неверное выражение")
	}

	a = in[0]
	sign = in[1]
	b = in[2]
	return a, sign, b, nil
}
