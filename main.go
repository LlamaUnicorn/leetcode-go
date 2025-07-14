package main

import (
	"fmt"
)

//func reverseWords(s string) string {
//	for _, word := range strings.Fields(s) {
//		fmt.Println(word)
//	}
//	return s
//}

//func reverseWords(s string) string {
//	n := len(s)
//	var words []string
//
//	// 1) Parse s manually, extract each word
//	i := 0
//	for i < n {
//		// Skip any leading spaces
//		for i < n && s[i] == ' ' {
//			i++
//		}
//		if i >= n {
//			break
//		}
//		// Mark the start of the word
//		start := i
//
//		// Advance until the next space or end
//		for i < n && s[i] != ' ' {
//			i++
//		}
//		// s[start:i] is one word
//		words = append(words, s[start:i])
//	}
//
//	// 2) Build the result by iterating words in reverse
//	var sb strings.Builder
//	for k := len(words) - 1; k >= 0; k-- {
//		sb.WriteString(words[k])
//		if k > 0 {
//			sb.WriteByte(' ')
//		}
//	}
//
//	return sb.String()
//}

func reverseWords(s string) string {
	// Convert to mutable buffer
	b := []byte(s)
	n := len(b)

	// 1) Trim leading/trailing spaces & collapse multiple spaces in-place
	write := 0
	read := 0

	// Skip initial spaces
	for read < n && b[read] == ' ' {
		read++
	}
	// Process until end
	for read < n {
		if b[read] == ' ' {
			// Write a single space, then skip all further spaces
			b[write] = ' '
			write++
			for read < n && b[read] == ' ' {
				read++
			}
		} else {
			// Copy non-space character
			b[write] = b[read]
			write++
			read++
		}
	}
	// Drop trailing space if any
	if write > 0 && b[write-1] == ' ' {
		write--
	}
	// Resize buffer to new length
	b = b[:write]

	// 2) Reverse the entire buffer
	reverse(b, 0, len(b)-1)

	// 3) Reverse each word within the buffer
	start := 0
	for i := 0; i <= len(b); i++ {
		if i == len(b) || b[i] == ' ' {
			reverse(b, start, i-1)
			start = i + 1
		}
	}

	return string(b)
}

// reverse swaps b[i] with b[j], moving inward
func reverse(b []byte, i, j int) {
	for i < j {
		b[i], b[j] = b[j], b[i]
		i++
		j--
	}
}

func main() {
	// Test cases
	fmt.Println(reverseWords("the sky is blue"))  // blue is sky the
	fmt.Println(reverseWords("  hello world  "))  // world hello
	fmt.Println(reverseWords("a good   example")) // example good a
}
