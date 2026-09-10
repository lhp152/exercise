package src

import (
	"math"
	"testing"
)

// ========== 测试用例 ==========
func TestMaxProfit(t *testing.T) {
	tests := []struct {
		name   string
		k      int
		prices []int
		want   int
	}{
		{
			name:   "示例: k=2, prices=[3,2,6,5,0,3]",
			k:      2,
			prices: []int{3, 2, 6, 5, 0, 3},
			want:   7,
		},
		{
			name:   "k=1, 单次交易",
			k:      1,
			prices: []int{7, 1, 5, 3, 6, 4},
			want:   5,
		},
		{
			name:   "k=2, 下降趋势",
			k:      2,
			prices: []int{5, 4, 3, 2, 1},
			want:   0,
		},
		{
			name:   "k=2, 上升趋势",
			k:      2,
			prices: []int{1, 2, 3, 4, 5},
			want:   4,
		},
		{
			name:   "k=0, 不能交易",
			k:      0,
			prices: []int{3, 2, 6, 5, 0, 3},
			want:   0,
		},
		{
			name:   "空数组",
			k:      2,
			prices: []int{},
			want:   0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxProfit(tt.k, tt.prices)
			if got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
func maxProfit(k int, prices []int) int {
	k = min(k, len(prices)/2) + 1
	sell := make([]int, k) //卖出 第j次交易的卖出行为，一次完整交易完成
	buy := make([]int, k)  //买入，第j 次交易的买入
	buy[0] = math.MinInt
	for i := range buy {
		buy[i] = math.MinInt // 保证第一次一定要买入
	}
	for i := range prices {
		for j := 1; j < k; j++ {
			buy[j] = max(sell[j-1]-prices[i], buy[j]) //buy[j] 是第j次完整交易的
			sell[j] = max(sell[j], buy[j]+prices[i])
		}
	}
	return sell[k-1]
}
