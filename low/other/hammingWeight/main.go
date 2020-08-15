package main

import (
	"fmt"
)

//uint32就是无符号的int32类型
func hammingWeight(num uint32) int {
	//用1去和32位做与（&）运算，如果结果不是0的话 ，说明这个位置的数字就是1
	//count := 0
	//mask := uint32(1)
	//for i:=0; i<32;i++ {
	//	fmt.Println(num & mask)
	//	if (num & mask) != 0{
	//		count++
	//	}
	//	//每次循环后，1的位置就向左移动一位，就像是一个指针扫描一遍这32位数一样
	//	mask = mask << 1
	//}
	//return count

	//上面的方法，速度会稍慢些，因为每次比较都要比较32位，我们可以反向操作，将num向右移1一位，这样省了内存的同时，每次循环，比较的次数都会变少，因为num的位数每次都右移，也就是位数-1
	for num != 0 {
		if num&1 == 1 { //让num与1进行按位与运算，取得num最低位判断是否位1
			count++
		}
		num >>= 1 //num右移一位
	}
	return count

	//在二进制表示中，数字 n 中最低位的 1 总是对应 n−1 中的 0 。因此，将 n 和 n−1 与运算总是能把 n 中最低位的 1 变成 0 ，并保持其他位不变。
	//sum := 0
	//for num != 0 {
	//	sum++
	//	num &= (num - 1)
	//}
	//return sum

}

func main() {
	num := uint32(00000000000000000000000000001011)
	res := hammingWeight(num)
	fmt.Println(res)
}
