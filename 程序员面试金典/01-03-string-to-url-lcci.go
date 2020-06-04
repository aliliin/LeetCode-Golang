package main

import (
	"bytes"
	"fmt"
	"strings"
)

/**
	面试题 01.03. URL化
	输入："Mr John Smith    ", 13
 	输出："Mr%20John%20Smith"
*/

func main() {
	S := "Mr John Smith    "
	length := 13
	fmt.Println(replaceSpaces(S, length))
	fmt.Println(replaceSpaces1(S, length))
	fmt.Println(replaceSpaces2(S, length))
	fmt.Println(replaceSpaces3(S, length))

}

func replaceSpaces(S string, length int) string {
	return strings.Replace(S[:length], " ", "%20", -1)
}
func replaceSpaces1(S string, length int) string {
	var buffer bytes.Buffer
	for i := 0; i < length; i++ {
		if S[i] == ' ' {
			buffer.WriteString("%20")
		} else {
			buffer.WriteString(string(S[i]))
		}
	}
	return buffer.String()
}

func replaceSpaces3(S string, length int) string {
	var builder strings.Builder
	for i := 0; i < length; i++ {
		if S[i] == ' ' {
			builder.WriteString("%20")
		} else {
			builder.WriteString(string(S[i]))
		}
	}
	return builder.String()
}

// 超出时间限制
func replaceSpaces2(S string, length int) string {
	newStr := ""
	for i := 0; i < length; i++ {
		if S[i] == ' ' {
			newStr += "%20"
		} else {
			newStr += string(S[i])
		}
	}
	return newStr
}
