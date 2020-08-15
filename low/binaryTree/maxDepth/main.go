package main

import (
	"fmt"
	"github.com/LIWENHUI121017/tools/tree"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Left  *Tree
 *     Value int
 *     Right *Tree
 * }
 */

//深度搜索遍历
//func maxDepth(root *tree.Tree) int {
//	if root == nil {
//		return 0
//	}
//	return tool.MaxIntValue(maxDepth(root.Left), maxDepth(root.Right)) + 1
//}

//层序遍历
func maxDepth(root *tree.Tree) int {
	if root == nil {
		return 0
	}
	queue := make([]*tree.Tree, 0)
	queue = append(queue, root)

	depth := 0
	for len(queue) > 0 {
		n := len(queue)
		for n > 0 {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			n--
		}
		depth++
	}
	return depth
}

func main() {
	root := &tree.Tree{nil, 5, nil}
	root = tree.Insert(root, 2)
	root = tree.Insert(root, 7)
	root = tree.Insert(root, 3)
	root = tree.Insert(root, 9)
	root = tree.Insert(root, 1)
	root = tree.Insert(root, 6)
	fmt.Println(root)
	fmt.Println(maxDepth(root))

}
