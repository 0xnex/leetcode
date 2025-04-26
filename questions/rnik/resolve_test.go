package main

import (
	"testing"
)

// ListNode represents a node in a singly-linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

// reverseKGroup reverses the nodes of a linked list k at a time
func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Next: head}
	prevGroup := dummy

	for {
		// Check if we have at least k nodes
		count := 0
		current := prevGroup.Next
		for current != nil && count < k {
			current = current.Next
			count++
		}

		if count < k {
			break
		}

		// Reverse the group of k nodes
		groupPrev := prevGroup.Next
		current = groupPrev.Next
		for i := 1; i < k; i++ {
			next := current.Next
			current.Next = groupPrev
			groupPrev = current
			current = next
		}

		// Connect the reversed group to the rest of the list
		nextGroup := prevGroup.Next
		prevGroup.Next.Next = current
		prevGroup.Next = groupPrev
		prevGroup = nextGroup
	}

	return dummy.Next
}

// Helper function to create a linked list from a slice
func createLinkedList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	current := head
	for i := 1; i < len(nums); i++ {
		current.Next = &ListNode{Val: nums[i]}
		current = current.Next
	}
	return head
}

// Helper function to convert linked list to slice
func linkedListToSlice(head *ListNode) []int {
	var result []int
	current := head
	for current != nil {
		result = append(result, current.Val)
		current = current.Next
	}
	return result
}

func TestReverseKGroup(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		k        int
		expected []int
	}{
		{
			name:     "Example 1",
			input:    []int{1, 2, 3, 4, 5},
			k:        2,
			expected: []int{2, 1, 4, 3, 5},
		},
		{
			name:     "Example 2",
			input:    []int{1, 2, 3, 4, 5},
			k:        3,
			expected: []int{3, 2, 1, 4, 5},
		},
		{
			name:     "Single group",
			input:    []int{1, 2, 3, 4, 5},
			k:        5,
			expected: []int{5, 4, 3, 2, 1},
		},
		{
			name:     "No reversal needed",
			input:    []int{1, 2, 3, 4, 5},
			k:        1,
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Empty list",
			input:    []int{},
			k:        2,
			expected: []int{},
		},
		{
			name:     "Single element",
			input:    []int{1},
			k:        2,
			expected: []int{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := createLinkedList(tt.input)
			result := reverseKGroup(head, tt.k)
			got := linkedListToSlice(result)

			if len(got) != len(tt.expected) {
				t.Errorf("length mismatch: got %v, want %v", got, tt.expected)
				return
			}

			for i := range got {
				if got[i] != tt.expected[i] {
					t.Errorf("at index %d: got %v, want %v", i, got, tt.expected)
					return
				}
			}
		})
	}
}
