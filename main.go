package main

func main() {
	head1 := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}

	/*head2 := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
		},
	}*/

	reverseList(&head1)
}
