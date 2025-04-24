package main

import (
	"container/heap"
	"testing"
)

// ListNode 定义链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}

// PriorityQueue 实现优先队列（最小堆）
type PriorityQueue []*ListNode

// 实现 heap.Interface 接口
func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].Val < pq[j].Val }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*ListNode))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

// mergeKLists 合并 K 个有序链表
// 关键点：
// 1. 使用优先队列（最小堆）来维护当前所有链表的最小值
// 2. 优化性能，减少不必要的操作
// 3. 时间复杂度：O(NlogK)，其中 N 是所有节点的总数，K 是链表数量
// 4. 空间复杂度：O(K)，堆的大小最多为 K
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	// 预处理：移除空链表，减少不必要的操作
	validLists := make([]*ListNode, 0, len(lists))
	for _, list := range lists {
		if list != nil {
			validLists = append(validLists, list)
		}
	}
	if len(validLists) == 0 {
		return nil
	}

	// 如果只有一个链表，直接返回
	if len(validLists) == 1 {
		return validLists[0]
	}

	// 初始化优先队列，预分配空间
	pq := make(PriorityQueue, 0, len(validLists))
	heap.Init(&pq)

	// 将所有链表的头节点加入优先队列
	for _, list := range validLists {
		heap.Push(&pq, list)
	}

	// 创建虚拟头节点
	dummy := &ListNode{}
	current := dummy

	// 不断从优先队列中取出最小值，直到队列为空
	for pq.Len() > 0 {
		// 取出当前最小值
		minNode := heap.Pop(&pq).(*ListNode)
		// 将最小值加入结果链表
		current.Next = minNode
		current = current.Next

		// 如果最小值节点还有下一个节点，将其加入优先队列
		if minNode.Next != nil {
			heap.Push(&pq, minNode.Next)
		}
	}

	return dummy.Next
}

// 辅助函数：创建链表
func createList(nums []int) *ListNode {
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

// 辅助函数：比较两个链表是否相等
func compareLists(l1, l2 *ListNode) bool {
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	return l1 == nil && l2 == nil
}

func TestMergeKLists(t *testing.T) {
	testCases := []struct {
		name     string
		lists    []*ListNode
		expected []int
	}{
		{
			name: "基本测试用例",
			lists: []*ListNode{
				createList([]int{1, 4, 5}),
				createList([]int{1, 3, 4}),
				createList([]int{2, 6}),
			},
			expected: []int{1, 1, 2, 3, 4, 4, 5, 6},
		},
		{
			name:     "空输入",
			lists:    []*ListNode{},
			expected: []int{},
		},
		{
			name: "包含空链表",
			lists: []*ListNode{
				createList([]int{1, 2}),
				nil,
				createList([]int{3, 4}),
			},
			expected: []int{1, 2, 3, 4},
		},
		{
			name: "单个链表",
			lists: []*ListNode{
				createList([]int{1, 2, 3}),
			},
			expected: []int{1, 2, 3},
		},
		{
			name: "包含重复元素",
			lists: []*ListNode{
				createList([]int{1, 1, 1}),
				createList([]int{1, 1, 1}),
				createList([]int{1, 1, 1}),
			},
			expected: []int{1, 1, 1, 1, 1, 1, 1, 1, 1},
		},
		{
			name: "边界值测试",
			lists: []*ListNode{
				createList([]int{-10000}),
				createList([]int{10000}),
			},
			expected: []int{-10000, 10000},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := mergeKLists(tc.lists)
			expected := createList(tc.expected)
			if !compareLists(result, expected) {
				t.Errorf("测试失败: %s", tc.name)
			}
		})
	}
}
