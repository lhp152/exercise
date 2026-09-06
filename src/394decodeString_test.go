package src

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
	遍历数组时，每个字符处理方式：使用[]byte（[]string 会不一样）
		数字：累积数字 cur_num = cur_num * 10 + num
		字母：累加成串
		[ ：临时数字入栈，将"["入结果栈；
		] ：出结果栈和数字, 数字栈 出栈一个数字，字符栈将[和一个完整字符串都出栈，用数字完成字符串倍数膨胀后入栈；（注意golang slice 引用类型特性，很多要深拷贝和清空操作）
		结束时：追加临时字符，是一个边界条件。
	算法：
		数据结构：
			字符结果栈，保存字母和“[”：res ： []byte{} 这里byte 比string方便，不用拼接string
			数字栈，num_stack：， []int {} ，保存转型后的数字；
			cur_str：[]byte，原始字符串中，当前连续字符；
			cur_num： 将连续数字拼接到一起,并转为int，
			其他参考：
				numStack := []int{}// 数字栈

strStack := []string{}//字符串栈，只保存字符串，不保存[
curStr := "" 结果
curNum := 0

	大周期：遍历整个数组，
		switch 检测
			数字：当前数字字符与cur_num 结合为新cur_num：cur_num = cur_num * 10 + num（先转型）
			字符：
				append到cur_str字符数组
			"[";
				说明前面是一个数字
				cur_num入数字栈；
				"[" 入栈
			"]"：
				进入小周期，两个栈都出栈并计算
				小周期-字符结果栈出栈，数字栈出栈+入栈：
					top+pop 字符栈，直到遇到[,可以最后截取，比如[index:]的方式，直接pop头插字符串拼接并不方便；
					top+pop 数字栈，得到[前的数字；
					将pop的字符串迭代相应top次数后入结果栈res；
					退出小周期；
	append  cur_str到res,不用检测；

遍历结束：res 栈转为string，就是结果
*/
func TestDecodeString(t *testing.T) {
	testset := []struct {
		s      string
		expect string
	}{
		{s: "3[a]2[bc]", expect: "aaabcbc"},
		{s: "3[a2[c]]", expect: "accaccacc"},
		{s: "2[abc]3[cd]ef", expect: "abcabccdcdcdef"},
		{s: "abc3[cd]xyz", expect: "abccdcdcdxyz"},
	}
	for _, v := range testset {
		assert.Equal(t, v.expect, decodeString(v.s))

	}
}

// 递归好做，迭代难
// []byte{}
func decodeString(s string) string {
	res := []byte{}
	nums := []int{}
	cur_num := 0
	cur_str := []byte{}

	for _, v := range []byte(s) {
		switch v {
		case '[':
			nums = append(nums, cur_num) //之前一定有一个数字
			cur_num = 0
			cur_str = append(cur_str, v) //入栈左括号
		case ']':
			res = append(res, cur_str...)
			cur_str = []byte{}
			for i := len(res) - 1; i >= 0; i-- {
				if res[i] == '[' {
					cur_str = make([]byte, len(res)-i-1) // 要深拷贝，避免slice 引用问题
					copy(cur_str, res[i+1:])
					res = res[:i] //出栈
					break
				}
			}
			cur_num = nums[len(nums)-1]
			nums = nums[:len(nums)-1]
			cur_str = []byte(strings.Repeat(string(cur_str), cur_num))
			res = append(res, cur_str...)
			cur_num = 0
			cur_str = []byte{}
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			cur_num = cur_num*10 + int(v-'0')
		default: // str
			cur_str = append(cur_str, v)
		}
	}
	res = append(res, cur_str...)
	return string(res)
}

//[]string{}
