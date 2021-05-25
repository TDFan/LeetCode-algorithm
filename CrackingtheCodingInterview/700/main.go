package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	var left *TreeNode = nil
	var right *TreeNode = nil
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	if root.Val > val {
		left = searchBST(root.Left, val)
	}
	if root.Val < val {
		right = searchBST(root.Right, val)
	}
	if left != nil {
		return left
	}
	return right
}
func main() {
	root := &TreeNode{
		4,
		&TreeNode{
			2,
			&TreeNode{
				1,
				nil,
				nil,
			},
			&TreeNode{
				3,
				nil,
				nil,
			},
		},
		&TreeNode{
			7,
			nil,
			nil,
		},
	}
	bst := searchBST(root, 5)
	fmt.Println(bst)
}

/**
[4,2,7,1,3]
2
*/
