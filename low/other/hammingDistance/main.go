package main

import (
	"fmt"
)

func hammingDistance(x int, y int) int {
	//count := 0
	//两个int每一位都和1与运算，如果结果不相等，计数加一
	//for x != y  {
	//	if x&1 != y&1 {
	//		count++
	//	}
	//	//然后往右移
	//	x >>= 1
	//	y >>= 1
	//}

	//利用异或的特性，1^1=0 1^0=1,例如1和4对比，他们的的二进制是：001 和 100，异或之后变成：011，1的数量有2位，也就代表了两个数字的二进制不相同的位置总共有两个，所以接下来我们只需要计算两个数异或后的数的二进制中，1的数量就是答案了
	//count := 0
	//num := x^y
	//for num!=0 {
	//	if num&1 == 1 {
	//		count++
	//	}
	//	num >>= 1
	//}

	//计数这个异或的时候可以用之前求位1的个数里面的方法，直接用n&n-1,就能把最右边1去掉，当所有的1都去掉了，n就会变成0，n&n-运算的次数就是异或后的二进制里面1的数量。
	//这是布赖恩·克尼根位计数算法的基本思想。该算法使用特定比特位和算术运算移除等于 1 的最右比特位。
	count := 0
	num := x ^ y
	for num != 0 {
		count++
		num = num & (num - 1)
	}
	return count
}

func main() {
	a := 1
	b := 4
	res := hammingDistance(a, b)
	fmt.Println(res)
}
