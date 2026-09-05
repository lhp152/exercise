package src

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 322 coinChange, 完全背包问题
func TestCoinChange(t *testing.T) {
	coins := []int{1, 2, 5}
	amount := 11

	if !assert.Equal(t, 3, coinChange(coins, amount)) {
		t.Error("wrong")
	}
}
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1) // need a sentinel
	dp[0] = 0                   //sentinel
	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt
	}
	for _, v := range coins {
		for j := v; j < amount+1; j++ {
			dp[j] = min(dp[j], dp[j-v]+1) //
		}
	}

	return dp[amount]
}
