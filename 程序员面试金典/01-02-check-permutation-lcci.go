package main

import (
	"reflect"
	"sort"
	"strings"
)

// 面试题 01.02. 判定是否互为字符重排
func main() {
	CheckPermutation("abc", "bca")
	CheckPermutation1("abc", "bca")
}

func CheckPermutation(s1 string, s2 string) bool {

	if len(s1) != len(s2) {
		return false
	}
	// 新建两个 slice
	strs1 := []int{}
	strs2 := []int{}
	for i := 0; i < len(s1); i++ {
		strs1 = append(strs1, int(s1[i]))
		strs2 = append(strs2, int(s2[i]))
	}
	// 排序 slice
	sort.Ints(strs1)
	sort.Ints(strs2)
	// 比较 slice 是否相等
	return reflect.DeepEqual(strs1, strs2)
}

func CheckPermutation1(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	s1Number := 0
	s2Number := 0

	for _, m := range []rune(s1) {
		s1Number += int(m)
	}
	for _, m := range []rune(s2) {
		s2Number += int(m)
	}
	return s1Number == s2Number
}

func CheckPermutation2(s1 string, s2 string) bool {
	a1 := strings.Split(s1, "")
	a2 := strings.Split(s2, "")
	sort.Strings(a1)
	sort.Strings(a2)
	return reflect.DeepEqual(a1, a2)
}
