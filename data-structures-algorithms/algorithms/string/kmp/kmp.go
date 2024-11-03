package main

import (
	"fmt"
)

// Function to compute the KMP prefix array
func computeKMPTable(s string) []int {
	n := len(s)
	lps := make([]int, n)
	length := 0
	i := 1

	for i < n {
		if s[i] == s[length] {
			length++
			lps[i] = length
			i++
		} else {
			if length != 0 {
				length = lps[length-1]
			} else {
				lps[i] = 0
				i++
			}
		}
	}
	return lps
}

// Function to find the shortest palindrome
func shortestPalindrome(s string) string {
	if len(s) == 0 {
		return s
	}

	// Create the modified string: original + "#" + reverse(original)
	rev := reverse(s)
	combined := s + "#" + rev

	// Compute the KMP table for the combined string
	lps := computeKMPTable(combined)

	// Find the characters to add to the beginning of the original string
	toAdd := rev[:len(s)-lps[len(lps)-1]]

	return toAdd + s
}

// Helper function to reverse a string
func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// Main function to test the solution
func main() {
	fmt.Println(shortestPalindrome("aacecaaa")) // Output: "aaacecaaa"
	fmt.Println(shortestPalindrome("abcd"))     // Output: "dcbabcd"
}
