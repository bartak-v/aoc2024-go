package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func evaluateExpression(operands []int, operators []rune) int {
	result := operands[0]
	for i := 0; i < len(operators); i++ {
		switch operators[i] {
		case '+':
			result += operands[i+1]
		case '*':
			result *= operands[i+1]
		}
	}
	return result
}

func generateOperatorCombinations(length int) [][]rune {
	operators := []rune{'+', '*'}
	var combinations [][]rune

	for i := 0; i < pow(len(operators), length); i++ {
		combo := make([]rune, length)
		n := i
		for j := 0; j < length; j++ {
			combo[j] = operators[n%len(operators)]
			n /= len(operators)
		}
		combinations = append(combinations, combo)
	}
	return combinations
}

func pow(base, exp int) int {
	result := 1
	for i := 0; i < exp; i++ {
		result *= base
	}
	return result
}

func solve(target int, operands []int) bool {
	// Number of spaces between operands where operators can be inserted
	operatorSpaces := len(operands) - 1

	// Generate all possible operator combinations
	operatorCombos := generateOperatorCombinations(operatorSpaces)

	for _, operators := range operatorCombos {
		if evaluateExpression(operands, operators) == target {
			return true
		}
	}
	return false
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	totalCalibration := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		target, _ := strconv.Atoi(parts[0])
		
		operandStrings := strings.Fields(parts[1])
		operands := make([]int, len(operandStrings))
		for i, s := range operandStrings {
			operands[i], _ = strconv.Atoi(s)
		}

		if solve(target, operands) {
			totalCalibration += target
			fmt.Printf("Solved: %d with %v\n", target, operands)
		}
	}

	fmt.Println("Total Calibration Result:", totalCalibration)
}