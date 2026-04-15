// https://leetcode.com/problems/happy-number/description/
package main

import (
	"fmt"
)

func isHappy(n int) bool {
	cur := n
	i := 0

	for ; cur != 1; i++ {
		sum := 0
		for {
			sum += (cur % 10) * (cur % 10)
			cur /= 10

			if cur == n*n {
				return false
			}
			if (cur % 10 == 0 && cur / 10 == 0) {
				break
			}
		}
		cur = sum

		if i == 10 {return false}
	}

	return cur == 1
}

func main() {
	fmt.Println(isHappy(6))
	// fmt.Println(isHappy(2))
	// fmt.Println(isHappy(1111111))
}