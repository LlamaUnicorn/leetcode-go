package main

import "fmt"

var arr = []int{1, 0, 2, 3, 0, 4, 5, 0} // [1,0,0,2,3,0,0,4]

func duplicateZeros(arr []int) {
	//fmt.Println(arr)
	for idx, elem := range arr {
		//fmt.Println(elem)
		if elem == 0 {
			for i := len(arr) - 1; i > idx; i-- {
				arr[i+1] = arr[i]
			}
		}
	}
	fmt.Println(arr)
}

func main() {
	duplicateZeros(arr)
}
