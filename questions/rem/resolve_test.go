package main

import "testing"

func isMatch(s string, p string) bool {
	m, n := len(s), len(p)

	// 使用单个数组优化空间
	dp := make([]bool, n+1)
	dp[0] = true

	// 预处理模式串，处理以*开头的情况
	for j := 1; j <= n; j++ {
		if p[j-1] == '*' {
			dp[j] = dp[j-2]
		}
	}

	// 使用临时变量存储上一行的状态
	var prev bool
	for i := 1; i <= m; i++ {
		prev = dp[0]
		dp[0] = false

		for j := 1; j <= n; j++ {
			temp := dp[j]
			if p[j-1] == s[i-1] || p[j-1] == '.' {
				dp[j] = prev
			} else if p[j-1] == '*' {
				// 处理*的情况
				if p[j-2] == s[i-1] || p[j-2] == '.' {
					dp[j] = dp[j] || dp[j-2]
				} else {
					dp[j] = dp[j-2]
				}
			} else {
				dp[j] = false
			}
			prev = temp
		}

		// 快速检查是否有匹配可能
		hasMatch := false
		for j := 0; j <= n; j++ {
			if dp[j] {
				hasMatch = true
				break
			}
		}
		if !hasMatch {
			return false
		}
	}

	return dp[n]
}

func TestIsMatch(t *testing.T) {
	tests := []struct {
		s, p string
		want bool
	}{
		{"aa", "a", false},
		{"aa", "a*", true},
		{"ab", ".*", true},
		{"aab", "c*a*b", true},
		{"mississippi", "mis*is*p*.", false},
	}

	for _, test := range tests {
		if got := isMatch(test.s, test.p); got != test.want {
			t.Errorf("isMatch(%q, %q) = %v; want %v", test.s, test.p, got, test.want)
		}
	}
}
