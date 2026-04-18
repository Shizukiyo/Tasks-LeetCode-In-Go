// https://leetcode.com/problems/mirror-distance-of-an-integer/description/
package main

func mirrorDistance(n int) int {
    copyn := n
	m := 0

	for copyn != 0 {
		m = m*10 + copyn%10
		copyn /= 10
	}

	if n >= m {
		return n - m
	} else {
		return m - n
	}
}