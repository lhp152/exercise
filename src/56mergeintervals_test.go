package src

import (
	"fmt"
	"sort"
)

/*
0910 百度面试题，将边界改了
56.合并区间 merge intervals 原题
算法题：合并服务异常时间段

SRE平台会从不同监控系统收到服务异常时间段。由于报警可能存在重叠，需要将所有重叠或相邻的异常时间段合并。

题目描述

给定一组无序时间区间 intervals，其中：
intervals[i] = []int{start, end}
表示服务从 start 到 end 处于异常状态。

请合并所有重叠或相邻的异常区间，并按开始时间升序返回。
示例
输入：

	[][]int{
	    {10, 20},
	    {15, 30},
	    {40, 50},
	    {31, 35},
	    {51, 60},
	}

输出：

	[][]int{
	    {10, 35},
	    {40, 60},
	}
*/
/*
两个思路：
	如下是第一种（推荐），一个loop，每次更新cur，index[1, n- 1] 移动；最后一个cur要追加，因为剩余
	第二种：两个loop（我在0909百度面试中想出来的），index[0, n-1] 每次比对[index+ 1]且index+1 合法，不重叠则合并，这个不

	区别：第二个算法，i+1 就是下一个对比区间，上个算法，i 就是下一个对比区间
算法：
	数据结构：
		res：[][2]int二维整型数组，可以直接确定内部是[2]int
		cur：[2]int当前要处理的区间元素，一维数组；
	排序：按照区间的起点/左端点，升序排序
		sort.Slice（，function）
		输入二维数组，两个相邻元素下标，这里会用slices库的排序函数
	初始化cur = arr[0]
	大周期（合并）：当前区间元素cur和要比较的区间元素下标index 区间 [1, n - 1]
		判断不重叠，追加到答案：
			检测前区间的右边界 > 后区间的左边界（这里用了排除法）
			将当前区间cur 追加到到答案里；
			更新cur = arr[index]//下一个
		否则合并
			更新cur右边界 = max(cur右边界，data[index]右边界)；
		index ++，并进入下一个大周期；//
	追加最后一个cur到res
返回res
*/
// 推荐一个loop
func merge1() {
	fmt.Println("ok")
	data := [][]int{
		{10, 20},
		{15, 30},
		{31, 35},
		{40, 50},
		{51, 60},
	}
	sort := func(nums [][]int) {
		sort.Slice(nums, func(i, j int) bool {
			return nums[i][0] < nums[j][0]
		})
	}
	sort(data)
	fmt.Println(data)
	//[[10 20] [15 30] [31 35] [40 50] [51 60]]
	res := [][]int{}
	cur := data[0]
	i := 1
	for i < len(data) {
		//找重叠/相邻
		if cur[1]+1 >= data[i][0] { //临接 处理为合并, 30 31
			cur[1] = max(cur[1], data[i][1])
		} else { //追加并更新cur
			res = append(res, cur)
			cur = data[i] //用不重叠的作为下一个目标值
		}
		i++
	}
	res = append(res, cur) //data中剩余最后一个cur区间会越界与 i= n+ 1 对比，导致cur没有合并，这里要补偿
	fmt.Println(res)
}

// 百度面试想到的，不如1好
func merge2() {
	fmt.Println("ok")
	data := [][]int{
		{10, 20},
		{15, 30},
		{31, 35},
		{40, 50},
		{51, 60},
	}
	sort := func(nums [][]int) {
		sort.Slice(nums, func(i, j int) bool {
			return nums[i][0] < nums[j][0]
		})
	}
	sort(data)
	fmt.Println(data)
	//[[10 20] [15 30] [31 35] [40 50] [51 60]]
	res := [][]int{}
	tmp := []int{}
	i := 0
	for i < len(data) {
		tmp = data[i]

		//找重叠/相邻
		for i+1 < len(data) && tmp[1]+1 >= data[i+1][0] { //临接 处理为合并, 30 31，这个算法i+1 就是下一个对比区间，上个算法，i 就是下一个对比区间
			tmp[1] = max(tmp[1], data[i+1][1])
			i++
		}
		i++ //合并的是i+1，不是i！！！所以下次tmp 要迁移到i+1
		res = append(res, tmp)

	}
	fmt.Println(res)
}
