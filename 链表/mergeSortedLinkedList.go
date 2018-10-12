package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// 取出最后一个元素
	var templ2 = l2
	var templ1 =l1
	// 遍历l2，取出元素再一一和l1的比较，如果小于，则插在前面
	for templ2.Next != nil {
		for templ1.Next!=nil{
			if templ2.Val < templ1.Val{
				// 我没办法判断链表的元素是不是头。。
			}

			templ1 = templ1.Next
		}
		templ2 = templ2.Next
	}
	return l1
}

func main() {
	//l1 := ListNode{Val: 1, Next: &ListNode{Val: 2, Next:nil }}
	//l2 := ListNode{Val: 3, Next: &ListNode{Val: 4, Next:nil }}
	//fmt.Println(mergeTwoLists(&l1,&l2).Next.Next.Next.Val)
	fmt.Println([]int{}==nil)
}
