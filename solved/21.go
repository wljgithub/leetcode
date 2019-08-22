// Merge two sorted linked lists and return it as a new list. The new list should be made by splicing together the nodes of the first two lists.
// 
// Example:
// 
// Input: 1->2->4, 1->3->4
// Output: 1->1->2->3->4->4

// 说实话，这道题，总感觉这个解答是不会AC的，因为会漏掉一个元素，我也有在本地跑过，但是却AC了
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := new(ListNode)
	for node := dummy; l1 != nil || l2 != nil; node = node.Next {
		if l1 == nil {
			node.Next = l2
			break
		} else if l2 == nil {
			node.Next = l1
			break
		} else if l1.Val < l2.Val {
			node.Next = l1
			l1 = l1.Next
		} else {
			node.Next = l2
			l2 = l2.Next
		}
	}
	return dummy.Next
}
