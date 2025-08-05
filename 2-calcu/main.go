package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	operator := input()
	inputStr := inputNumbers()

	numbers := []int{} // Empty slice

	// Clean the input by removing all whitespace first
	cleanedInput := strings.ReplaceAll(inputStr, " ", "")
	// Then split by commas if any
	numStrings := strings.Split(cleanedInput, ",")

	for _, numStr := range numStrings {
		if numStr == "" {
			continue
		}
		num, err := strconv.Atoi(numStr)
		if err != nil {
			fmt.Printf("Error converting %v: %v\n", numStr, err)
			continue
		}
		numbers = append(numbers, num)
	}

	if len(numbers) == 0 {
		fmt.Println("No valid numbers provided")
		return
	}

	var result interface{}
	switch strings.ToUpper(operator) {
	case "AVG":
		result = average(numbers)
	case "SUM":
		result = sum(numbers)
	case "MED":
		result = median(numbers)
	default:
		fmt.Println("Invalid operator. Please choose AVG, SUM, or MED")
		return
	}

	fmt.Printf("Here is the final answer of the %s operation on %v:\n", operator, numbers)
	fmt.Println(result)
}

func input() string {
	var reply string
	text := "Choose one of the following operations: AVG, SUM, MED"
	fmt.Println(text)
	fmt.Scan(&reply)
	return reply
}

func inputNumbers() string {
	var reply string
	text := "Input the numbers to which you want to perform the statistical operation on (separated by commas):"
	fmt.Println(text)
	fmt.Scanln(&reply) // Use Scanln to handle potential spaces
	return reply
}

func sum(numbers []int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

func average(numbers []int) float64 {
	if len(numbers) == 0 {
		return 0
	}
	return float64(sum(numbers)) / float64(len(numbers))
}

func median(numbers []int) float64 {
	if len(numbers) == 0 {
		return 0
	}

	nums := make([]int, len(numbers))
	copy(nums, numbers)

	sort.Ints(nums)

	n := len(nums)
	if n%2 == 1 {
		return float64(nums[n/2])
	}
	return float64(nums[n/2-1]+nums[n/2]) / 2.0
}
