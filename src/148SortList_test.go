package src

import (
	"reflect"
	"testing"
)

/*
0906
链表排序：归并/插入/快排都可，各有注意事项
快排：

	链表的拆分和合并都有注意点
	拆分时，注意两个分开链表的tail 可能有关联 要断开
	合并时，拼接而不是归并的合并两个有序链表

归并：

插入算法：双指针
基本思想：用dummy 保存结果链表的头节点；双指针，一个tail 一个pre，pre找到剩余链表中合适的节点将这个节点插入到tail尾部,移动tail; 开始新遍历，pre从tail 开始新遍历，pre.Next ！= nil ；直到tail.Next = nil 结束

	注意next是下一个目标节点的pre

算法：可以将排序好的和剩余链表分开，也可以部分 两种策略而已; 这里是未断开
数据结构：

	有序链表的dummy:
	有序链表的tail:
	遍历剩余链表pre: 每次初始化为tail，因为单向链表 要得到目标节点的前一个节点
	剩余链表的最小节点 presmallest: 每次初始化为tail

步骤：

	初始化dummy tail节点
	大周期：输入一个新剩余链表，tail != nil 且tail.Next != nil
		初始化pre和presmallest = tail
		小周期：遍历剩余链表，找到最小节点的前置节点， pre.Next != nil
			pre.Next.Val presmallest.Next.Val 对比
				如果pre 更小则更新presmallest为pre
			pre= pre.Next 移动pre
		头插法，提取最小节点并插入到有序链表后侧
			next 等于presmallest.Next
			presmallest.Next = presmallest.Next.Next

			//插入tail后面
			next.Next = tail.next
			tail.Next = next
			tail = tail.Next
		下一个大周期

return dummy.Next
*/
func TestSortList(t *testing.T) {
	// 辅助函数：数组转链表
	buildList := func(nums []int) *ListNode {
		dummy := &ListNode{}
		cur := dummy
		for _, v := range nums {
			cur.Next = &ListNode{Val: v}
			cur = cur.Next
		}
		return dummy.Next
	}

	// 辅助函数：链表转数组（便于比较）
	listToSlice := func(head *ListNode) []int {
		res := []int{}
		for head != nil {
			res = append(res, head.Val)
			head = head.Next
		}
		return res
	}

	testset := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "正常乱序",
			input: []int{4, 2, 1, 3},
			want:  []int{1, 2, 3, 4},
		},
		{
			name:  "已经有序",
			input: []int{1, 2, 3, 4},
			want:  []int{1, 2, 3, 4},
		},
		{
			name:  "逆序",
			input: []int{5, 4, 3, 2, 1},
			want:  []int{1, 2, 3, 4, 5},
		},
		{
			name:  "空链表",
			input: []int{},
			want:  []int{},
		},
		{
			name:  "单个节点",
			input: []int{1},
			want:  []int{1},
		},
		{
			name:  "有重复值",
			input: []int{3, 1, 2, 3, 1},
			want:  []int{1, 1, 2, 3, 3},
		},
	}

	for _, tt := range testset {
		t.Run(tt.name, func(t *testing.T) {
			head := buildList(tt.input)
			sorted := sortListInert(head) // 调用被测函数
			got := listToSlice(sorted)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortList(%v) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 快排 (快拍和归并都用到了分治思想 不要混了)
// 注意链表根据pivot 拆分时，最后两个节点要断开，避免循环引用或链表相交
func sortList(head *ListNode) *ListNode {
	var sort func(head *ListNode) (root *ListNode)
	sort = func(head *ListNode) *ListNode {
		if head == nil {
			return nil
		}
		pivot := head
		head = head.Next
		pivot.Next = nil
		right := &ListNode{}
		dummy_r := right
		left := &ListNode{}
		dummy_l := left
		for head != nil { //链表分离，注意将head 先摘出来，避免最后一位循环引用或链表相交，也就是结束时left 和right 的 Next都是nil
			// next := head.Next
			// head.Next = nil
			if head.Val < pivot.Val {
				left.Next = head
				left = left.Next

			} else {
				right.Next = head
				right = right.Next
			}
			// head = next
			head = head.Next
		} //这里退出时，最后一个移动的left/right，其末尾节点会被另一个链表尾部也指向，所以for 内部每次将目标节点先断开
		right.Next = nil //要么每次断开head，要么最后将right left 与后序节点断开
		left.Next = nil
		//
		left = sort(dummy_l.Next)
		right = sort(dummy_r.Next)

		//merge
		pivot.Next = right //记得合并pivot
		if left == nil {
			return pivot
		}
		dummy := left
		for left.Next != nil { //有环
			left = left.Next
		}
		left.Next = pivot

		return dummy
	}
	return sort(head)
}

// 链表的插入排序，连续
func sortListInert(head *ListNode) *ListNode {
	dummy := &ListNode{}
	tail := dummy // tail 闭区间
	tail.Next = head
	for tail != nil && tail.Next != nil { // 边界，tail 是最后一个非空节点时退出
		pre := tail
		presmallest := tail
		for pre.Next != nil {
			if pre.Next.Val < presmallest.Next.Val {
				presmallest = pre
			}
			pre = pre.Next
		}
		//头插法
		// 提取目标节点
		next := presmallest.Next
		presmallest.Next = presmallest.Next.Next
		// 插入tail 后面并移动tail
		next.Next = tail.Next
		tail.Next = next
		tail = tail.Next
	}
	return dummy.Next
}

// 链表插入排序优化，断开；其实代码差不多
func sortListInert_good(head *ListNode) *ListNode {
	dummy := &ListNode{}
	tail := dummy // tail 闭区间

	// smallest 从链表中得到最小节点，以及删除改节点后的链表的head
	smallest := func(head *ListNode) (newhead, target *ListNode) {
		if head == nil {
			return nil, nil
		}
		dummy := &ListNode{}
		dummy.Next = head
		presmallest := dummy
		pre := dummy
		for pre.Next != nil {
			if pre.Next.Val < presmallest.Next.Val {
				presmallest = pre
			}
			pre = pre.Next

		}
		//提取节点， 优化后，这里就不用判断 presmallest.Next 的边界问题，这个问题在函数入口就规避了
		target = presmallest.Next
		presmallest.Next = presmallest.Next.Next
		return dummy.Next, target
	}
	for head != nil {
		head, tail.Next = smallest(head)
		tail = tail.Next
	}
	return dummy.Next
}
