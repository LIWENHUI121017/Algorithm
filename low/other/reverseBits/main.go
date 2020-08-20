package main

import (
	"fmt"
)

func reverseBits(num uint32) uint32 {
	num = num<<16 | num>>16
	num = ((num & 0xff00ff00) >> 8) | ((num & 0x00ff00ff) << 8)
	num = ((num & 0xf0f0f0f0) >> 4) | ((num & 0x0f0f0f0f) << 4)
	num = ((num & 0xcccccccc) >> 2) | ((num & 0x33333333) << 2)
	num = ((num & 0xaaaaaaaa) >> 1) | ((num & 0x55555555) << 1)
	return num

	////直接从最后面开始循环，一个一个转换到前面
	//res := uint32(0)
	//p := uint32(31) //左移的位数，初始为31
	//for num != 0 {
	//	res += (num&1) << p
	//	num = num >> 1 //num位数-1
	//	p-- //左移的位数-1
	//}
	//return res

}

func main() {
	//15的二进制为1111
	num := uint32(15)
	res := reverseBits(num)
	fmt.Printf("%b", res)
}
