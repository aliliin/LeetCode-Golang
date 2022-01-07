package main

import (
	"fmt"
)

type StackOfPlates struct {
	Cap   int
	Stack [][]int
}

func Constructor(cap int) StackOfPlates {
	return StackOfPlates{
		Cap:   cap,
		Stack: make([][]int, 0),
	}
}

func (this *StackOfPlates) Push(val int) {
	if this.Cap == 0 {
		return
	}
	n := len(this.Stack)
	if n == 0 {
		newStack := append(make([]int, 0), val)
		this.Stack = append(this.Stack, newStack)
		return
	}

	top := this.Stack[len(this.Stack)-1]
	if len(top) == this.Cap {
		newStack := append(make([]int, 0), val)
		this.Stack = append(this.Stack, newStack)
		return
	}
	top = append(top, val)
	this.Stack[len(this.Stack)-1] = top
}

func (this *StackOfPlates) Pop() int {
	if len(this.Stack) == 0 {
		return -1
	}
	plate := this.Stack[len(this.Stack)-1]
	v := plate[len(plate)-1]
	plate = plate[:len(plate)-1]
	this.Stack[len(this.Stack)-1] = plate
	if len(plate) == 0 {
		this.Stack = this.Stack[:len(this.Stack)-1]
	}
	return v
}

func (this *StackOfPlates) PopAt(index int) int {
	n := len(this.Stack)
	if index >= n {
		return -1
	}

	plate := this.Stack[index]
	v := plate[len(plate)-1]
	plate = plate[0 : len(plate)-1]
	this.Stack[index] = plate

	if len(plate) == 0 {
		tmp := this.Stack[index+1:]
		this.Stack = this.Stack[:index]
		this.Stack = append(this.Stack, tmp...)
	}
	return v
}

func main() {
	// 堆盘子。设想有一堆盘子，堆太高可能会倒下来。因此，在现实生活中，盘子堆到一定高度时，我们就会另外堆一堆盘子。
	//请实现数据结构SetOfStacks，模拟这种行为。SetOfStacks应该由多个栈组成，并且在前一个栈填满时新建一个栈。
	//此外，SetOfStacks.push()和SetOfStacks.pop()应该与普通栈的操作方法相同（也就是说，pop()返回的值，应该跟只有一个栈时的情况一样）。
	//进阶：实现一个popAt(int index)方法，根据指定的子栈，执行pop操作。
	//
	//当某个栈为空时，应当删除该栈。当栈中没有元素或不存在该栈时，pop，popAt 应返回 -1.

	// 该题的重点是需要事先准确定义各个变量，然后在编写代码的过程中，要经常思考：修改一个变量后，该变量以及其它相关变量是否仍然满足其定义？
	//0、栈的长度为0时push方法直接return
	//1、一个栈用一个一维数组表示，有多个栈，即二维数组
	//2、每个栈的长度是一样的
	//3、栈满的时候，将元素放到新的栈里
	//4、栈空的时候要删除

	Stack := Constructor(1)
	fmt.Println(Stack)
	// ["StackOfPlates", "push", "push", "popAt", "pop", "pop"]
	// [[1], [1], [2], [1], [], []]
	// 输出： [null, null, null, 2, 1, -1]

	Stack.Push(1)
	Stack.Push(2)
	fmt.Println(Stack.PopAt(1)) // 2
	fmt.Println(Stack.Pop())    // 1
	fmt.Println(Stack.Pop())    // -1

	//输入： ["StackOfPlates", "push", "push", "push", "popAt", "popAt", "popAt"]
	// [[2], [1], [2], [3], [0], [0], [0]]
	//输出： [null, null, null, null, 2, 1, 3]
	Stacks := Constructor(2)
	fmt.Println(Stacks)
	Stacks.Push(1)
	Stacks.Push(2)
	Stacks.Push(3)
	fmt.Println(Stacks.PopAt(0)) // 2
	fmt.Println(Stacks.PopAt(0)) // 1
	fmt.Println(Stacks.PopAt(0)) // 3

}
