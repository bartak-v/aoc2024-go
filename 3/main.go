package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}
	str := string(file) // convert content to a 'string'
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := re.FindAllString(str, -1)
	result := make([][]int, 0)

	// Convert matches to integers and append to result
	for _, match := range matches {
		num1, _ := strconv.Atoi(string(match[1]))
		num2, _ := strconv.Atoi(string(match[2]))
		result = append(result, []int{num1, num2})
	}
	fmt.Printf("%q\n",result)
}
