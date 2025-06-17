package main

import (
	"fmt"
	"strings"
)

func addBinary(a string, b string) string {
	var result strings.Builder
	carry := 0
	i, j := len(a)-1, len(b)-1

	// Process from right to left
	for i >= 0 || j >= 0 || carry > 0 {
		sum := carry

		if i >= 0 {
			sum += int(a[i] - '0')
			i--
		}

		if j >= 0 {
			sum += int(b[j] - '0')
			j--
		}

		result.WriteByte(byte(sum%2) + '0')
		carry = sum / 2
	}

	// Reverse the result since we built it backwards
	s := result.String()
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func main() {
	// Test cases
	fmt.Println(addBinary("11", "1"))      // "100"
	fmt.Println(addBinary("1010", "1011")) // "10101"
}

// initiate sum
// if both are ones:
// save 0
// sum = 1
// iterate over the next, keeping the sum

//func addBinary(a string, b string) string {
//	for i := len(a) - 1; i >= 0; i-- {
//		fmt.Printf("a[%d] %c\n", i, a[i])
//		//fmt.Print(b[i])
//	}
//	return a + b
//}
//
//func main() {
//	fmt.Println(addBinary("11", "1")) // Output: 100
//	//fmt.Println(addBinary("1010", "1011")) // Output: 10101
//}
