package src

// 0903 一遍过 手写heap
import (
	"fmt"
)

func TestHeap() {
	nums := []int{
		7, 5, 2, 6, 3, 1,
	}
	Initheap(nums)
	for i := 0; i < len(nums); i++ {
		fmt.Println(Popheap(nums))
	}
}

func Initheap(nums []int) {
	n := len(nums)
	for i := n / 2; i >= 0; i-- {
		ProcH(nums, i)
	}
}
func ProcH(nums []int, idx int) {
	largest := idx
	left := 2*largest + 1
	right := 2*largest + 2
	if left < len(nums) && nums[largest] < nums[left] {
		largest = left
	}
	if right < len(nums) && nums[largest] < nums[right] {
		largest = right
	}
	nums[idx], nums[largest] = nums[largest], nums[idx]
	if largest == idx { //即时退出 避免栈溢出
		return
	}
	ProcH(nums, largest)
}
func Popheap(nums []int) int {
	top := nums[0]
	nums[0] = nums[len(nums)-1]
	nums = nums[:len(nums)-1]
	ProcH(nums, 0)
	return top
}
