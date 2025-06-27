package reverse_string

//Write a function that reverses a string. The input string is given as an array of characters s.
//
//You must do this by modifying the input array in-place with O(1) extra memory.
//
//
//
//Example 1:
//
//Input: s = ["h","e","l","l","o"]
//Output: ["o","l","l","e","h"]
//Example 2:
//
//Input: s = ["H","a","n","n","a","h"]
//Output: ["h","a","n","n","a","H"]
//
//
//Constraints:
//
//1 <= s.length <= 105
//s[i] is a printable ascii character.
//Hide Hint #1
//The entire logic for reversing a string is based on using the opposite directional two-pointer approach!

func reverseString(s []string) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

//func main() {
//	// Test cases
//	strings := []string{"h", "e", "l", "l", "o"}
//	reverseString([]string{"h", "e", "l", "l", "o"}) // ["o","l","l","e","h"]
//	fmt.Println(strings)                             // ["o","l","l","e","h"]
//}
