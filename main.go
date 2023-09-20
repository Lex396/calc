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
	fmt.Printf("Output: %d \n", result)
}

func calculate(expression map[int]string) int {
	var result int

	var operand1, _ = strconv.Atoi(expression[0])
	var operand2, _ = strconv.Atoi(expression[2])
	var operator = expression[1]
	switch operator {
	case "+":
		result = operand1 + operand2
	case "-":
		result = operand1 - operand2
	case "*":
		result = operand1 * operand2
	case "/":
		if operand2 == 0 {
			fmt.Println("деление на ноль")
			os.Exit(1)
		}
		result = operand1 / operand2
	default:
		fmt.Println("неизвестный оператор", operand1)
		os.Exit(1)
	}
	return result
}

func check(expr string) map[int]string {
	expr = strings.TrimSpace(expr)

	parts := strings.Split(expr, " ")

	if len(parts) > 3 {
		fmt.Print("неверное выражение, много аргументов")
		os.Exit(1)
	} else if len(parts) < 2 {
		fmt.Print("неверное выражение, мало аргументов")
		os.Exit(1)
	}

	operand1, _ := strconv.Atoi(parts[0])
	operand2, _ := strconv.Atoi(parts[2])

	if operand1 <= 0 || operand1 > 10 {
		fmt.Print("неверный первый операнд, должен быть целым числом от 1 до 10")
		os.Exit(1)
	}

	if operand2 <= 0 || operand2 > 10 {
		fmt.Print("неверный второй операнд, должен быть целым числом от 1 до 10")
		os.Exit(1)
	}

	var expCheck = map[int]string{
		0: strconv.Itoa(operand1),
		1: parts[1],
		2: strconv.Itoa(operand2),
	}
	return expCheck
}

func integerToRoman(number int) string {
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

//получение выражения из консоли
//разбить выражение на массив по пробелу

/*определить значение в нулевом индексе;
если не число, сравнить с массивом римских цифр;
если не совпадает, вывести ошибку
если вначале минус, вывести ошибку
если не число и не римская цифра, то вывести ошибку;
если не целое число вывести ошибку;
*/

// определить первый индекс массива
// определить математическую операцию
// вызвать функию с мат операцие
// проверить результат, если результат не удовлетворяет условиям вывести ошибку
// вывести результат
