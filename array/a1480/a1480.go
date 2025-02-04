package a1480

// RunningSum calculates the running sum of the input array.
func RunningSum(nums []int) []int {
	var runSum []int
	var digit = 0
	for _, i := range nums {
		digit += i
		runSum = append(runSum, digit)
	}
	return runSum
}

//package main
//
//import "fmt"
//
//var nums = []int{1, 2, 3, 4}
//
//func runningSum(nums []int) []int {
//	var runSum []int
//	var digit = 0
//	for _, i := range nums {
//		fmt.Println(i)
//		digit += i
//		runSum = append(runSum, digit)
//	}
//	return runSum
//}
//
//func main() {
//	fmt.Println(runningSum(nums))
//}
