package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	//定义结果数组
	var res [][]int
	n := len(nums)
	if n < 3 {
		return res
	}

	sort.Ints(nums)
	for i := 0; i < n-2; i++ {
		if nums[i] > 0 {
			return res
		}
		//每次循环重新初始话前指针和后指针
		L, R := i+1, n-1
		if i > 0 && nums[i] == nums[n-1] {
			continue
		}
		//当前指针小于后指针的时候
		for L < R {
			sum := nums[i] + nums[L] + nums[R]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[L], nums[R]})
				//前指针往后扫描，如果有值相同的，直接跳过，去掉重复的结果
				for L < R && nums[L] == nums[L+1] {
					L++
				}

				//后指针往前扫描，如果有值相同的，直接跳过，去掉重复的结果
				for L < R && nums[R] == nums[R-1] {
					R--
				}
				L++
				R--
			} else if sum > 0 {
				R--
			} else {
				L++
			}
		}
		//外面的指针i也要去重，遇到重复的值直接跳过
		for i < len(nums)-3 && nums[i] == nums[i+1] {
			i++
		}
	}
	return res
}
func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	res := threeSum(nums)
	fmt.Print(res)
}
