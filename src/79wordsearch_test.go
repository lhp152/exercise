package src

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
0906

	dfs算法：
		数据结构
			遍历过数组 visited [][]int或者用map
			四个方向direction 数组：加速dfs 遍历

		步骤
			递归检测函数： 输入位置 i,j，要找的字符串中字符下标
				i,j 边界检测
				idx 边界检测
					越界 return true 说明找到一个
				是否visited遍历过
				[i][j] 与string[idx] 是否相同，不同则return false
				相同:记录visited；并4个方向找 下一个
					递归检测函数，输入i j 与direction的计算，idx+1
					检测递归检测函数返回值，如果true ，则return true 不再递归
				恢复visited
				return false
			检测外循环i [0, n-1]
				检测内循环 j [0, m - 1]
					递归检测函数：输入i,j ,0；
					检测函数返回值，true 则直接返回true

return false
*/
func TestExist(t *testing.T) {
	testset := []struct {
		board [][]byte
		word  string
		want  bool
	}{
		{
			board: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			word: "ABCCED",
			want: true,
		},
		{
			board: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			word: "SEE",
			want: true,
		},
		{
			board: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			word: "ABCB",
			want: false,
		},
	}
	for _, v := range testset {
		assert.Equal(t, v.want, exist(v.board, v.word))
	}
}

func exist(board [][]byte, word string) bool {
	var visited [][]bool
	for i := range board {
		visited = append(visited, make([]bool, len(board[i])))
	}

	direction := [][]int{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}
	var check func(i, j, idx int) bool
	check = func(i, j, idx int) bool {

		if idx >= len(word) {
			return true
		}
		if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) {
			return false
		}
		if visited[i][j] || board[i][j] != word[idx] { //检测是否走过，是否是目标字符
			return false
		}
		visited[i][j] = true // 记录，来时路
		for _, v := range direction {
			if check(i+v[0], j+v[1], idx+1) { //找到一个即可返回
				return true
			}
		}
		visited[i][j] = false // 还原

		return false // 说明遍历完没有找到
	}
	for i, r := range board {
		for j := range r {
			if check(i, j, 0) { //找到一个即可退出
				return true
			}
		}
	}
	return false
}
