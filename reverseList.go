package main

import (
	"strconv"
	"strings"
)

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode

	for head != nil {
		tmp := head.Next
		head.Next = prev
		prev = head
		head = tmp
	}

	return prev
}

func printList(head *ListNode) string {
	sb := strings.Builder{}
	for head != nil {
		sb.WriteString(strconv.Itoa(head.Val))
		sb.WriteRune('>')
		head = head.Next
	}
	return sb.String()
}
