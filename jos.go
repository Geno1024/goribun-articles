package main

import (
	"fmt"
)

func main() {
	Joseph(3, 1500000)
}

//����ڵ�
type Node struct {
	num  int
	next *Node
}

var p, r, list *Node

func Joseph(m, n int) {
	//��������
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

	p.next = list //ʹ����ѭ��
	p = list      //pָ��ͷ���
	r = p

	//ѭ��ɾ�������еĽ�㣬������
	fmt.Print("���������У�")
	for p.next != p {
		for i := 1; i < m; i++ {
			r = p
			p = p.next
		}
		r.next = p.next
		fmt.Printf("%d ", p.num)
		p = r.next

	}
	fmt.Printf("\n������µ����ǣ�%d\n", p.num)
}

