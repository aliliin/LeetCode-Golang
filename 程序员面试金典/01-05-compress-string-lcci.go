package main

import (
	"fmt"
	"strconv"
	"strings"
)

/**
面试题 01.06. 字符串压缩
输入："aabcccccaaa"
输出："a2b1c5a3"

用 ch 记录当前要压缩的字符，
cnt 记录 ch 出现的次数，如果当前枚举到的字符 s[i] 等于 ch ，
我们就更新 cnt 的计数，即 cnt = cnt + 1，否则我们按题目要求将 ch 以及
cnt 更新到答案字符串 ans 里，即 ans = ans + ch + cnt，完成对 ch 字符的压缩。
随后更新 ch 为 s[i]，cnt 为 11，表示将压缩的字符更改为 s[i]
*/
func main() {
	fmt.Println(compressString1("aabcccccaaa"))
	fmt.Println(compressString("aabcccccaaa"))
}

func compressString(S string) string {
	lenNum := len(S)
	// 空串处理
	if lenNum == 0 {
		return S
	}

	var sb strings.Builder
	curr := S[0]
	currLen := 1

	for i := 1; i < lenNum; i++ {
		if S[i] == curr {
			currLen++
		} else {
			sb.WriteByte(curr)
			sb.WriteString(strconv.Itoa(currLen))
			curr = S[i]
			currLen = 1
		}
	}
	
	sb.WriteByte(curr)
	sb.WriteString(strconv.Itoa(currLen))

	if sb.Len() >= len(S) {
		return S
	}

	return sb.String()

}

func compressString1(S string) string {
	lenNum := len(S)
	// 空串处理
	if lenNum == 0 {
		return S
	}
	var ans string
	cnt := 1
	ch := S[0]
	for i := 1; i < lenNum; i++ {
		if ch == S[i] {
			cnt++
		} else {
			ans = strings.Join([]string{ans, string(ch), strconv.Itoa(cnt)}, "")
			ch = S[i]
			cnt = 1
		}
	}
	ans = strings.Join([]string{ans, string(ch), strconv.Itoa(cnt)}, "")
	if len(ans) >= lenNum {
		return S
	} else {
		return ans
	}
}
