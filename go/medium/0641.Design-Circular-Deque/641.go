// https://leetcode.com/problems/design-circular-deque/
// Level: Medium

package leetcode

type Node struct {
	val  int
	prev *Node
	next *Node
}

type MyCircularDeque struct {
	size   int
	length int
	head   *Node
	tail   *Node
}

func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		size:   k,
		length: 0,
		head:   nil,
		tail:   nil,
	}
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}

	cur := &Node{val: value}
	if this.IsEmpty() {
		this.head = cur
		this.tail = cur
	} else {
		cur.next = this.head
		this.head.prev = cur
		this.head = cur
	}
	this.length++
	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}

	cur := &Node{val: value}
	if this.IsEmpty() {
		this.head = cur
		this.tail = cur
	} else {
		cur.prev = this.tail
		this.tail.next = cur
		this.tail = cur
	}
	this.length++
	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}
	if this.head == this.tail {
		this.head = nil
		this.tail = nil
	} else {
		this.head = this.head.next
		this.head.prev = nil
	}
	this.length--
	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}

	if this.head == this.tail {
		this.head = nil
		this.tail = nil
	} else {
		this.tail = this.tail.prev
		this.tail.next = nil
	}
	this.length--
	return true
}

func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}

	return this.head.val
}

func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}

	return this.tail.val
}

func (this *MyCircularDeque) IsEmpty() bool {
	return this.length == 0
}

func (this *MyCircularDeque) IsFull() bool {
	return this.size == this.length
}
