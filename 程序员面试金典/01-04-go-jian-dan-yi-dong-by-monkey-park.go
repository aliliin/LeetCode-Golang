package main

import "fmt"

/**
面试题 01.04. 回文排列
	输入："tactcoa"
	输出：true
给定一个字符串，编写一个函数判定其是否为某个回文串的排列之一。

回文串是指正反两个方向都一样的单词或短语。排列是指字母的重新排列。

回文串不一定是字典当中的单词。

*/
func main() {
	fmt.Println(canPermutePalindrome("abctycba"))
	fmt.Println(canPermutePalindrome1("abctycba"))
}

func canPermutePalindrome1(s string) bool {
	m := make(map[rune]int)
	for _, v := range s {
		m[v]++
		if m[v] == 2 {
			delete(m, v)
		}
	}
	return len(m) <= 1
}

func canPermutePalindrome(s string) bool {
	m := make(map[rune]int)
	for _, v := range s {
		m[v]++
		if m[v]%2 == 0 {
			delete(m, v)
		}
	}
	return len(m) <= 1
}
