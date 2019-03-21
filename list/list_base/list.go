package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

// 初始化头节点
func Init() *Node {
	return nil
}

var head Node

// 增加节点数据
func Insert(head *Node, num int) error {
	newNode := new(Node)
	newNode.data = num
	for {
		if head.next == nil {
			head.next = newNode
			return nil
		} else {
			head = head.next
		}
	}

}

// 删除节点

func Delete(head *Node, num int) error {
	pre := FindListPrevics(head, num)
	for {
		if head == nil {
			return nil
		} else {
			if num == head.data {
				pre.next = head.next
			}
			head = head.next
		}
	}
}

func FindListPrevics(head *Node, num int) *Node {

	for {
		if head.next != nil && head.next.data != num {
			head = head.next
		} else {
			fmt.Println("previces", head)
			return head
		}
	}
}

// 获取节点

func Get(head *Node) []*Node {
	nodes := []*Node{}
	for {
		if head == nil {
			break
		} else {
			tmp := head
			nodes = append(nodes, tmp)
			fmt.Println("data ", tmp.data)
			head = head.next
		}
	}
	return nodes
}

func main() {
	headPtr := &Node{
		next: nil,
	}
	// headPtr := node.Init()
	for i := 0; i < 10; i++ {
		Insert(headPtr, i)
	}
	FindListPrevics(headPtr, 8)
	Delete(headPtr, 8)
	nodes := Get(headPtr)
	fmt.Println("nodes===", nodes)
	/* 	for _, v := range nodes {
		tmp := v
		fmt.Println(v.data, &tmp)
	} */
}

/* package main

import (
	"fmt"
)

type Node struct {
	Value int
	Next  *Node
}

//测试链表是否为空
func IsEmpty(list *Node) bool {
	return list == nil
}

//测试是否为最后节点
func IsLast(position *Node, list *Node) bool {
	return position.Next == nil
}

func Find(value int, list *Node) *Node {
	p := list
	for p.Value != value {
		p = p.Next
	}
	return p
}

//插入节点
func Insert(value int, list *Node, position *Node) {
	tempCell := new(Node)
	if tempCell == nil {
		fmt.Println("out of space!")
	}
	tempCell.Value = value
	tempCell.Next = position.Next
	position.Next = tempCell
}

//删除节点
func Delete(value int, list *Node) {
	pPrev := FindPrevious(value, list)
	p := Find(value, list)
	pPrev.Next = p.Next
	p = nil
}

func FindPrevious(value int, list *Node) *Node {
	p := list
	for p.Next != nil && p.Next.Value != value {
		p = p.Next
	}
	return p
}

//删除链表 注意点：Golang中函数传递指针变量，其实是传递实参
func DeleteList(list **Node) {
	p := list
	for *p != nil {
		tmp := *p
		*p = nil
		*p = tmp.Next
	}
}

//打印列表
func PrintList(list *Node) {
	p := list
	for p != nil {
		fmt.Printf("%d-%p-%p ", p.Value, p, p.Next)
		p = p.Next
	}
	fmt.Println()
}

func main() {
	headNode := &Node{
		Value: 0,
		Next:  nil,
	}
	list := headNode
	Insert(1, list, headNode)
	PrintList(list)
	fmt.Println(IsEmpty(list))
	fmt.Println(IsLast(headNode, list))
	p := Find(0, list)
	Insert(2, list, p)
	PrintList(list)
	Delete(1, list)
	PrintList(list)
	DeleteList(&list)
	PrintList(list)
} */
