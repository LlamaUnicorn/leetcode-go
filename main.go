package main

import (
	"fmt"
	//"strings"
)

//func longestCommonPrefix(strs []string) string {
//	if len(strs) == 0 {
//		return ""
//	}
//
//	prefix := strs[0]
//	for _, word := range strs[1:] {
//		for !strings.HasPrefix(word, prefix) {
//			if len(prefix) == 0 {
//				return ""
//			}
//			prefix = prefix[:len(prefix)-1]
//		}
//	}
//
//	return prefix
//}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]
	for _, word := range strs[1:] {
		prefix = commonPrefix(prefix, word)
		if prefix == "" {
			return ""
		}
	}

	return prefix
}

func commonPrefix(a, b string) string {
	minLen := len(a)
	if len(b) < minLen {
		minLen = len(b)
	}

	i := 0
	for i < minLen && a[i] == b[i] {
		i++
	}

	return a[:i]
}

func main() {
	// Test cases
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"})) // "fl"
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))    // ""
}
