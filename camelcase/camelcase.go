// Package camelcase provides a solution to the following problem:
// https://www.hackerrank.com/challenges/camelcase/problem
package camelcase

import "strings"

func camelCase(words string) int {
	n := 0
	uppercase := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	if len(words) > 0 {
		n++
	}

	for _, c := range words {
		if strings.Contains(uppercase, string(c)) {
			n++
		}
	}

	return n
}
