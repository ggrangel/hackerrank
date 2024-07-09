// Package caesar provides a solution to the following problem:
// https://www.hackerrank.com/challenges/caesar-cipher-1/problem
package caesar

import "strings"

func caesarCipher(plaintext string, shift int) string {
	ciphertext := make([]rune, len(plaintext))
	lowerChars := "abcdefghijklmnopqrstuvwxyz"
	upperChars := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var index int

	for i, c := range plaintext {
		if strings.Contains(lowerChars, string(c)) {
			index = strings.IndexRune(lowerChars, c)
			ciphertext[i] = rune(lowerChars[((index+shift)%26+26)%26])
		} else if strings.Contains(upperChars, string(c)) {
			index = strings.IndexRune(upperChars, c)
			ciphertext[i] = rune(upperChars[((index+shift)%26+26)%26])
		} else {
			ciphertext[i] = c
		}
	}

	return string(ciphertext)
}
