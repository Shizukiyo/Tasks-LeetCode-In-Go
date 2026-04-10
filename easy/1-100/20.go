// https://leetcode.com/problems/valid-parentheses/description/
package main

func isValid(s string) bool {
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case ')':
			if !(i != 0 && s[i-1] == '(') {
				return false
			}
			s = s[:i-1] + s[i+1:]
			i-=2
		case ']':
			if !(i != 0 && s[i-1] == '[') {
				return false
			}
			s = s[:i-1] + s[i+1:]
			i-=2
		case '}':
			if !(i != 0 && s[i-1] == '{') {
				return false
			}
			s = s[:i-1] + s[i+1:]
			i-=2
		}
	}
	if len(s) != 0 && (s[len(s)-1] == '(' || s[len(s)-1] == '[' || s[len(s)-1] == '{') {
		return false
	}
	return true
}