package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Input")
	expression, _ := reader.ReadString('\n')

	expressionCheck := check(expression)
	result := calculate(expressionCheck)
	fmt.Printf("Output: %s \n", result)
}

var romansToArabic = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

var arabicsToRoman = map[int]string{
	1:    "I",
	4:    "IV",
	5:    "V",
	9:    "IX",
	10:   "X",
	40:   "XL",
	50:   "L",
	90:   "XC",
	100:  "C",
	400:  "CD",
	500:  "D",
	900:  "CM",
	1000: "M",
}

func calculate(expression map[int]string) string {
	var res int
	var result string

	var operand1, _ = strconv.Atoi(expression[0])
	var operand2, _ = strconv.Atoi(expression[2])
	var operator = expression[1]
	var flag = expression[3]

	switch operator {
	case "+":
		res = operand1 + operand2
	case "-":
		res = operand1 - operand2
	case "*":
		res = operand1 * operand2
	case "/":
		if operand2 == 0 {
			fmt.Println("деление на ноль")
			os.Exit(1)
		}
		res = operand1 / operand2
	default:
		fmt.Println("неизвестный оператор", operator)
		os.Exit(1)
	}

	if flag == "roman" {
		if res <= 0 {
			fmt.Print("в римской системе исчесления нет отрицательных чисел")
			os.Exit(1)
		}
		result = arabicToRoman(res)
	} else {
		result = strconv.Itoa(res)
	}

	return result
}

func check(expr string) map[int]string {
	expr = strings.TrimSpace(expr)

	parts := strings.Split(expr, " ")
	var flag1 string
	var flag2 string

	if len(parts) > 3 {
		fmt.Print("неверное выражение, много аргументов")
		os.Exit(1)
	} else if len(parts) < 2 {
		fmt.Print("неверное выражение, мало аргументов")
		os.Exit(1)
	}

	parts0Rom := romanToArabic(parts[0])
	parts2Rom := romanToArabic(parts[2])
	var operand1, operand2 string

	if parts0Rom != 0 {
		operand1 = strconv.Itoa(parts0Rom)
		flag1 = "roman"
	} else {
		parts0Arab, _ := strconv.Atoi(parts[0])
		if parts0Arab <= 0 || parts0Arab > 10 {
			fmt.Print("неверный первый операнд, должен быть целым числом от 1 до 10")
			os.Exit(1)
		}
		operand1 = strconv.Itoa(parts0Arab)
		flag1 = "arabic"
	}

	if parts2Rom != 0 {
		operand2 = strconv.Itoa(parts2Rom)
		flag2 = "roman"
	} else {
		parts2Arab, _ := strconv.Atoi(parts[2])
		if parts2Arab <= 0 || parts2Arab > 10 {
			fmt.Print("неверный второй операнд, должен быть целым числом от 1 до 10")
			os.Exit(1)
		}
		operand2 = strconv.Itoa(parts2Arab)
		flag2 = "arabic"
	}

	if flag1 != flag2 {
		fmt.Print("оба операнда долджны быть из одной системы исчесления")
		os.Exit(1)
	}

	var expCheck = map[int]string{

		0: operand1,
		1: parts[1],
		2: operand2,
		3: flag1,
	}
	return expCheck
}

func arabicToRoman(number int) string {
	maxRomanNumber := 3999
	if number > maxRomanNumber {
		return strconv.Itoa(number)
	}

	conversions := []struct {
		value int
		digit string
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

	var roman strings.Builder
	for _, conversion := range conversions {
		for number >= conversion.value {
			roman.WriteString(conversion.digit)
			number -= conversion.value
		}
	}

	return roman.String()
}

func romanToArabic(romanNum string) int {
	fChar := strings.Split(romanNum, "")

	if fChar[0] == "-" {
		fmt.Print("Римская цифра не может быть отрицательной")
		os.Exit(1)
	}

	result := 0
	for i := 0; i < len(romanNum); i++ {
		current := romansToArabic[string(romanNum[i])]
		if i+1 < len(romanNum) {
			next := romansToArabic[string(romanNum[i+1])]
			if current < next {
				result -= current
			} else {
				result += current
			}
		} else {
			result += current
		}
	}
	return result
}
