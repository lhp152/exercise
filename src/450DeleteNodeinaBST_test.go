package src

import (
	"reflect"
	"testing"
)

/*
0908

算法：

	数据结构：
		successor： 临时变量，可以是left 子树的最右侧节点，或right 子树的最左侧节点，解法不同
	步骤：

递归删除函数 deleteNode(root, key): key 是目标值，返回一个节点，表示用来替换的节点；

	如果 root == nil:
		返回 nil

	如果 key < root.Val:
		root.Left = deleteNode(root.Left, key)
		返回 root

	如果 key > root.Val:
		root.Right = deleteNode(root.Right, key)
		返回 root

	//说明root.Val == key // 至此，找到了要删除的节点 root

	//找到root，现在找successor，有三种情况（或者是两种，一侧为空；两侧都为空）
	场景1：如果 root.Left == nil:
		返回 root.Right           // 右子树直接顶替root，可以不考虑root.Right 是否也是空

	场景2：如果 root.Right == nil:
		返回 root.Left            // 左子树直接顶替root

	场景3：左右子树都存在，左子树最右侧非空节点或者右子树最左侧非空
		以右子树最左侧非空子节点为例
		1. 迭代遍历找后继节点：
			successor = root.Right
			while successor.Left != nil:
				successor = successor.Left

		2. 复制值覆盖：//这里官方也是值替换，没有用指针
			root.Val = successor.Val

		3. 删除右子树中的后继节点：
		复用删除函数,将目标顶替节点从右子树中删除;这里复用是因为successor的删除还是需要考虑很多情况，比如要记住successor 的父节点，再比如successor 可能就是root.Right,所以直接复用最优
			root.Right = deleteNode(root.Right, successor.Val)

	4. 返回 root
*/
func TestDeleteNode(t *testing.T) {
	tests := []struct {
		name     string
		rootArr  []interface{}
		key      int
		expected []interface{}
	}{
		{
			name:     "LeetCode 450 示例",
			rootArr:  []interface{}{5, 3, 6, 2, 4, nil, 7},
			key:      3,
			expected: []interface{}{5, 4, 6, 2, nil, nil, 7},
		},
		// 可以继续添加更多测试用例
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := arrayToTree(tt.rootArr)
			result := deleteNode(root, tt.key)
			got := treeToArray(result)

			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("deleteNode() = %v, want %v", got, tt.expected)
			}
		})
	}
}

// 输入树和key，返回删除节点后的新树的root
func deleteNode(root *TreeNode, key int) *TreeNode {

	//search
	//delete
	if root == nil {
		return nil
	}

	if root.Val < key { //目标节点在右侧，且可能就是right
		root.Right = deleteNode(root.Right, key)
		return root
	} else if root.Val > key {
		root.Left = deleteNode(root.Left, key)
		return root
	}
	// 找到目标删除节点 root.Val == key

	// 基本删除原则：find the most right node in the left or the most left node in the right；
	// 两种情况 1 一侧是nil，则直接用另一侧非空子节点顶替当前root位置； 2 两侧都为为空，则需要用基本删除原则，找到目标节点，并使用删除函数，递归删除掉这个目标顶替节点
	//第一种 一侧为空，另一侧非空，则非空子节点就是最好的，是剩余数据中最接近root 的节点
	if root.Left == nil {
		return root.Right // 右子树直接顶替root，可以不考虑root.Right 是否也是空
	} else if root.Right == nil {
		return root.Left // 左子树直接顶替root
	}
	// 第二种情况，可以在左或用中找继任，我这里是右侧的最左侧节点
	successor := root.Right
	for successor.Left != nil {
		successor = successor.Left
	}
	root.Val = successor.Val //这里官方也是值替换，没有用指针
	//复用删除函数,将目标顶替节点从右子树中删除;这里复用是因为successor的删除还是需要考虑很多情况，比如要记住successor 的父节点，再比如successor 可能就是root.Right,所以直接复用最优
	root.Right = deleteNode(root.Right, successor.Val)
	return root
}

// ============ 辅助函数：数组 ↔ 二叉树 ============

// arrayToTree 将层序遍历数组转为二叉树，null 用 nil 表示
func arrayToTree(arr []interface{}) *TreeNode {
	if len(arr) == 0 || arr[0] == nil {
		return nil
	}

	root := &TreeNode{Val: arr[0].(int)}
	queue := []*TreeNode{root}
	i := 1

	for len(queue) > 0 && i < len(arr) {
		node := queue[0]
		queue = queue[1:]

		// 左子节点
		if i < len(arr) && arr[i] != nil {
			node.Left = &TreeNode{Val: arr[i].(int)}
			queue = append(queue, node.Left)
		}
		i++

		// 右子节点
		if i < len(arr) && arr[i] != nil {
			node.Right = &TreeNode{Val: arr[i].(int)}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}

// treeToArray 将二叉树转为层序遍历数组（包含 null）
func treeToArray(root *TreeNode) []interface{} {
	if root == nil {
		return []interface{}{}
	}

	res := []interface{}{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node != nil {
			res = append(res, node.Val)
			queue = append(queue, node.Left, node.Right)
		} else {
			res = append(res, nil)
		}
	}

	// 去掉末尾的 nil，保持紧凑
	for len(res) > 0 && res[len(res)-1] == nil {
		res = res[:len(res)-1]
	}

	return res
}
