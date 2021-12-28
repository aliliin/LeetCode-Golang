package main

import (
	"fmt"
	"sync"
)

// ArrayQueue 数组队列，先进先出
type ArrayQueue struct {
	array []string   // 底层切片
	size  int        // 队列的元素数量
	lock  sync.Mutex // 为了并发安全使用的锁
}

func main() {
	q := new(ArrayQueue)
	q.Add("1")
	q.Add("2")
	q.Add("3")
	fmt.Println(q.Remove())
	fmt.Println(q.Remove())
	fmt.Println(q.Remove())
}

func (queue *ArrayQueue) Add(s string) {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	// 后进去的元素放到最后面
	queue.array = append(queue.array, s)
	// 队中元素数量+1
	queue.size = queue.size + 1
}

func (queue *ArrayQueue) Remove() string {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	// 队列是否为空了
	if queue.size == 0 {
		panic("empty ArrayQueue")
	}
	// 队列最前面元素
	v := queue.array[0]
	// 直接原位移动，原数组缩容,但缩容后继的空间不会被释放

	// 创建新的数组，移动次数过多
	mewQueue := make([]string, queue.size-1, queue.size-1)
	for i := 1; i < queue.size; i++ {
		mewQueue[i-1] = queue.array[i]
	}
	queue.array = mewQueue
	queue.size = queue.size - 1
	return v

}
