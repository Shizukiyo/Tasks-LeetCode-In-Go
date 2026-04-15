// https://leetcode.com/problems/add-binary/description/
package main

import (
	"strconv"
)

func addBinary(a string, b string) string {
	if len(a) < len(b) {
		a, b = b, a
	}

	add := 0
	res := ""
	for i := 0; i < len(a); i++ {
		if len(b) - i > 0 {
			inta, _ := strconv.Atoi(string(a[len(a)-1-i]))
			intb, _ := strconv.Atoi(string(b[len(b)-1-i]))
			cur := add + inta + intb 
			res = strconv.Itoa(cur % 2) + res
			add = cur / 2
		} else {
			inta, _ := strconv.Atoi(string(a[len(a)-1-i]))
			cur := add + inta
			res = strconv.Itoa(cur % 2) + res
			add = cur / 2
		}
	}
	if add == 1 {
		res = "1" + res
	}
	
	return res
}