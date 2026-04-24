// https://leetcode.com/problems/valid-palindrome/description/
package main

import (
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
    left := 0
	right := len(s)-1
	s = strings.ToLower(s)
	for left < right {
		if !unicode.IsLetter(rune(s[left])) && !unicode.IsDigit(rune(s[left])) {
			left++
			continue
		}
		if !unicode.IsLetter(rune(s[right])) && !unicode.IsDigit(rune(s[right])) {
			right--
			continue
		}
		if s[left] == s[right] {
			left++
			right--
		} else {
			return false
		}
	}

	return true
}