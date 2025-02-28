package main

import (
	"fmt"
)

var arr = []int{1, 0, 2, 3, 0, 4, 5, 0} // [1,0,0,2,3,0,0,4]

func duplicateZeros(arr []int) {
	//fmt.Println(arr)
	for _, elem := range arr {
		fmt.Println(elem)
	}
}

func main() {
	duplicateZeros(arr)
}
