// https://leetcode.com/problems/excel-sheet-column-number/description/
package main

func titleToNumber(columnTitle string) int {
    res := 0

	for i := 0; i < len(columnTitle); i++ {
		letter := columnTitle[len(columnTitle)-i-1] - 64
		res = pow(26, i) * int(letter) + res
	}

	return res
}

func pow(num, factor int) int {
	n := 1
	for range factor {
		n *= num
	}
	return n
}