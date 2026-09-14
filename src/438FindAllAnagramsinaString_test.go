package src

import (
	"reflect"
	"testing"
)

/*
260914
空间复杂度o(1), 时间复杂度O（n + m）
方法二 推荐：只用一个hash([26]int)保存模式和被检测字符串的字符频率差值，并使用diff来统计有几个字符的频率是不同的，注意不是频率差而是有哪几个字符的频率不同：

	空间复杂度o(1), 时间复杂度O（n + m）
	数据结构：
		diff： 频率不同的字符的个数，diff = 0 时，找到一个字母异位词，初始为0
		h[26]int: 在一个上面同时统计子串和目标字符串的频率，目标字符串字符++， 子串字符--，这样当h全部统计都是0时，两个串字母频率一致
		res []int: 保存符合字母异位词的子串的首元素下标
	初始化：
		同时对p和s[0: len(p)]的进行字符串频率统计，p的频率是正数增长，s的频率是负数增长(这里的增长方向决定了diff 的计算方式)；
		统计diff：如果初始窗口符合anagram，则[26]int的所有值都是0，刚好抵消；如果有一个不为0，则diff++，表示一个字符的频率不一致，但是不记录是哪个字符
		先比较diff，初始化窗口是否合适
	大周期：[1, len(s) - len(p) -1]移动并只统计发生变化的（从频率相等到不相等和从不相等到相等，一直相等或不相等的，并不会影响diff）：
		对s[len(p):]开始检测;并统计diff变化；
		左侧移除后：
			移除[i-1]元素，h对应位置--
			比较h[s[i-1]-'a'] ==0,说明有个字符频率一致了，diff--
			否则diff++
		右侧移入后：
			移入[i+len(p) -1]元素，h对应位置--
			比较h[s[i+len(p) -1]-'a'] ==0,说明有个字符频率一致了，diff--
			否则diff++
		diff == 0,说明当前子串是一个字母异位词，i 入res末尾

方法一我的算法:

	空间复杂度o(1), 时间复杂度O（26*n + m）

	数据结构
		SMap：s子串的字母频率 []int, 26
		PMap：p的全部字母频率 []int, 26
		结果集 res: []int 保存字母异位词的子串首元素坐标
		size 子串长度
	check: 同步遍历SMap PMap，保证两个数组的频率一致
	步骤
		初始化：
			PMap 统计p的字母频率
			SMap 统计s[0: size -1 ] 的区间频率，作为一个初始化

	大周期： i [0, len(s)-len(p)]  左右闭合区间，，i 表示当前窗口左边界，即结果下标
		// i == 0 时：sMap 已经是第一个窗口的统计结果，
		i > 0 时：
			左出：sMap[s[i-1]-'a']--
			右进：sMap[s[i+len(p)-1]-'a']++
		调用 check()，返回 true 则记录 i
*/
func TestFindAnagrams(t *testing.T) {
	tests := []struct {
		name string
		s    string
		p    string
		want []int
	}{
		{
			name: "示例1",
			s:    "cbaebabacd",
			p:    "abc",
			want: []int{0, 6},
		},
		{
			name: "示例2",
			s:    "abab",
			p:    "ab",
			want: []int{0, 1, 2},
		},
		{
			name: "s 比 p 短",
			s:    "a",
			p:    "ab",
			want: []int{},
		},
		{
			name: "完全相同",
			s:    "abc",
			p:    "abc",
			want: []int{0},
		},
		{
			name: "无匹配",
			s:    "abcdefg",
			p:    "xyz",
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findAnagrams(tt.s, tt.p)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findAnagrams(%q, %q) = %v, 期望 %v",
					tt.s, tt.p, got, tt.want)
			}
		})
	}
}

// 方法二 优化 推荐：使用diff 的方式来检测anagram
// 时间复杂度：O(n+m) 空间复杂度O(1)
func findAnagrams(s string, p string) []int {
	lenS := len(s)
	lenP := len(p)
	if lenS < lenP {
		return []int{}
	}
	res := []int{}
	diff := 0
	h := make([]int, 26)
	//初始化h和diff
	for i := range p {
		h[p[i]-'a']++
		h[s[i]-'a']--
	}
	for i := range h {
		if h[i] != 0 {
			diff++
		}
	}
	if diff == 0 {
		res = append(res, 0)
	}
	//移动窗口，我这里是(l, r]的方式
	for i := 1; i <= len(s)-len(p); i++ {
		//out of left boundary
		h[s[i-1]-'a']++
		if h[s[i-1]-'a'] == 0 {
			diff--
		} else {
			diff++
		}

		//new right boundary
		h[s[i+len(p)-1]-'a']--
		if h[s[i+len(p)-1]-'a'] == 0 {
			diff--
		} else {
			diff++
		}
		if diff == 0 {
			res = append(res, i)
		}
	}
	return res
}

// 我的
func findAnagrams_my(s string, p string) []int {
	if len(p) > len(s) {
		return []int{}
	}
	res := []int{}
	pMap := make([]int, 26)
	for i := range p {
		pMap[p[i]-'a']++
	}
	sMap := make([]int, 26)
	check := func() bool {
		for i := range pMap {
			if sMap[i] != pMap[i] {
				return false
			}
		}
		return true
	}
	for i := range s[0:len(p)] {
		sMap[s[i]-'a']++
	}
	for i := 0; i <= len(s)-len(p); i++ {
		if i > 0 {
			sMap[s[i-1]-'a']--
			sMap[s[i+len(p)-1]-'a']++
		}

		if check() {
			res = append(res, i)
		}

	}
	return res
}
