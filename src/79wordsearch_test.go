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
