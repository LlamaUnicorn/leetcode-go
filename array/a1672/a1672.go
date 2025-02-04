package a1672

import (
	"fmt"
)

func MaximumWealth(accounts [][]int) int {
	maxWealth := 0
	for account := range accounts { // [[1,5] [7,3] [3,5]]
		fmt.Println(accounts[account])
		accountSum := 0
		for _, value := range accounts[account] { // [1,5]
			accountSum += value
		}
		fmt.Println(accountSum)
		fmt.Println(maxWealth)
		if accountSum > maxWealth {
			maxWealth = accountSum
		}
	}
	return maxWealth
}
