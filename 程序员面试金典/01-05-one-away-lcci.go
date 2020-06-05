package main

import (
	"fmt"
)

/**
	面试题 01.05. 一次编辑
输入:
first = "pale"
second = "ple"
输出: True

输入:
first = "pales"
second = "pal"
输出: False

*/

func main() {
	fmt.Println(oneEditAway("intention", "execution"))
	fmt.Println(oneEditAway2("intention", "execution"))
	fmt.Println(oneEditAway4("intention", "execution"))
}

// dp 的解法
func oneEditAway2(first string, second string) bool {
	len1, len2 := len(first), len(second)
	if first == second {
		return true
	}
	if absInt(int64(len1-len2)) > 1 {
		return false
	}

	dp := make([][]int, len1+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len2+1)
	}
	for i := 0; i < len(dp); i++ {
		dp[i][0] = i
	}
	for j := 0; j < len(dp[0]); j++ {
		dp[0][j] = j
	}
	for i := 1; i <= len1; i++ {
		for j := 1; j <= len2; j++ {
			dp[i][j] = minDis(dp[i-1][j], minDis(dp[i][j-1], dp[i-1][j-1]))
			if first[i-1] != second[j-1] {
				dp[i][j]++
			}
		}
	}
	return dp[len1][len2] < 2

}

func oneEditAway4(first string, second string) bool {
	len1, len2 := len(first), len(second)
	if first == second {
		return true
	}
	if absInt(int64(len1-len2)) > 1 {
		return false
	}

	dp := make([][]int, len1+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len2+1)
	}
	for i := 0; i < len(dp); i++ {
		dp[i][0] = i
	}
	for j := 0; j < len(dp[0]); j++ {
		dp[0][j] = j
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[0]); j++ {
			diagonal := 1
			if first[i-1] == second[j-1] {
				diagonal = 0
			}
			dp[i][j] = minDis(dp[i-1][j]+1, minDis(dp[i][j-1]+1, dp[i-1][j-1]+diagonal))
		}
	}
	return dp[len1][len2] < 2
}
func minDis(a, b int) int {
	min := a
	if a > b {
		min = b
	}
	return min
}

// 双循环解法
func oneEditAway(first string, second string) bool {
	if first == second {
		return true
	}

	len1 := len(first)
	len2 := len(second)
	num := len1 - len2
	if absInt(int64(num)) > 1 {
		return false
	}
	i := 0
	j := len1 - 1
	k := len2 - 1
	for i < len1 && i < len2 && first[i] == second[i] {
		i++
	}
	for j >= 0 && k >= 0 && first[j] == second[k] {
		j--
		k--
	}
	return j-i < 1 && k-i < 1
}

func absInt(n int64) int64 {

	if n < 0 {
		n = -n
	}
	return n
}
