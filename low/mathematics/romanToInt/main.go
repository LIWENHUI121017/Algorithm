package main

import "fmt"

/**
*执行结果：通过
*执行用时：12 ms, 在所有 Go 提交中击败了47.29%的用户
*内存消耗：3.1 MB, 在所有 Go 提交中击败了100.00%的用户
**/
var m = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

//func romanToInt(s string) int {
//	sum := 0
//	//因为每次都要拿后一个元素做对比，所以后指针会越界，所以循环的长度要-1，然后在最后面再把最后一个元素加上
//	for i := 0; i < len(s)-1; i++ {
//		//双指针，如果后一个元素比前面的大，说明是减法
//		if m[s[i]] < m[s[i+1]] {
//			sum -= m[s[i]]
//		} else {
//			//如果后一个元素比前面的小，说明是加法
//			sum += m[s[i]]
//		}
//	}
//	sum += m[s[len(s)-1]]
//	return sum
//}

func main() {
	str := "MCMXCIV"
	fmt.Println(romanToInt(str))
}

/**
*执行结果：通过
*执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
*内存消耗：3.1 MB, 在所有 Go 提交中击败了100.00%的用户
 */
//哈希表的时间还能再省一些，使用switch
func getInt(r byte) int {
	switch r {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000
	default:
		panic(fmt.Sprintf("没有这个字母: %q", string(r)))
	}
}

func romanToInt(s string) int {
	sum := 0
	//因为每次都要拿后一个元素做对比，所以后指针会越界，所以循环的长度要-1，然后在最后面再把最后一个元素加上
	for i := 0; i < len(s)-1; i++ {
		//双指针，如果后一个元素比前面的大，说明是减法
		if getInt(s[i]) < getInt(s[i+1]) {
			sum -= getInt(s[i])
		} else {
			//如果后一个元素比前面的小，说明是加法
			sum += getInt(s[i])
		}
	}
	sum += getInt(s[len(s)-1])
	return sum
}
