package a412

import (
	"fmt"
	"strconv"
)

func FizzBuzz(n int) []string {
	fmt.Println(n)
	arr := make([]string, n)
	for i := 1; i <= n; i++ {
		arr = append(arr, strconv.Itoa(i))
	}
	fmt.Println(arr)
	//for _, i := range arr {
	//	switch {
	//	case i%3 == 0 && i%5 == 0:
	//		{
	//			arr[i] = "FizzBuzz"
	//		}
	//	case i%3 == 0:
	//		{
	//			arr[i] = "Fizz"
	//		}
	//	case i%5 == 0:
	//		{
	//			arr[i] = "Buzz"
	//		}
	//	default:
	//		arr[i] = strconv.Itoa(i)
	//	}
	//}
	return arr
}
