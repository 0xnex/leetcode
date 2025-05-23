# K 个一组翻转链表

## 问题描述

给定一个链表的头节点 `head`，每 k 个节点一组进行翻转，并返回修改后的链表。

k 是一个正整数，且小于或等于链表的长度。如果节点总数不是 k 的整数倍，那么最后剩余的节点保持原有顺序。

## 解题思路

### 方法一：递归法

1. 首先检查是否还有至少 k 个节点
2. 如果有 k 个节点：
   - 反转这 k 个节点
   - 递归处理剩余的链表
3. 如果不足 k 个节点，保持原样

### 方法二：迭代法（优化空间）

1. 使用一个虚拟头节点（dummy）来统一处理边界情况
2. 维护一个 `prevGroup` 指针，指向当前要处理组的前一个节点
3. 对于每一组 k 个节点：
   - 检查是否还有足够的节点
   - 反转当前组内的节点
   - 将反转后的组连接到前一组和后一组
   - 移动到下一组

## 关键实现细节

1. **虚拟头节点**：

   - 使用 dummy 节点可以统一处理头节点的特殊情况
   - 避免了对头节点的特殊判断

2. **组内反转**：

   - 使用三个指针（groupPrev, current, next）进行组内反转
   - 反转过程中保持链表的连续性

3. **组间连接**：
   - 反转后需要正确连接前一组和后一组
   - 通过 prevGroup 和 nextGroup 维护组间的连接

## 复杂度分析

### 时间复杂度

- O(n)，其中 n 是链表的长度
- 每个节点只被访问一次

### 空间复杂度

- O(1)，只使用了常数个额外空间
- 相比递归方法的 O(n/k) 空间复杂度有了显著优化

## 示例分析

以输入 `[1,2,3,4,5]`，k = 2 为例：

1. 初始状态：

   ```
   dummy -> 1 -> 2 -> 3 -> 4 -> 5
   ```

2. 第一组反转后：

   ```
   dummy -> 2 -> 1 -> 3 -> 4 -> 5
   ```

3. 第二组反转后：

   ```
   dummy -> 2 -> 1 -> 4 -> 3 -> 5
   ```

4. 最终结果：
   ```
   [2,1,4,3,5]
   ```
