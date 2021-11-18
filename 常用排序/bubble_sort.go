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

	BubbleSort(data)

	fmt.Println(data)
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
