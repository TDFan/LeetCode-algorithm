package main

func main() {

}

/*
0 ms  100.00%
2.2 MB  100.00%
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

var temp bool

func kthToLast(head *ListNode, k int) int {
	temp = false
	return sovle(head, k)
}
func sovle(head *ListNode, k int) int {
	if head == nil {
		return k
	}
	k = sovle(head.Next, k)
	if temp {
		return k
	} else {
		if k <= 1 {
			temp = true
			return head.Val
		} else {
			return k - 1
		}
	}
	return k
}
