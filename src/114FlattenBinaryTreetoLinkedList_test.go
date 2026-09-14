package src

import (
	"fmt"
	"reflect"
	"testing"
)

/*
260914

关键：算法有很多种，我这里展示第三种，都是O(n)的时间复杂度，只有第三种是O(1)的空间复杂度。

自顶向下迭代（推荐）算法：

	空间复杂度是时间复杂度是O(n)！！ 经常考察的点
		虽然有两个for循环，但是每个节点在内循环中只访问一次，因为内循环只遍历root.Left 的右侧节点来找到最右侧非空节点，阶段性展开时，这些遍历过的节点都在root.Right 上，不会再重复遍历，所以最终时间复杂度是外循环的n次，因为外循环是每个节点都看一次。
	数据结构：
		tail：root的left 子树最右侧非空节点
	大周期：输入局部root，非 nil，否则说明处理完了，直接返回；
		root.left 不为空，则进行一次迭代小周期：
			迭代小周期：输入root
				root.left子树中，迭代找最右侧非空叶子节点，更新tail为这个节点；
				将left子树嵌入其root和right兄弟节点之间：
					tmp = root.right//断开root 右子树准备拼接
					root.right = root.left
					删除root左子树(必须)：root.left = nil，
					tail.right 指向tmp, (将root.left子树嵌入到root和root.right之间）
			root = root.Right， 移动root，进入下一大周期
		left为空，说明left子树不用展开

移动root：root = root.right
*/
func TestFlatten(t *testing.T) {
	//        1
	//       / \
	//      2   5
	//     / \   \
	//    3   4   6
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 3},
			Right: &TreeNode{Val: 4},
		},
		Right: &TreeNode{
			Val:   5,
			Right: &TreeNode{Val: 6},
		},
	}

	flatten(root)

	want := []int{1, 2, 3, 4, 5, 6}
	var got []int
	for node := root; node != nil; node = node.Right {
		got = append(got, node.Val)
		if node.Left != nil {
			t.Errorf("节点 %d 的 Left 不为 nil", node.Val)
		}
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("flatten 结果 = %v, 期望 %v", got, want)
	}
}
func flatten(root *TreeNode) {
	curr := root
	for curr != nil { //外层遍历所有节点
		if curr.Left != nil {
			tail := curr.Left
			for tail.Right != nil { //内层每次指遍历left 子树的right 侧节点；left 的right 子树在拼接后，不会被后序内循环遍历到！！
				tail = tail.Right
			}
			tail.Right = curr.Right
			curr.Right = curr.Left
			curr.Left = nil
		}
		curr = curr.Right
	}
	fmt.Println(root)
}
