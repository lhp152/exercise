package src

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
