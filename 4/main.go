package main

import (
	"bufio"
	"fmt"
	"os"
)

// Coordinates of "X"
type XPosition struct {
	Row int
	Col int
}

func main() {
	var totalXmasCounter int = 0

	// Open the file
	file, err := os.Open("input")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a 2D array to hold the character matrix
	var characterMatrix [][]string

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var chars []string
		for _, char := range line {
			chars = append(chars, string(char))
		}
		characterMatrix = append(characterMatrix, chars)
	}

	// Check for errors in reading the file
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Iterate through all cells in the matrix
	for i := 0; i < len(characterMatrix); i++ {
		for j := 0; j < len(characterMatrix[0]); j++ {
			// Check for "XMAS" in all directions from each cell
			totalXmasCounter += checkHorizontal(i, j, characterMatrix, 1) // Right
			totalXmasCounter += checkHorizontal(i, j, characterMatrix, -1) // Left
			totalXmasCounter += checkVertical(i, j, characterMatrix, 1) // Down
			totalXmasCounter += checkVertical(i, j, characterMatrix, -1) // Up
			totalXmasCounter += checkDiagonal(i, j, characterMatrix, 1, 1) // Down-Right
			totalXmasCounter += checkDiagonal(i, j, characterMatrix, 1, -1) // Down-Left
			totalXmasCounter += checkDiagonal(i, j, characterMatrix, -1, 1) // Up-Right
			totalXmasCounter += checkDiagonal(i, j, characterMatrix, -1, -1) // Up-Left
		}
	}

	fmt.Println(totalXmasCounter)
}

// Checks for "XMAS" in a horizontal direction
func checkHorizontal(row, col int, matrix [][]string, step int) int {
	if col < 0 || col+3*step < 0 || col+3*step >= len(matrix[0]) {
		return 0
	}
	if matrix[row][col] == "X" && matrix[row][col+step] == "M" &&
		matrix[row][col+2*step] == "A" && matrix[row][col+3*step] == "S" {
		return 1
	}
	return 0
}

// Checks for "XMAS" in a vertical direction
func checkVertical(row, col int, matrix [][]string, step int) int {
	if row < 0 || row+3*step < 0 || row+3*step >= len(matrix) {
		return 0
	}
	if matrix[row][col] == "X" && matrix[row+step][col] == "M" &&
		matrix[row+2*step][col] == "A" && matrix[row+3*step][col] == "S" {
		return 1
	}
	return 0
}

// Checks for "XMAS" in a diagonal direction
func checkDiagonal(row, col int, matrix [][]string, rowStep, colStep int) int {
	if row < 0 || col < 0 || row+3*rowStep < 0 || row+3*rowStep >= len(matrix) ||
		col+3*colStep < 0 || col+3*colStep >= len(matrix[0]) {
		return 0
	}
	if matrix[row][col] == "X" && matrix[row+rowStep][col+colStep] == "M" &&
		matrix[row+2*rowStep][col+2*colStep] == "A" && matrix[row+3*rowStep][col+3*colStep] == "S" {
		return 1
	}
	return 0
}
