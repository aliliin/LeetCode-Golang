package main

import "fmt"

func main() {
	obj := Constructor()
	obj.Push(1)
	obj.Push(2)
	param_2 := obj.Pop()
	fmt.Println(param_2)
	param_3 := obj.Peek()
	fmt.Println(param_3)
	param_4 := obj.Empty()
	fmt.Println(param_4)
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
// 双栈
// 将一个栈当作输入栈，用于压入 \texttt{push}push 传入的数据；另一个栈当作输出栈，用于 \texttt{pop}pop 和 \texttt{peek}peek 操作。
//每次 \texttt{pop}pop 或 \texttt{peek}peek 时，若输出栈为空则将输入栈的全部数据依次弹出并压入输出栈，这样输出栈从栈顶往栈底的顺序就是队列从队首往队尾的顺序。

type MyQueue struct {
	inStack, outStack []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{}
}

func (q *MyQueue) in2out() {
	for len(q.inStack) > 0 {
		q.outStack = append(q.outStack, q.inStack[len(q.inStack)-1])
		q.inStack = q.inStack[:len(q.inStack)-1]
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.inStack = append(this.inStack, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if len(this.outStack) == 0 {
		this.in2out()
	}
	x := this.outStack[len(this.outStack)-1]
	this.outStack = this.outStack[:len(this.outStack)-1]
	return x
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if len(this.outStack) == 0 {
		this.in2out()
	}
	return this.outStack[len(this.outStack)-1]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.inStack) == 0 && len(this.outStack) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
