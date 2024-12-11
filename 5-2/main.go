package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Hashmap of the ordering rules
	// X -> {A,B,C} is a lookup map which shows that the key must be before all values in an update
	ordering_rules := make(map[int][]int)
	// Input of the update pages
	var update_pages [][]int
	var invalid_update_pages [][]int

	// Open the first input file
	file1, err := os.Open("input1")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file1.Close()

	// Open the second input file
	file2, err := os.Open("input2")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file2.Close()

	// Read the ordering rules and append right parts to arrays that are indexed in the map by the page on the left
	scanner := bufio.NewScanner(file1)
	for scanner.Scan() {
		line := scanner.Text()                    // Get line of input
		words := strings.Split(line, "|")         // Split line inbetween
		index_number, _ := strconv.Atoi(words[0]) // Get the index number
		rule_number, _ := strconv.Atoi(words[1])  // Get the rule
		ordering_rules[index_number] = append(ordering_rules[index_number], rule_number)
	}
	// Read the update pages instructions and append them to an array
	scanner2 := bufio.NewScanner(file2)
	for scanner2.Scan() {
		line := scanner2.Text()           // Get line of input
		words := strings.Split(line, ",") // Split line inbetween
		var update_pages_line []int
		for _, v := range words {
			int_val, _ := strconv.Atoi(v)
			update_pages_line = append(update_pages_line, int_val)
		}
		update_pages = append(update_pages, update_pages_line)
	}

	// Parse each update instruction and based on the map, determine if it is valid or not.
	for _, update := range update_pages {
		seen_elements := make(map[int]bool, len(update))
		valid := true
		for _, current_element := range update {
			// Make mapping for element seen
			seen_elements[current_element] = true
			// Go through the ordering rules for current element and check validity
			for _, element := range ordering_rules[current_element] {
				if seen_elements[element] {
					// An element has been seen before the current element, which breaks the rule.
					// This is invalid update rule.
					valid = false
				}
			}
		}
		if !valid {
			invalid_update_pages = append(invalid_update_pages, update)
		}
	}

	for _, invalid_update := range invalid_update_pages{
		valid_update := sort_invalid_update(invalid_update, ordering_rules)
	}
}

func sort_invalid_update(update []int, rules map[int][]int) []int{
	 // Go through the update and for each element, if you find the next element in the ordering rules, put it immediatedly before the element.
	 
}