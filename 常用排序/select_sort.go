package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	list := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
	SelectSort(list)
	fmt.Println(list)

	list4 := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6}
	SelectSortGood(list4)
	fmt.Println(list4)

	var data []int
	//避免伪随机
	rand.Seed(time.Now().Unix())
	for i := 0; i < 100; i++ {
		data = append(data, rand.Intn(1000))
	}

	SelectSortGood(data)
	fmt.Println(data)
}

func SelectSort(list []int) {
	n := len(list)
	if list == nil || n <= 1 {
		return
	}
	for i := 0; i < n-1; i++ {
		// 从i的位置开始i查询最小的数据
		min := list[i] // 最小数
		minIndex := i  // 最小数据的下标
		for j := i + 1; j < n; j++ {
			if list[j] < min {
				min = list[j]
				minIndex = j
			}
		}
		// 这一轮找到的最小数的下标不等于最开始的下标，交换元素
		if i != minIndex {
			list[i], list[minIndex] = list[minIndex], list[i]
		}
	}
}

func SelectSortGood(list []int) {
	n := len(list)
	if list == nil || n <= 1 {
		return
	}

	for i := 0; i < n/2; i++ {
		minIndex := i
		maxIndex := i
		for j := i + 1; j < n-i; j++ {
			// 找到最大值下标
			if list[j] > list[maxIndex] {
				maxIndex = j // 这一轮这个是大的，直接 continue
				continue
			}
			// 找到最小值下标
			if list[j] < list[minIndex] {
				minIndex = j // 这一轮这个是大的，直接 continue
			}
		}

		if maxIndex == i && maxIndex != n-i-1 {
			// 如果最大值是开头的元素，而最小值不是最尾的元素
			// 先将最大值和最尾的元素交换
			list[n-i-1], list[maxIndex] = list[maxIndex], list[n-i-1]
			// 然后最小的元素放在最开头
			list[i], list[minIndex] = list[minIndex], list[i]
		} else if maxIndex == i && minIndex == n-i-1 {
			// 如果最大值在开头，最小值在结尾，直接交换
			list[minIndex], list[maxIndex] = list[maxIndex], list[minIndex]
		} else {
			// 否则先将最小值放在开头，再将最大值放在结尾
			list[i], list[minIndex] = list[minIndex], list[i]
			list[n-i-1], list[maxIndex] = list[maxIndex], list[n-i-1]
		}
	}
}
