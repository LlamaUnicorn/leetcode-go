package longest_common_prefix

//Write a function to find the longest common prefix string amongst an array of strings.
//
//If there is no common prefix, return an empty string "".
//
//
//
//Example 1:
//
//Input: strs = ["flower","flow","flight"]
//Output: "fl"
//Example 2:
//
//Input: strs = ["dog","racecar","car"]
//Output: ""
//Explanation: There is no common prefix among the input strings.
//
//
//Constraints:
//
//1 <= strs.length <= 200
//0 <= strs[i].length <= 200
//strs[i] consists of only lowercase English letters if it is non-empty.

import (
	"fmt"
	//"strings"
)

//Use builtin function

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

// No built in functions
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
