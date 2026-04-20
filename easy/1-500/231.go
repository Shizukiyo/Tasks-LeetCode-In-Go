// https://leetcode.com/problems/power-of-two/description/
package main

func isPowerOfTwo(n int) bool {
	// 8 - 1000
	// 7 - 0111
	// 1000 & 0111 = 0000
	return n > 0 && (n & (n - 1)) == 0
}
