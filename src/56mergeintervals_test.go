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
func merge() {

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
		for i+1 < len(data) && tmp[1]+1 >= data[i+1][0] { //临接 处理为合并, 30 31
			tmp[1] = max(tmp[1], data[i+1][1])
			i++
		}
		i++ //合并的是i+1，不是i！！！所以下次tmp 要迁移到i+1
		res = append(res, tmp)

	}
	fmt.Println(res)
}
