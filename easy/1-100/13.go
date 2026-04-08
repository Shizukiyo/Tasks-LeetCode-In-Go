// https://leetcode.com/problems/roman-to-integer/
package main

import "fmt"

func romanToInt(s string) int {

	res := 0

	for i := 0; i < len(s); i++ {
		switch s[i] {
			case 'I':
				if i != len(s) - 1 && (s[i+1] == 'V' || s[i+1] == 'X') {
					if s[i+1] == 'V' {
						res += 4
						i++
					} else {
						res += 9
						i++
					}
				} else {
					res += 1
				}
			case 'V':
				res += 5
			case 'X':
				if i != len(s) - 1 && (s[i+1] == 'L' || s[i+1] == 'C') {
					if s[i+1] == 'L' {
						res += 40
						i++
					} else {
						res += 90
						i++
					}
				} else {
					res += 10
				}
			case 'L':
				res += 50
			case 'C':
				if i != len(s) - 1 && (s[i+1] == 'D' || s[i+1] == 'M') {
					if s[i+1] == 'D' {
						res += 400
						i++
					} else {
						res += 900
						i++
					}
				} else {
					res += 100
				}
			case 'D':
				res += 500
			case 'M':
				res += 1000
		}
	}

    return res
}

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}