// https://leetcode.com/problems/palindrome-number/description/
package main

import (
	"strconv"
)

func isPalindrome(x int) bool {

	sx := strconv.Itoa(x)
    
	leftI := 0
	rightI := len(sx)-1

	for {
		if sx[leftI] == sx[rightI] {
			leftI++
			rightI--
		} else {
			return false
		}
		if leftI >= rightI {
			break
		}
	}

	return true
}
