package main

import (
	"fmt"
	"strconv"
	"strings"
)

func diffWaysToCompute(expression string) []int {
	return compute(expression, make(map[string][]int))
}

func compute(expression string, memo map[string][]int) []int {
	// Check if result is already computed and stored in memo
	if result, found := memo[expression]; found {
		return result
	}

	// If the expression is a single number, return it as the result
	if !strings.ContainsAny(expression, "+-*") {
		num, _ := strconv.Atoi(expression)
		return []int{num}
	}

	var results []int

	// Iterate over the expression to find operators and split the expression
	for i, char := range expression {
		if char == '+' || char == '-' || char == '*' {
			left := expression[:i]
			right := expression[i+1:]

			leftResults := compute(left, memo)
			rightResults := compute(right, memo)

			// Combine results of left and right expressions based on the operator
			for _, l := range leftResults {
				for _, r := range rightResults {
					switch char {
					case '+':
						results = append(results, l+r)
					case '-':
						results = append(results, l-r)
					case '*':
						results = append(results, l*r)
					}
				}
			}
		}
	}

	// Store computed results in memo
	memo[expression] = results
	return results
}

func main() {
	expression1 := "2-1-1"
	expression2 := "2*3-4*5"

	results1 := diffWaysToCompute(expression1)
	results2 := diffWaysToCompute(expression2)

	fmt.Println("Results for expression", expression1, ":", results1)
	fmt.Println("Results for expression", expression2, ":", results2)
}
