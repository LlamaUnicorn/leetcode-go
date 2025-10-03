// #two-pointers
//
// https://leetcode.com/problems/merge-strings-alternately/description/?envType=study-plan-v2&envId=leetcode-75
// You are given two strings word1 and word2. Merge the strings by adding letters in alternating order, starting with word1.
// If a string is longer than the other, append the additional letters onto the end of the merged string.

// Return the merged string.

// Example 1:

// Input: word1 = "abc", word2 = "pqr"
// Output: "apbqcr"
// Explanation: The merged string will be merged as so:
// word1:  a   b   c
// word2:    p   q   r
// merged: a p b q c r
// Example 2:

// Input: word1 = "ab", word2 = "pqrs"
// Output: "apbqrs"
// Explanation: Notice that as word2 is longer, "rs" is appended to the end.
// word1:  a   b
// word2:    p   q   r   s
// merged: a p b q   r   s
// Example 3:

// Input: word1 = "abcd", word2 = "pq"
// Output: "apbqcd"
// Explanation: Notice that as word1 is longer, "cd" is appended to the end.
// word1:  a   b   c   d
// word2:    p   q
// merged: a p b q c   d

// Constraints:

// 1 <= word1.length, word2.length <= 100
// word1 and word2 consist of lowercase English letters.

package s1768

import (
	"fmt"
	"log"
)

// mergeStrings loops over existing string chars
// until one of the strings ends.
// then it adds extra char from longer word that still has chars.
// Might be a bug for different length words. Indeed a bug.
func mergeStrings(word1 string, word2 string) string {
	i, j := 0, 0
	result := []byte{}
	for i < len(word1) && j < len(word2) {
		result = append(result, word1[i])
		result = append(result, word2[j])
		i++
		j++
	}

	if i < len(word1) {
		result = append(result, word1[i])
	}
	if j < len(word2) {
		result = append(result, word2[j])
	}
	return string(result)
}

func Run1768() {
	word1, word2 := "abc", "pqr"
	result := mergeStrings(word1, word2)
	expected := "apbqcr"
	if expected != result {
		log.Fatalf("Error! Expected %+v == %+v", expected, result)
	}
	fmt.Printf("Expected %+v, got %+v \n", expected, result)
	fmt.Println(expected == result)

	word3, word4 := "ab", "pqrs"
	result = mergeStrings(word3, word4)
	expected = "apbqrs"
	if expected != result {
		log.Fatalf("Error 2: Expected %+v == %+v", expected, result)
	}
	fmt.Printf("Expected %+v, got %+v \n", expected, result)
	fmt.Println(expected == result)
}
