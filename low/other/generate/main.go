package main

import (
	"fmt"
)

func generate(numRows int) [][]int {
	//声明结果数组
	var res = [][]int{}
	if numRows == 0 {
		return res
	}
	res = append(res, []int{1})
	for i := 1; i < numRows; i++ {
		//因为每加一层，其实就是比上一层多了一个元素
		//所以我们可以添加一个0在开头的位置，后续就可以利用当前位置的值等于当前元素加后一个元素的值来计算了
		row := []int{0}
		//直接将上一层数组加入到当前数组里面
		row = append(row, res[i-1]...)
		//这样循环的时候就不需要获取上一个
		for j := 0; j < i; j++ {
			row[j] = row[j] + row[j+1]
		}
		res = append(res, row)
	}
	return res

	//var res = [][]int{}
	//for i:=0;i<numRows;i++ {
	//	row := []int{};
	//	for j:=0;j<=i;j++{
	//		//第一个位置和最后一个位置的元素为1
	//		if (j==0 || j==i){
	//			row = append(row,1);
	//		}else {
	//			//上一行的元素进行相加
	//			row = append(row,res[i-1][j-1]+res[i-1][j]);
	//		}
	//	}
	//	res = append(res,row)
	//}
	//return res
}

func main() {
	//15的二进制为1111
	res := generate(6)
	fmt.Print(res)
}
