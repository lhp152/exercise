package src

import "testing"

/*
0915
方法二 推荐：前缀和 + 哈希优化(更好)

	将前缀和及其频率信息保存在哈希表中，记录前缀和频率，方便检索;
	相对前缀和数组的优化：从每次遍历O(1)变成一次查询频率O(1)时间复杂度，这样一次就能检索出所有preSum - target的区间了。
	关键：前缀和频率，初始化和为0 的子串频率为1，目标公式preSum - target，不要记反了；
	算法
		数据结构
			前缀频率和及其频率 h : map[int]int， 初始化h[0] = 1,
			结果 res ： int
			presum：当前前缀和，每次都记录到h中
		步骤
			h[0] = 1
			大周期： 正序遍历 i nums[0, n- 1]
				presum += nums[i]
				检测前缀和频率中，是否存在某个子串的前缀和是preSum - target
				有则res += h[preSum-target]
				记录当前前缀和到hash中 ：h[preSum]++
				下一个大周期
			return res
	时间复杂度o(N)

空间复杂度O(N)
*/
func TestSubarraySum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{
			name: "示例1",
			nums: []int{1, 1, 1},
			k:    2,
			want: 2,
		},
		{
			name: "示例2",
			nums: []int{1, 2, 3},
			k:    3,
			want: 2,
		},
		{
			name: "包含负数",
			nums: []int{1, -1, 0},
			k:    0,
			want: 3,
		},
		{
			name: "单元素等于k",
			nums: []int{5},
			k:    5,
			want: 1,
		},
		{
			name: "无匹配",
			nums: []int{1, 2, 3},
			k:    7,
			want: 0,
		},
		{
			name: "全零",
			nums: []int{0, 0, 0},
			k:    0,
			want: 6,
		},
		{
			name: "从开头累加等于k",
			nums: []int{3, 4, 7, 2, -3, 1, 4, 2},
			k:    7,
			want: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := subarraySum(tt.nums, tt.k)
			if got != tt.want {
				t.Errorf("subarraySum(%v, %d) = %d, 期望 %d",
					tt.nums, tt.k, got, tt.want)
			}
		})
	}
}

func subarraySum(nums []int, k int) int {
	h := make(map[int]int)
	presum := 0
	res := 0
	h[0] = 1
	for i := range nums {
		presum += nums[i]
		if v, ok := h[presum-k]; ok {
			res += v
		}
		h[presum]++
	}
	return res
}
