package implement_string

//Implement strStr()
//
//Solution
//Given two strings needle and haystack, return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.
//
//
//Example 1:
//
//Input: haystack = "sadbutsad", needle = "sad"
//Output: 0
//Explanation: "sad" occurs at index 0 and 6.
//The first occurrence is at index 0, so we return 0.
//Example 2:
//
//Input: haystack = "leetcode", needle = "leeto"
//Output: -1
//Explanation: "leeto" did not occur in "leetcode", so we return -1.
//
//Constraints:
//
//1 <= haystack.length, needle.length <= 104
//haystack and needle consist of only lowercase English characters.

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

//func main() {
//	// Test cases
//	fmt.Println(strStr("sadbutsad", "sad"))  // 0
//	fmt.Println(strStr("leetcode", "leeto")) // -1
//}
