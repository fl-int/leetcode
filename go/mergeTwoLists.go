package main

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// 1 -> 1 -> 3 -> 5
	// 2 -> 3 -> 3 -> 7
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	if list1.Val > list2.Val {
		list1, list2 = list2, list1
	}
	l1, l2 := list1, list2
	for l1.Next != nil && l2 != nil {
		if l1.Next.Val < l2.Val {
			l1 = l1.Next
		} else {
			next := l2.Next
			l2.Next = l1.Next
			l1.Next = l2
			l2 = next

			l1 = l1.Next
		}
	}

	if l2 != nil {
		l1.Next = l2
	}

	return list1
}
