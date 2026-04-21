// https://leetcode.com/problems/integer-to-roman/
package main

func intToRoman(num int) string {
	values := make([]int, 0)
	add := 1
	for {
		values = append(values, (num%10)*add)
		num /= 10
		add *= 10
		if num/10 == 0 && num%10 == 0 {
			break
		}
	}

	res := ""

	for _, value := range values {
		switch {
		case value <= 9 && value != 0:
			switch {
			case value == 9:
				res = "IX" + res
			case value == 4:
				res = "IV" + res
			case value < 4:
				for range value {
					res = "I" + res
				}
			case value < 9:
				for range value - 5 {
					res = "I" + res
				}
				res = "V" + res
			}

		case value <= 90 && value != 0:
			switch {
			case value == 90:
				res = "XC" + res
			case value == 40:
				res = "XL" + res
			case value < 40:
				for i := 0; i < value; i += 10 {
					res = "X" + res
				}
			case value < 90:
				for i := 0; i < value-50; i += 10 {
					res = "X" + res
				}
				res = "L" + res
			}

		case value <= 900 && value != 0:
			switch {
			case value == 900:
				res = "CM" + res
			case value == 400:
				res = "CD" + res
			case value < 400:
				for i := 0; i < value; i += 100 {
					res = "C" + res
				}
			case value < 900:
				for i := 0; i < value-500; i += 100 {
					res = "C" + res
				}
				res = "D" + res
			}
		case value >= 1000:
			for i := 0; i < value; i += 1000 {
				res = "M" + res
			}
		}
	}

	return res
}
