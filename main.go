package main

import (
	"fmt"
)

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

func main() {
	fmt.Println(findDiagonalOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})) // Output: [1,2,4,7,5,3,6,8,9]
	fmt.Println(findDiagonalOrder([][]int{{1, 2}, {3, 4}}))                  // Output: [1,2,3,4// ]
}
