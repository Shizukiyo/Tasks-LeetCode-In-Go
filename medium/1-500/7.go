// https://leetcode.com/problems/reverse-integer/description/
package main

func reverse(x int) int {
	values := make([]int, 0)
    res := 0
	add := 1

	for {
		values = append(values, x%10)
		x/=10
		if (x % 10 == 0 && x / 10 == 0) {
			break
		}
	}

	for i := len(values)-1; i >= 0; i-- {
		if i == len(values)-1 && values[i] == 0 {
			continue
		} 
		res = values[i] * add + res
		add *= 10
	}
	
	if !(-2147483648 < res && res < 2147483647) {
		return 0
	} else {
		return res
	}
}
