// https://leetcode.com/problems/longest-common-prefix/
package main

func longestCommonPrefix(strs []string) string {
    longPrefix := strs[0]

	for _, str := range strs {
		if longPrefix == "" {
			return ""
		}
		i := 0
		for ; i < len(min(str, longPrefix)); i++ {
			if str[i] != longPrefix[i] {
				break
			}
		}
		longPrefix = longPrefix[:i]
	}

	return longPrefix
}