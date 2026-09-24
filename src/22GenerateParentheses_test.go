package src

import "testing"

/*
算法：

	数据结构：
		左右括号数量 nums:[2] int, 初始化 [0] = n, [1] = n;
		res : []string
		tmp := []byte
		层数：level， 初始化为0, [0,2*n - 1]；
	dfs：输入level（能进入dfs 说明之前合法）
		level >= 2 *n,， 一个合法括号组合, tmp 转为string，入res；
		继续
		for range 2
			nums[i]--
			if num[i] >= 0 && nums[0] <= nums[1] ,合法，进入下一层
				tmp 根据i 收集一种括号
				dfs(level+1)
				nums[i]++，tmp 剔除末尾元素， 恢复环境
			否则不合法，下一个周期；
		return

返回res，就是所有合法的括号，
*/
func TestGeneratePatenthesis(t *testing.T) {
	got := generateParenthesis(3)
	want := []string{"((()))", "(()())", "(())()", "()(())", "()()()"}
	if len(got) != len(want) {
		t.Fatalf("got %v, want %v", got, want)
	}
	// 顺序可能不同，用 map 比对
	m := map[string]bool{}
	for _, s := range got {
		m[s] = true
	}
	for _, s := range want {
		if !m[s] {
			t.Fatalf("missing %s in %v", s, got)
		}
	}
}
func generateParenthesis(n int) []string {
	res := []string{}
	tmp := []byte{}
	var sources [2]int
	sources[0] = n
	sources[1] = n
	var dfs func(level int)
	dfs = func(level int) {
		if level >= 2*n {
			res = append(res, string(tmp))
			return
		}
		for i := range 2 {
			sources[i]--
			if sources[i] >= 0 && sources[0] <= sources[1] {
				if i == 0 {
					tmp = append(tmp, '(')
				} else {
					tmp = append(tmp, ')')
				}
				dfs(level + 1)
				tmp = tmp[:len(tmp)-1]
			}
			sources[i]++
		}
	}
	dfs(0)
	return res
}
