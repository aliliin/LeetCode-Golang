package main

import "fmt"

//将三合一的栈Stacks[]平均分成三份
//使用 Len[] 保存三个栈的长度，入栈就将对应的 Len[Num]++,出栈就 Len[Num]--
//Size 保存初始化的每个栈大小，每次 push() or peek()判断当前栈是否占满，占满则啥也不做
//如何索引具体位置？使用Num * Size来索引该栈的开始位置，后面加上 Len[Num] 即当前元素应该占的位置，出栈则-1。

type TripleInOne struct {
	Stacks []int
	Len    []int
	Size   int
}

func Constructor(stackSize int) TripleInOne {
	return TripleInOne{
		Stacks: make([]int, stackSize*3),
		Len:    make([]int, 3),
		Size:   stackSize,
	}
}

func (this *TripleInOne) Push(stackNum int, value int) {
	if this.Len[stackNum] == this.Size {
		return
	}
	this.Stacks[stackNum*this.Size+this.Len[stackNum]] = value
	this.Len[stackNum]++
}

func (this *TripleInOne) Pop(stackNum int) int {
	if this.Len[stackNum] == 0 {
		return -1
	}
	x := this.Stacks[stackNum*this.Size+this.Len[stackNum]-1]
	this.Len[stackNum]--
	return x
}

func (this *TripleInOne) Peek(stackNum int) int {
	if this.Len[stackNum] == 0 {
		return -1
	}
	return this.Stacks[stackNum*this.Size+this.Len[stackNum]-1]
}

func (this *TripleInOne) IsEmpty(stackNum int) bool {
	return this.Len[stackNum] == 0
}

func main() {

	// 输入：["TripleInOne", "push", "push", "pop", "pop", "pop", "isEmpty"]
	// [[1], [0, 1], [0, 2], [0], [0], [0], [0]]
	// 输出：[null, null, null, 1, -1, -1, true]
	tripleInOne := Constructor(1)
	tripleInOne.Push(0, 1)
	tripleInOne.Push(0, 2)
	fmt.Println(tripleInOne.Pop(0))
	fmt.Println(tripleInOne.Pop(0))
	fmt.Println(tripleInOne.Pop(0))
	fmt.Println(tripleInOne.IsEmpty(0))
}
