package src

import "fmt"

func TestInitSearchTree() {
	root := InitSearchTree()
	fmt.Println(root.Val)
}

func InitSearchTree() *TreeNode {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	var init func(nums []int, left, right int) *TreeNode
	init = func(nums []int, left, right int) *TreeNode {
		if left > right {
			return nil
		}
		mid := (left + right) / 2
		root := &TreeNode{
			Val: nums[mid],
		}
		root.Left = init(nums, left, mid-1)
		root.Right = init(nums, mid+1, right)
		return root
	}
	return init(nums, 0, len(nums)-1)
}
