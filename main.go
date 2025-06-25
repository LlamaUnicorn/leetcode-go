package main

import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {
	//longest := ""
	//for _, word := range strs {
	//	fmt.Println(word)
	//	for i := 0; i < len(word); i++ {
	//		fmt.Println(string(word[i]))
	//	}
	//}
	for _, letter := range strs[0] {
		fmt.Println(string(letter))
		for i := 1; i < len(strs); i++ {
			fmt.Println(string(strs[i]))
			for j := 0; j < len(strs[i]); j++ {
				fmt.Println(string(strs[i][j]))
			}
		}
	}
	return ""
}

func main() {
	// Test cases
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"})) // "fl"
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))    // ""
}
