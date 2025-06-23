package main

import (
	"fmt"
)

func strStr(haystack string, needle string) int {
	n := len(needle)
	if n == 0 {
		return 0
	}

	h := len(haystack)
	if h < n {
		return -1
	}

	for i := 0; i <= h-n; i++ {
		if haystack[i:i+n] == needle {
			return i
		}
	}

	return -1
}

func main() {
	// Test cases
	fmt.Println(strStr("sadbutsad", "sad"))  // 0
	fmt.Println(strStr("leetcode", "leeto")) // -1
}
