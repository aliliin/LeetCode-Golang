package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 冒泡排序
func main() {
	var data []int
	//避免伪随机
	rand.Seed(time.Now().Unix())

	for i := 0; i < 100; i++ {
		data = append(data, rand.Intn(1000))
	}

	BubbleSort3(data)

	fmt.Println(data)

	list := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
	BubbleSort1(list)
	fmt.Println(list)

	BubbleSort2(list)
	fmt.Println(list)
}

func BubbleSort3(data []int) {
	if data == nil || len(data) <= 1 {
		return
	}
	l := len(data)
	swap := false
	for i := 0; i < l; i++ {
		for j := 0; j < l-i-1; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
				swap = true
			}
		}
	}
	if !swap {
		return
	}
}

func BubbleSort2(array []int) {
	n := len(array)
	swap := false

	for i := n - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
				swap = false
			}
		}
	}

	if !swap {
		return
	}

}

func BubbleSort1(list []int) {
	n := len(list)
	// 在一轮中有没有交换过
	didSwap := false
	// 进行 N-1 轮迭代
	for i := n - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			// 如果前面的数比后面的大，那么交换
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
				didSwap = true
			}
		}
	}
	// 如果在一轮中没有交换过，那么已经排好序了，直接返回
	if !didSwap {
		return
	}
}

func BubbleSort(data []int) {
	if data == nil || len(data) <= 1 {
		return
	}

	swap := true
	for i := 0; i < len(data) && swap; i++ {
		swap = false
		for j := 0; j < len(data)-i-1; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
				swap = true
			}
		}
	}
}
