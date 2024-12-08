package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Check if sequence is strictly increasing
func isIncreasing(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1]-nums[i] > 3 {
			return false
		}
		if nums[i] >= nums[i+1] {
			return false
		}
	}
	return true
}

// Check if sequence is strictly decreasing
func isDecreasing(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i]-nums[i+1] > 3 {
			return false
		}
		if nums[i] <= nums[i+1] {
			return false
		}
	}
	return true
}

// Check if report is safe without Problem Dampener
func isSafeReport(nums []int) bool {
	return isIncreasing(nums) || isDecreasing(nums)
}

// Check if report is safe with Problem Dampener
func isSafeReportWithDampener(nums []int) bool {
	// First, check if already safe
	if isSafeReport(nums) {
		return true
	}

	// Try removing each level
	for i := 0; i < len(nums); i++ {
		// Create a copy of the slice without the i-th element
		reducedNums := make([]int, 0, len(nums)-1)
		reducedNums = append(reducedNums, nums[:i]...)
		reducedNums = append(reducedNums, nums[i+1:]...)

		if isIncreasing(reducedNums) || isDecreasing(reducedNums) {
			return true
		}
	}

	return false
}

func main() {
	// Open the input file
	file, err := os.Open("input")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	safeWithoutDampener := 0
	safeWithDampener := 0

	for scanner.Scan() {
		// Parse the line of numbers
		line := scanner.Text()
		strNums := strings.Fields(line)
		nums := make([]int, len(strNums))

		for j, strNum := range strNums {
			num, _ := strconv.Atoi(strNum)
			nums[j] = num
		}

		// Count safe reports
		if isSafeReport(nums) {
			safeWithoutDampener++
		}
		if isSafeReportWithDampener(nums) {
			safeWithDampener++
		}
	}

	fmt.Printf("Safe reports without Problem Dampener: %d\n", safeWithoutDampener)
	fmt.Printf("Safe reports with Problem Dampener: %d\n", safeWithDampener)

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
