package a1672

import "fmt"

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

// Accepted, Runtime: 0 ms, Memory Usage: 3.3 MB
// Runtime: 0 ms, faster than 100.00% of Go online submissions for Richest Customer Wealth.
func maximumWealth(accounts [][]int) (max int) {
	calculateWealth := func(account []int, sumCh chan int) {
		sum := 0
		for _, val := range account {
			sum += val
		}
		sumCh <- sum
	}

	for _, account := range accounts {
		sumCh := make(chan int)
		go calculateWealth(account, sumCh)
		accountWealth := <-sumCh
		if accountWealth > max {
			max = accountWealth
		}
		close(sumCh)
	}
	return
}
