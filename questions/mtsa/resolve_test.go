package main

import (
	"fmt"
	"math"
	"testing"
)

// findMedianSortedArrays 使用二分查找实现
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// 处理两个数组都为空的情况
	if len(nums1) == 0 && len(nums2) == 0 {
		return 0.0
	}

	// 确保 nums1 是较短的数组
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}

	m, n := len(nums1), len(nums2)
	total := m + n
	half := (total + 1) / 2

	// 在较短的数组上进行二分查找
	left, right := 0, m
	for left <= right {
		// 在 nums1 中找到一个分割点
		i := (left + right) / 2
		// 在 nums2 中对应的分割点
		j := half - i

		// 处理边界情况
		var nums1Left, nums1Right, nums2Left, nums2Right int

		if i == 0 {
			nums1Left = math.MinInt
		} else {
			nums1Left = nums1[i-1]
		}

		if i == m {
			nums1Right = math.MaxInt
		} else {
			nums1Right = nums1[i]
		}

		if j == 0 {
			nums2Left = math.MinInt
		} else {
			nums2Left = nums2[j-1]
		}

		if j == n {
			nums2Right = math.MaxInt
		} else {
			nums2Right = nums2[j]
		}

		// 检查分割是否满足条件
		if nums1Left <= nums2Right && nums2Left <= nums1Right {
			if total%2 == 0 {
				return float64(max(nums1Left, nums2Left)+min(nums1Right, nums2Right)) / 2
			} else {
				return float64(max(nums1Left, nums2Left))
			}
		} else if nums1Left > nums2Right {
			right = i - 1
		} else {
			left = i + 1
		}
	}

	return 0.0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func TestResovle(t *testing.T) {
	testCases := []struct {
		nums1   []int
		nums2   []int
		expects float64
	}{
		// 基本测试用例
		{[]int{1, 3}, []int{2}, 2.0},                   // 奇数长度
		{[]int{1, 2}, []int{3, 4}, 2.5},                // 偶数长度
		{[]int{0, 0}, []int{0, 0}, 0.0},                // 全零
		{[]int{}, []int{1}, 1.0},                       // 一个数组为空
		{[]int{2}, []int{}, 2.0},                       // 另一个数组为空
		{[]int{}, []int{}, 0.0},                        // 两个数组都为空
		{[]int{1, 2, 3}, []int{4, 5, 6}, 3.5},          // 两个数组长度相等
		{[]int{1, 3, 5}, []int{2, 4, 6}, 3.5},          // 交替分布的数组
		{[]int{1, 2, 3, 4}, []int{5, 6, 7, 8, 9}, 5.0}, // 一个数组比另一个长
		{[]int{1, 2, 3, 4, 5}, []int{6, 7, 8, 9}, 5.0}, // 另一个数组比一个长

		// 边界测试用例
		{[]int{-5, -3, -1}, []int{-4, -2, 0}, -2.5},  // 包含负数
		{[]int{1}, []int{2, 3, 4, 5, 6}, 3.5},        // 一个数组只有一个元素
		{[]int{1, 2, 3, 4, 5}, []int{6}, 3.5},        // 另一个数组只有一个元素
		{[]int{1}, []int{2}, 1.5},                    // 两个数组各一个元素
		{[]int{1, 2, 3}, []int{100, 200, 300}, 51.5}, // 两个数组差异较大
	}

	// 执行测试
	for i, tc := range testCases {
		result := findMedianSortedArrays(tc.nums1, tc.nums2)
		if result != tc.expects {
			t.Errorf("Test case %d failed: expected %v, got %v", i+1, tc.expects, result)
		} else {
			fmt.Printf("Test case %d passed: expected %v, got %v\n", i+1, tc.expects, result)
		}
	}
}
