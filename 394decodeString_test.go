package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
