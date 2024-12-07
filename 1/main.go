package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var a []string           // Left side of input
	var b []string           // Right side of input
	var total_length float64 // Total length

	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()        // Get line of input
		words := strings.Fields(line) // Split file inbetween
		a = append(a, words[0])       // Append left side of split file to left side arr
		b = append(b, words[1])       // Append right side of split file to right side arr
	}

	// Convert read strings to ints
	intsA := make([]int, len(a))
	intsB := make([]int, len(b))

	for i, s := range a {
		intsA[i], _ = strconv.Atoi(s)
	}
	for i, s := range b {
		intsB[i], _ = strconv.Atoi(s)
	}

	// Sort by size
	sort.Ints(intsA)
	sort.Ints(intsB)

	// Measure the difference
	for i := 0; i < max(len(a), len(b)); i++ {
		total_length = total_length + math.Abs(float64(intsA[i])-float64(intsB[i]))
	}

	// Result part 1
	fmt.Print(int(total_length), "\n")

	// Part 2 of AOC1
	total_freq := 0

	for _, num1 := range intsA {
		freq := 0
		for _, num2 := range intsB {
			if num1 == num2 {
				freq = freq + 1
			}
		}
		total_freq = total_freq + num1*freq
	}

	fmt.Print(total_freq)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
