package main

import "fmt"

type Node struct {
	val  int
	next *Node
}

type MyLinkedList struct {
	head   *Node
	length int
}

func Constructor() MyLinkedList {
	return MyLinkedList{&Node{}, 0}
}

// 带头节点的链表
// 获取链表中第 index 个节点的值。如果索引无效，则返回-1
func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.length {
		return -1
	}
	p := this.head
	for i := 0; i <= index; i++ {
		p = p.next
	}
	return p.val
}

// 在链表的第一个元素之前添加一个值为 val 的节点。插入后，新节点将成为链表的第一个节点。
func (this *MyLinkedList) AddAtHead(val int) {
	this.head.next = &Node{val, this.head.next}
	this.length++
}

// 将值为 val 的节点追加到链表的最后一个元素。
func (this *MyLinkedList) AddAtTail(val int) {
	// 先找到链表的最后一个元素
	p := this.head
	for p.next != nil {
		p = p.next
	}
	p.next = &Node{val, nil}
	this.length++
}

// 在链表中的第 index 个节点之前添加值为 val  的节点。如果 index 等于链表的长度，则该节点将附加到链表的末尾。如果 index 大于链表长度，则不会插入节点。如果index小于0，则在头部插入节点。
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index <= 0 {
		this.AddAtHead(val)
		return
	} else if index == this.length {
		this.AddAtTail(val)
		return
	} else if index > this.length {
		return
	}

	p := this.head
	// 在第index个节点之前添加值，则需要找到第index-1个结点的位置
	for i := 0; i < index; i++ {
		p = p.next
	}
	q := &Node{val, nil}
	q.next = p.next
	p.next = q
	this.length++
}

// 如果索引 index 有效，则删除链表中的第 index 个节点
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.length {
		return
	}
	p := this.head
	// 删除第index个结点，则需要找到第index-1个结点
	for i := 0; i < index; i++ {
		p = p.next
	}
	q := p.next
	p.next = q.next
	q.next = nil
	this.length--
}

//打印完整链表
func (this *MyLinkedList) PrintList() {
	p := this.head.next
	fmt.Println("----------print list-----------")
	for p != nil {
		fmt.Print(p.val)
		fmt.Print(" ")
		p = p.next
	}
	fmt.Println()
	fmt.Println("-------------------------------")
}
