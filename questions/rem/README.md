[Regular Expression Matching](https://leetcode.com/problems/regular-expression-matching/)

Given an input string s and a pattern p, implement regular expression matching with support for '.' and '\*' where:

```sh
'.' Matches any single character.​​​​
'*' Matches zero or more of the preceding element.
```

The matching should cover the entire input string (not partial).

## Example 1:

```
Input: s = "aa", p = "a"
Output: false
Explanation: "a" does not match the entire string "aa".
```

## Example 2:

```
Input: s = "aa", p = "a*"
Output: true
Explanation: '*' means zero or more of the preceding element, 'a'. Therefore, by repeating 'a' once, it becomes "aa".
```

## Example 3:

```
Input: s = "ab", p = ".*"
Output: true
Explanation: ".*" means "zero or more (*) of any character (.)".
```

## Constraints:

```
1 <= s.length <= 20
1 <= p.length <= 20
s contains only lowercase English letters.
p contains only lowercase English letters, '.', and '*'.
It is guaranteed for each appearance of the character '*', there will be a previous valid character to match.
```

## 思路

### 动态规划解法

1. **状态定义**：
   - 使用二维数组 `dp[i][j]` 表示字符串 s 的前 i 个字符和模式 p 的前 j 个字符是否匹配
   - 最终答案存储在 `dp[m][n]` 中，其中 m 和 n 分别是字符串和模式的长度

2. **初始化**：
   - `dp[0][0] = true`：空字符串匹配空模式
   - 处理模式以 `*` 开头的情况（如 `a*`, `a*b*` 等）：
     ```go
     for j := 1; j <= n; j++ {
         if p[j-1] == '*' {
             dp[0][j] = dp[0][j-2]
         }
     }
     ```

3. **状态转移**：
   - 如果当前字符匹配或模式是 '.'：
     ```go
     if p[j-1] == s[i-1] || p[j-1] == '.' {
         dp[i][j] = dp[i-1][j-1]
     }
     ```
   - 如果当前模式是 '*'：
     - 情况1：'*' 匹配零个前面的元素：`dp[i][j] = dp[i][j-2]`
     - 情况2：'*' 匹配一个或多个前面的元素：`dp[i][j] = dp[i][j] || dp[i-1][j]`

4. **优化**：
   - 空间优化：使用两个一维数组替代二维数组，空间复杂度从 O(m*n) 降低到 O(n)
   - 提前终止：在每一行处理完后检查是否有匹配的可能，如果没有则提前返回 false

### 时间复杂度分析
- 时间复杂度：O(m*n)，其中 m 和 n 分别是字符串和模式的长度
- 空间复杂度：O(n)，经过优化后只需要 O(n) 的额外空间

### 关键点
1. 正确处理 '*' 的两种情况：匹配零个或多个前面的元素
2. 使用动态规划避免重复计算
3. 通过空间优化减少内存使用
4. 利用提前终止优化性能
