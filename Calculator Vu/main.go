package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	var x string
	var y string
	var z string
	var operator string

	fmt.Scanf("%s", &x)
	fmt.Scanf("%s", &operator)
	fmt.Scanf("%s", &y)
	fmt.Scanf("%s", &z)

	if z != "" {
		fmt.Println("Неверная операция")
		return
	}

	roman1 := make(map[string]int)
	roman1["I"] = 1
	roman1["II"] = 2
	roman1["III"] = 3
	roman1["IV"] = 4
	roman1["V"] = 5
	roman1["VI"] = 6
	roman1["VII"] = 7
	roman1["VIII"] = 8
	roman1["IX"] = 9
	roman1["X"] = 10
	var a int = roman1[x]

	roman2 := make(map[string]int)
	roman2["I"] = 1
	roman2["II"] = 2
	roman2["III"] = 3
	roman2["IV"] = 4
	roman2["V"] = 5
	roman2["VI"] = 6
	roman2["VII"] = 7
	roman2["VIII"] = 8
	roman2["IX"] = 9
	roman2["X"] = 10
	var b int = roman2[y]

	if a > 0 && a <= 10 && b > 0 && b <= 10 {

		switch operator {

		case "+":
			var add int
			add = a + b
			integerToRoman(add)
			fmt.Println(integerToRoman(add))
			return

		case "-":
			var sub int
			sub = a - b
			if sub <= 0 {
				fmt.Println("Отрицательное или нулевое значение")
				return
			}
			integerToRoman(sub)
			fmt.Println(integerToRoman(sub))
			return

		case "*":
			var multiply int
			multiply = a * b
			integerToRoman(multiply)
			fmt.Println(integerToRoman(multiply))
			return

		case "/":
			var divide int
			{
				divide = a / b
				integerToRoman(divide)
				fmt.Println(integerToRoman(divide))
				return
			}

		default:
			fmt.Println("Неверная команда")
			return
		}
	}

	k, _ := strconv.Atoi(x)
	m, _ := strconv.Atoi(y)

	if k > 10 || k == 0 || k < 0 {
		fmt.Println("Неверное значение")
		return
	}
	if m > 10 || m == 0 || m < 0 {
		fmt.Println("Неверное значение")
		return
	}

	switch operator {

	case "+":
		var add int
		add = k + m
		fmt.Println(add)

	case "-":
		var sub int
		sub = k - m
		fmt.Println(sub)

	case "*":
		var multiply int
		multiply = k * m
		fmt.Println(multiply)

	case "/":
		var divide int

		divide = k / m
		fmt.Println(divide)

	default:
		fmt.Println("Неверная команда")
	}

}

func integerToRoman(number int) string {

	conversions := []struct {
		value int
		digit string
	}{
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

	var roman strings.Builder
	for _, conversion := range conversions {
		for number >= conversion.value {
			roman.WriteString(conversion.digit)
			number -= conversion.value
		}
	}

	return roman.String()
}
