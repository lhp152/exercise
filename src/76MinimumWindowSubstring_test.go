package src

import "testing"

/*
	 0906
		算法：
			数据结构：
				targetMap统计目标字符串种类和个数：map[byte]int；
				windowMap:map[byte]int
				minWide 最小宽度 可初始化为len(s) + 1 表示没有
				窗口左右 边界 left, right: [left, right) 左闭右开
				初始化窗口条件：窗口左右边界都是0,，窗口长度是len(s) + 1；
			覆盖条件判断函数：（我这里优化了，但是对大周期有边界要求，如果这里再统计频率并对比频率，则大周期函数能简化）
				比对windowsmap和targetMap中字符频率 ，如果windowMap对应key的value小于，则不满足覆盖 return false
				满足覆盖则小于上一次宽度？
					记录一次子串
					更新最小宽度值
				return true
			大周期-滑动窗口遍历源字符串：
				外循环：right <= len(s)   // right是窗口开区间
					内循环：for 覆盖条件成立？//这个内循环是个优化点
							在windowMap中移除窗口左边界字符
							缩小窗口左边界
					检测right ！= len(s)
						在windowMap中增加窗口right边界字符
						增大窗口右边界
					否则break，不能扩展了// 这里其实打破了外循环边界条件，为了配合覆盖条件函数的优化
			返回结果：
				检测初始窗口边界是否变化？
					返回计算结果（注意不同编程语言，左右边界截取字符串时边界闭开条件，一般都是左闭右开）
				没变化

返回空结果
*/
func TestMinWindow(t *testing.T) {
	testset := []struct {
		s      string
		t      string
		expect string
	}{
		{
			s:      "ADOBECODEBANC",
			t:      "ABC",
			expect: "BANC",
		},
		{
			s:      "a",
			t:      "a",
			expect: "a",
		},
		{
			s:      "a",
			t:      "aa",
			expect: "",
		},
		{
			s:      "bba",
			t:      "ab",
			expect: "ba", // "bba" 中包含 "ab" 的最短子串是 "ba"
		},
	}

	for _, v := range testset {
		got := minWindow(v.s, v.t)
		if got != v.expect {
			t.Errorf("minWindow(%q, %q) = %q, want %q", v.s, v.t, got, v.expect)
		}
	}
}
func minWindow(s string, t string) string {
	hash_t := make(map[byte]int)
	for _, v := range []byte(t) {
		hash_t[v]++
	}
	hash_str := make(map[byte]int)
	res := ""

	minLen := len(s) + 1
	check := func(left, right int) bool {
		for i := range hash_t {
			if hash_t[i] > hash_str[i] {
				return false
			}
		}
		//覆盖并分析length
		if minLen > right-left {
			minLen = right - left
			res = s[left:right]
		}
		return true
	}

	left, right := 0, 0
	for right <= len(s) {
		for check(left, right) { //move left
			hash_str[s[left]]--
			left++
		}
		//move right
		if right == len(s) {
			break
		}
		hash_str[s[right]]++
		right++

	}
	return res
}
