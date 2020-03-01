package main

func main() {

}

/*
20 ms   31.25%
5.8 MB   100.00%
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeDuplicateNodes(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	arr := make([]bool, 20001)
	begin := head
	end := head.Next
	arr[begin.Val] = true
	for end != nil {
		if begin.Val != end.Val && !arr[end.Val] {
			begin.Next = end
			arr[end.Val] = true
			begin = begin.Next
		}
		end = end.Next
	}
	begin.Next = nil
	return head
}
