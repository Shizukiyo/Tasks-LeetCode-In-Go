// https://leetcode.com/problems/plus-one/
package main

func plusOne(digits []int) []int {
	res := append([]int{0}, digits...)

	for i := len(res)-1; i >= 0; i-- {
		if res[i] == 9 {
			res[i] = 0
			continue
		} else {
			res[i]++
			break
		}
	}

	if res[0] == 0 {
		return res[1:]
	} else {
		return res
	}
}