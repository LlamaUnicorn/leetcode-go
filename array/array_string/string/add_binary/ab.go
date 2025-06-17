package add_binary

//Given two binary strings a and b, return their sum as a binary string.
//
//
//Example 1:
//
//Input: a = "11", b = "1"
//Output: "100"
//Example 2:
//
//Input: a = "1010", b = "1011"
//Output: "10101"
//
//Constraints:
//
//1 <= a.length, b.length <= 104
//a and b consist only of '0' or '1' characters.
//Each string does not contain leading zeros except for the zero itself.

//import (
//	"fmt"
//	"strings"
//)
//
//func addBinary(a string, b string) string {
//	var result strings.Builder
//	carry := 0
//	i, j := len(a)-1, len(b)-1
//
//	// Process from right to left
//	for i >= 0 || j >= 0 || carry > 0 {
//		sum := carry
//
//		if i >= 0 {
//			sum += int(a[i] - '0')
//			i--
//		}
//
//		if j >= 0 {
//			sum += int(b[j] - '0')
//			j--
//		}
//
//		result.WriteByte(byte(sum%2) + '0')
//		carry = sum / 2
//	}
//
//	// Reverse the result since we built it backwards
//	s := result.String()
//	runes := []rune(s)
//	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
//		runes[i], runes[j] = runes[j], runes[i]
//	}
//
//	return string(runes)
//}
//
//func main() {
//	// Test cases
//	fmt.Println(addBinary("11", "1"))      // "100"
//	fmt.Println(addBinary("1010", "1011")) // "10101"
//}
