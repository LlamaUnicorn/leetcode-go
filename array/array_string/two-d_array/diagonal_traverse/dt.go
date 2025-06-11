package diagonal_traverse

//Given an m x n matrix mat, return an array of all the elements of the array in a diagonal order.
//
//Example 1:
//
//Input: mat = [[1,2,3],[4,5,6],[7,8,9]]
//Output: [1,2,4,7,5,3,6,8,9]
//Example 2:
//
//Input: mat = [[1,2],[3,4]]
//Output: [1,2,3,4]
//
//
//Constraints:
//
//m == mat.length
//n == mat[i].length
//1 <= m, n <= 104
//1 <= m * n <= 104
//-105 <= mat[i][j] <= 105

func findDiagonalOrder(mat [][]int) []int {
	m, n := len(mat), len(mat[0])
	result := make([]int, m*n)
	idx := 0

	// Process each diagonal (identified by sum i+j)
	for diag := 0; diag < m+n-1; diag++ {
		if diag%2 == 0 {
			// Even diagonal: traverse up-right (decreasing row, increasing col)
			for i := min(diag, m-1); i >= 0 && diag-i < n; i-- {
				j := diag - i
				result[idx] = mat[i][j]
				idx++
			}
		} else {
			// Odd diagonal: traverse down-left (increasing row, decreasing col)
			for i := max(0, diag-n+1); i <= diag && i < m; i++ {
				j := diag - i
				result[idx] = mat[i][j]
				idx++
			}
		}
	}

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//func main() {
//	fmt.Println(findDiagonalOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})) // Output: [1,2,4,7,5,3,6,8,9]
//	fmt.Println(findDiagonalOrder([][]int{{1, 2}, {3, 4}}))                  // Output: [1,2,3,4// ]
//}
