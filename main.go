package main

import (
	"fmt"
)

// initiate sum
// if both are ones:
// save 0
// sum = 1
// iterate over the next, keeping the sum

func addBinary(a string, b string) string {
	for i := len(a) - 1; i >= 0; i-- {
		fmt.Printf("a[%d] %c\n", i, a[i])
		//fmt.Print(b[i])
	}
	return a + b
}

func main() {
	fmt.Println(addBinary("11", "1")) // Output: 100
	//fmt.Println(addBinary("1010", "1011")) // Output: 10101
}
