package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var operator string
	var result float64 = 0

	if len(os.Args) < 4 {
		fmt.Println("Limited arguments to calculate")
		return
	}

	// First part of os.Args is always the binary exe file path
	for i := 1; i < len(os.Args); i++ {

		if i == 1 {
			val, err := parseNumber(os.Args[i])
			if err != nil {
				return
			}

			result = val
			continue
		}

		if i%2 == 0 {
			if !checkOperator(os.Args[i]) {
				fmt.Println("An operator is required between both operands")
				return
			} else {
				operator = os.Args[i]
				continue
			}
		}

		val, err := parseNumber(os.Args[i])

		if err != nil {
			fmt.Println("Can't parse this number ", os.Args[i])
			return
		}

		result = handleCalculations(result, operator, val)

		fmt.Println("Result of the calculations are : ", result)
	}
}

func handleCalculations(num1 float64, operator string, num2 float64) float64 {
	if operator == "+" {
		return num1 + num2
	} else if operator == "-" {
		return num1 - num2
	} else if operator == "*" {
		return num1 * num2
	} else if operator == "/" {
		if num2 == 0 {
			fmt.Println("Can't divide the things with 0, skipping this calculation")
			return num1
		}
		return num1 / num2
	}
	return num1
}

func checkOperator(op string) bool {
	if op == "+" || op == "-" || op == "*" || op == "/" {
		return true
	} else {
		return false
	}
}

func parseNumber(operand string) (float64, error) {
	val, err := strconv.ParseFloat(operand, 64)

	if err != nil {
		fmt.Printf("%v is not a valid number", operand)
		return 0, err
	}

	return val, nil
}
