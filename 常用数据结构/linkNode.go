package main

import "fmt"

// LinkNode 简单的链表如下
type LinkNode struct {
	Data     int64
	NextNode *LinkNode
}

// Ring 循环链表
type Ring struct {
	next, prev *Ring       // 前驱和后驱节点
	Value      interface{} // 数据
}

// 初始化空的循环链表，前驱和后驱都指向自己，因为是循环的
func (r *Ring) init() *Ring {
	r.next = r
	r.prev = r
	return r
}

func main() {
	r := new(Ring)
	r.init()
	AddRingLink(5)

	node := new(LinkNode)
	node.Data = 2

	node1 := new(LinkNode)
	node1.Data = 3

	node.NextNode = node1

	node2 := new(LinkNode)
	node2.Data = 4

	node1.NextNode = node2

	// 按顺序打印数据

	newNode := node

	for {
		if newNode != nil {
			// 打印节点值
			fmt.Println(newNode.Data)
			// 下一个节点值
			newNode = newNode.NextNode
			continue
		}
		// 如果下一个节点为空，表示链表结束了
		break
	}

}

// AddRingLink 指定大小 N 的循环链表
func AddRingLink(n int) *Ring {
	if n <= 0 {
		return nil
	}
	r := new(Ring)
	p := r
	for i := 1; i < n; i++ {
		p.next = &Ring{prev: p, Value: i}
		p = p.next
	}
	p.next = r
	r.prev = p
	return r
}
