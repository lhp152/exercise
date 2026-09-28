package src

import "testing"

/*
260928
算法：

	数据结构：
	步骤：根据num长度，比如这里32位 就需要写明32个二进制位 字面值
		第一次 ：res = (0b01010101... & res)<<1 | (0b10101010... & res)>>1 // 0x555555..., 0xaaaaaa...
		。。。
		第5次，题目32位，2^5 = 32：res =  (0b00000001111111... & res)<<16 | (0b11111100000... & res) >> 16// 0x000..ff... 0xffff...0000...
	return res

边界：

	Constraints:
		• 0 <= n <= 231 - 2 // 不用考虑负数的补码等
		n is even.
*/
func TestReverseBits(t *testing.T) {
	tests := []struct {
		in, want int
	}{
		{0b00000010100101000001111010011100, 0b00111001011110000010100101000000},
		{43261596, 964176192},
		{0, 0},
		// {-1, -1}, // 全 1 翻转还是全 1， 题目边界条件 没有这个
	}
	for _, tt := range tests {
		if got := reverseBits(tt.in); got != tt.want {
			t.Errorf("reverseBits(%b) = %b, want %b", tt.in, got, tt.want)
		}
	}
}
func reverseBits(n int) int {
	res := n

	res = (0b01010101010101010101010101010101&res)<<1 | (0b10101010101010101010101010101010&res)>>1
	res = (0b00110011001100110011001100110011&res)<<2 | (0b11001100110011001100110011001100&res)>>2
	res = (0b00001111000011110000111100001111&res)<<4 | (0b11110000111100001111000011110000&res)>>4
	res = (0b00000000111111110000000011111111&res)<<8 | (0b11111111000000001111111100000000&res)>>8
	res = (0b00000000000000001111111111111111&res)<<16 | (0b11111111111111110000000000000000&res)>>16

	/*
		res = (res&0x55555555)<<1 | (res&0xAAAAAAAA)>>1
		res = (res&0x33333333)<<2 | (res&0xCCCCCCCC)>>2
		res = (res&0x0F0F0F0F)<<4 | (res&0xF0F0F0F0)>>4
		res = (res&0x00FF00FF)<<8 | (res&0xFF00FF00)>>8
		res = (res&0x0000FFFF)<<16 | (res&0xFFFF0000)>>16
	*/
	return res
}
