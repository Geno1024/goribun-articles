package main

import (
	"fmt"
)

func main() {
	Joseph(3, 1500000)
}

//定义节点
type Node struct {
	num  int
	next *Node
}

var p, r, list *Node

func Joseph(m, n int) {
	//创建链表
	for i := 1; i <= n; i++ {
		p = &Node{num: 0, next: nil}
		p.num = i
		if list == nil {
			list = p
		} else {
			r.next = p
		}
		r = p
	}

	p.next = list //使链表循环
	p = list      //p指向头结点
	r = p

	//循环删除队列中的结点，即出列
	fmt.Print("出列者序列：")
	for p.next != p {
		for i := 1; i < m; i++ {
			r = p
			p = p.next
		}
		r.next = p.next
		fmt.Printf("%d ", p.num)
		p = r.next

	}
	fmt.Printf("\n最后留下的人是：%d\n", p.num)
}

