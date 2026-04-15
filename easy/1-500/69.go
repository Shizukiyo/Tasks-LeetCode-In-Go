// https://leetcode.com/problems/sqrtx/description/
package main

func mySqrt(x int) int {
	for i := 0; i <= x / 2 + 1; i++ {
		if i * i == x {
			return i
		} else if i * i > x {
			return i - 1
		}
	}

	return 0
}