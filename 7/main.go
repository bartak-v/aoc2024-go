package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/Knetic/govaluate"
)

type OperationInputs struct {
	Result   int
	Operands []int
}

// Recursive function to build and evaluate all combinations of operators
// Recursive function to build and evaluate all combinations of operators
func evaluateWithGovaluate(operands []int, target int, index int, expression string) bool {
	// Base case: If we've processed all operands, evaluate the expression
	if index == len(operands) {
		eval, err := govaluate.NewEvaluableExpression(expression)
		if err != nil {
			log.Fatalf("Error creating expression: %v", err)
		}

		// Evaluate the expression
		result, err := eval.Evaluate(nil)
		if err != nil {
			log.Fatalf("Error evaluating expression: %v", err)
		}

		// Check if the result matches the target
		if int(result.(float64)) == target {
			fmt.Printf("Expression that matches: %s = %d\n", expression, target)
			return true
		}
		return false
	}

	// Recursive case: Add "+" or "*" between the current expression and the next operand
	nextOperand := operands[index]
	if evaluateWithGovaluate(operands, target, index+1, expression+" + "+fmt.Sprintf("%d", nextOperand)) {
		return true
	}
	if evaluateWithGovaluate(operands, target, index+1, expression+" * "+fmt.Sprintf("%d", nextOperand)) {
		return true
	}

	// No match found
	return false
}

func main() {
	var operationCandidates []OperationInputs
	totalSum := 0 // Total sum of legit operations
	// Open the file
	file, err := os.Open("input")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	// Scan and parse the whole file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()            // Get line of input
		words := strings.Split(line, ":") // Split line inbetween

		var operands []int                           // Intermediate slice of operands
		operands_str := strings.Split(words[1], " ") // Slice of string representation of operands
		for _, v := range operands_str {
			// Convert string operands to ints
			operand_int, _ := strconv.Atoi(v)
			operands = append(operands, operand_int)
		}
		op := OperationInputs{}
		op.Result, _ = strconv.Atoi(words[0])
		op.Operands = operands
		operationCandidates = append(operationCandidates, op)
	}

	// Cycle through the operationCandidates and determine the legit candidates
	for _, v := range operationCandidates {
		// Start recursion with the first operand
		if !evaluateWithGovaluate(v.Operands, v.Result, 1, fmt.Sprintf("%d", v.Operands[0])) {
			fmt.Println("No combination of + or * gives the target.")
		} else {
			totalSum += v.Result
		}
	}
	fmt.Print(totalSum)
}
