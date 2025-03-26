package check_if_n_and_its_double_exist

//Given an array arr of integers, check if there exist two indices i and j such that :

// i != j
// 0 <= i, j < arr.length
// arr[i] == 2 * arr[j]

// Input: arr = [10,2,5,3]
// Output: true
// Explanation: For i = 0 and j = 2, arr[i] == 10 == 2 * 5 == 2 * arr[j]
//var arr = []int{10, 2, 5, 3}

//Input: arr = [3,1,7,11]
//Output: false
//Explanation: There is no i and j that satisfy the conditions.
//var arr = []int{3,1,7,11}

// Input:
// [7,1,14,11]
// Output:
// false
// Expected:
// true
var arr = []int{7, 1, 14, 11}

func checkIfExist(arr []int) bool {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if i != j && arr[i] == 2*arr[j] {
				return true
			}
		}
	}
	return false
}
