// Examples Taken From https://dev.to/jpoly1219/linked-lists-in-go-3g63#:~:text=If%20you%20are%20a%20student%20majoring
package main

func main() {
	list := LinkedList{nil, 0}
}

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head   *Node
	length int
}

func (l *LinkedList) insertAtHead(data int) {
	temp1 := &Node{data, nil}

	if l.head == nil {
		l.head = temp1
	} else {
		temp2 := l.head
		l.head = temp1
		temp1.next = temp2
	}
	l.length += 1
}

func (l *LinkedList) insertAtTail(data int) {
	temp1 := &Node{data, nil}

	if l.head == nil {
		l.head = temp1
	} else {
		temp2 := l.head
		for temp2.next != nil {
			temp2 = temp2.next
		}
		temp2.next = temp1
	}
	l.length += 1
}

func (l *LinkedList) insert(n, data int) {
	if n == 0 {
		l.insertAtHead(data)
	} else if n == l.length-1 {
		l.insertAtTail(data)
	} else {
		temp1 := &Node{data, nil}
		temp2 := l.head

		for i := 0; i < n-1; i++ {
			temp2 = temp2.next
		}
		temp1.next = temp2.next
		temp2.next = temp1
		l.length += 1
	}
}
func (l *LinkedList) deleteAtHead() {
	temp := l.head
	l.head = temp.next

	l.length -= 1
}

func (l *LinkedList) deleteAtTail() {
	temp1 := l.head
	var temp2 *Node
	for temp1.next != nil {
		temp2 = temp1
		temp1 = temp1.next
	}
	temp2.next = nil

	l.length -= 1
}

func (l *LinkedList) delete(n int) {
	if n == 0 {
		l.deleteAtHead()
	} else if n == l.length-1 {
		l.deleteAtTail()
	} else {
		temp1 := l.head
		for i := 0; i < n-1; i++ {
			temp1 = temp1.next
		}
		temp2 := temp1.next
		temp1.next = temp2.next
		l.length -= 1
	}
}

func (l *LinkedList) Reverse() {
	var curr, prev, next *Node
	curr = l.head
	prev = nil

	for curr != nil {
		next = curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
	l.head = prev
}
