package main

import (
	"fmt"
	"strings"
)

// 面试题 01.09. 字符串轮转
func main() {
	s1 := "waterbottle"
	s2 := "erbottlewat"
	fmt.Println("res", isFlipedString(s1, s2))
}

func isFlipedString(s1 string, s2 string) bool {
	return len(s1) == len(s2) && strings.Contains(s2+s2, s1)
}
