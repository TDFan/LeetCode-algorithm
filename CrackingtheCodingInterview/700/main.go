package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val > val {
		return searchBST(root.Left, val)
	}
	if root.Val < val {
		return searchBST(root.Right, val)
	}
	return root
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
