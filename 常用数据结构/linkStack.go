package main

import (
	"fmt"
	"sync"
)

// LinkStack  链表形式的下压栈，后进先出
type LinkStack struct {
	root *Link      // 链表起点
	size int        // 栈的元素数量
	lock sync.Mutex // 为了并发安全使用的锁
}

// Link 链表节点
type Link struct {
	Next  *Link
	Value string
}

func main() {
	linkStack := new(LinkStack)
	linkStack.Push("g")
	linkStack.Push("y")
	linkStack.Push("l")
	fmt.Println("size:", linkStack.Size())
	fmt.Println("pop:", linkStack.Pop())
	fmt.Println("pop:", linkStack.Pop())
	fmt.Println("size:", linkStack.Size())
	linkStack.Push("gyl")
	fmt.Println("pop:", linkStack.Pop())
	fmt.Println("pop:", linkStack.Pop())
	fmt.Println("size:", linkStack.Size())
	fmt.Println("IsEmpty:", linkStack.IsEmpty())
}

// Push 入栈
func (stack *LinkStack) Push(v string) {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	if stack.root == nil {
		stack.root = new(Link)
		stack.root.Value = v
	} else {
		// 否则新元素插入 原来的链表 的头部
		prevNode := stack.root

		// 新节点
		newNode := new(Link)
		newNode.Value = v

		// 原来的链表链接到新元素后面
		newNode.Next = prevNode

		// 新节点放在头部
		stack.root = newNode
	}
	stack.size = stack.size + 1
}

// Pop 出栈
func (stack *LinkStack) Pop() string {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	if stack.size == 0 {
		panic("Empty LinkStack")
	}

	// 获取顶部元素
	top := stack.root
	v := top.Value

	// 将顶部元素的后继链接链接上
	stack.root = top.Next

	stack.size = stack.size - 1
	return v
}

// Peek 获取栈顶元素
func (stack *LinkStack) Peek() string {
	// 栈中元素已空
	if stack.size == 0 {
		panic("Empty LinkStack")
	}

	// 顶部元素值
	v := stack.root.Value
	return v
}

// Size 栈大小
func (stack *LinkStack) Size() int {
	return stack.size
}

// IsEmpty 栈是否为空
func (stack *LinkStack) IsEmpty() bool {
	return stack.size == 0
}
