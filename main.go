package main

import (
	"fmt"
)

func reverseString(s []string) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	// Test cases
	strings := []string{"h", "e", "l", "l", "o"}
	reverseString([]string{"h", "e", "l", "l", "o"}) // ["o","l","l","e","h"]
	fmt.Println(strings)                             // ["o","l","l","e","h"]
	//fmt.Println(reverseString([]string{"H", "a", "n", "n", "a", "h"})) // ["h","a","n","n","a","H"]
}
