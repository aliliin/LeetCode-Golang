package main

// 面试题 01.01. 判定字符是否唯一

func main() {
	isUnique("leetcode")
}

func isUnique(astr string) bool {
	strLen := len(astr)
	if strLen <= 1 {
		return true
	}
	for i := 0; i < strLen; i++ {
		for j := i + 1; j < strLen; j++ {
			if i != j && astr[i] == astr[j] {
				return false
			}
		}
	}
	return true
}
