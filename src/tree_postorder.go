package src

// 0904
import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func TestPostOrder() {
	root := InitSearchTree()
	res := PostOrder(root)
	fmt.Println(res)
}
func PostOrder(root *TreeNode) []int {

	res := []int{}
	stack := []*TreeNode{}
	stack = append(stack, root)
	prev := &TreeNode{}
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		//pop
		if cur.Left == nil && cur.Right == nil { //当前节点是叶子节点，pop 更新prev
			res = append(res, cur.Val)
			stack = stack[:len(stack)-1]
			prev = cur
		} else if prev == cur.Left && cur.Right == nil { // 当前节点的left 遍历完，且right 空，pop 更新prev
			res = append(res, cur.Val)
			stack = stack[:len(stack)-1]
			prev = cur
		} else if prev == cur.Right { // 当前节点的right 遍历完， pop 更新prev
			res = append(res, cur.Val)
			stack = stack[:len(stack)-1]
			prev = cur
		} else { //push
			if cur.Right != nil {
				stack = append(stack, cur.Right)
			}
			if cur.Left != nil {
				stack = append(stack, cur.Left)
			}
		}
	}
	return res
}
